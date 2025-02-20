package main

import (
  "image"
  "image/color"
  "image/png"
  "image/draw"
  "os"
  "math/rand"
)

type colorPoint struct {
  point image.Point
  color color.Color
}

// Width and Height of png 
const WIDTH = 500
const HEIGHT = 500

// Rates that are 1 out of x
const KILLRATE = 77
const BIRTHRATE = 35

// Max number of points that can be living
const MAXPOINTS = 50

// How fast the color changes
const VARIANCE = 5
const HALF_VAR = int(VARIANCE / 2)

func changeColor(color uint32) uint8 {
  return uint8(rand.Intn(VARIANCE) - HALF_VAR + int(color))
}

func main() {
  upLeft := image.Point{0, 0}
  lowRight := image.Point{WIDTH, HEIGHT}

  img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
  draw.Draw(img, img.Bounds(), image.White, image.Pt(0, 0), draw.Src)

  points := []colorPoint{colorPoint{image.Point{int(WIDTH / 2), 0}, color.Black}}

  for len(points) > 0 {
    tempPoints := []colorPoint{} 

    for i, point := range points {
      img.Set(point.point.X, point.point.Y, point.color)

      
      // If 0, kill this point
      deathNum := rand.Intn(KILLRATE)
      
      if deathNum == 0 && len(points) > 2 {
        continue
      }
  
      r, g, b, _ := point.color.RGBA()
      point.color = color.RGBA{changeColor(r), changeColor(g), changeColor(b), uint8(255)} 


      randNum := rand.Intn(70)

      switch {
      // Move left if not on border
      case randNum < 32:
        if point.point.X >= WIDTH - 1 {
          continue
        } else {
          tempPoints = append(tempPoints, colorPoint{image.Point{points[i].point.X + 1, points[i].point.Y}, point.color})
        }
      case randNum < 64:
        if point.point.X <= 0 {
          continue
        } else {
          tempPoints = append(tempPoints, colorPoint{image.Point{points[i].point.X - 1, points[i].point.Y}, point.color})
        }
      default:
        if point.point.Y >= HEIGHT - 1 {
          continue
        } else {
          tempPoints = append(tempPoints, colorPoint{image.Point{points[i].point.X, points[i].point.Y + 1}, point.color})
        }
      }
      createNum := rand.Intn(BIRTHRATE)

      if createNum == 0 && len(points) < MAXPOINTS {
        tempPoints = append(tempPoints, point)
      }
    }
    points = tempPoints
  }

  // Encode as PNG
  f, _ := os.Create("image.png")
  png.Encode(f, img)
}

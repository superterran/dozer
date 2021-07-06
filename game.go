package main

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

var unitSize int = 20

func drawLevel(da *gtk.DrawingArea, cr *cairo.Context) {

	offset := 0

	for y, row := range LevelMap {

		if y > height {
			break
		}

		for x, char := range row {

			if char == "S" && !isSpawned { // S is the maps spawn point
				playerX = x
				playerY = y
				// char = " "
				isSpawned = true
			}

			switch char {
			case "x":
				cr.SetSourceRGB(255, 0, 0)
			case "*":
				cr.SetSourceRGB(0, 255, 0)
			case " ":
				cr.SetSourceRGB(255, 255, 255)
			case "S":
				cr.SetSourceRGB(255, 255, 0)
			}

			cr.Rectangle(float64(offset+x*unitSize), float64(offset+y*unitSize), float64(unitSize), float64(unitSize))
			cr.Fill()

			if x == width {
				break
			}
		}
	}
}

func isPositionOccupied(x int, y int) bool {
	return LevelMap[int(y)][int(x)] != " " && LevelMap[int(y)][int(x)] != "S"
}

func movePlayer(origX int, origY int, newX int, newY int) {
	LevelMap[origY][origX] = " "
	LevelMap[newY][newX] = "S"
}

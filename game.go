package main

import (
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

func drawLevel(da *gtk.DrawingArea, cr *cairo.Context) {

	unitSize := 20

	offset := 70

	rows := strings.Split(LevelMap, "\n")
	for y, row := range rows {
		col := strings.Split(row, "")
		for x, char := range col {

			if char == "S" && !isSpawned { // S is the maps spawn point
				playerX = x
				playerY = y
				// char = " "
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
		}
	}
}

package main

import (
	"fmt"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

var unitSize int = 30

func drawLevel(da *gtk.DrawingArea, cr *cairo.Context) {

	for y, row := range LevelMap {

		if y > height {
			break
		}

		for x, char := range row {

			if char == "S" {
				if !isSpawned { // S is the maps spawn point

					origX = x
					origY = y

					playerX = x
					playerY = y

					frontX = x
					frontY = y

					isSpawned = true
				} else {
					if playerX != x || playerY != y {
						char = " "
					}

				}
			}

			var tileName string = ""

			switch char {

			case " ":
				// cr.SetSourceRGB(255, 50, 25)
				tileName = ""
			case "x":
				cr.SetSourceRGB(255, 0, 0)
				tileName = "tile008.png"
			case "*":
				cr.SetSourceRGB(0, 255, 0)
				tileName = "tile001.png"
			case "o":
				cr.SetSourceRGB(0, 0, 0)
				tileName = "tile010.png"
			case "S":
				cr.SetSourceRGB(255, 255, 0)

				switch playerDirection {
				case "up":
					tileName = "tile012.png"
				case "down":
					tileName = "tile013.png"
				case "left":
					tileName = "tile014.png"
				case "right":
					tileName = "tile011.png"
				}

			}

			// cr.IdentityMatrix()

			offset := 0
			_ = offset

			// cr.Save()

			if tileName != "" {
				surface, _ := cairo.NewSurfaceFromPNG("sprites/" + tileName)
				cr.SetSourceSurface(surface, float64(x*unitSize), float64(y*unitSize))
				cr.Paint()

			} else {
				cr.Rectangle(float64(offset+x*unitSize), float64(offset+y*unitSize), float64(unitSize), float64(unitSize))
				cr.Fill()
			}

			// cr.Clip()

			// cr.Restore()

			if x == width {
				break
			}
		}
	}
}

func isPushable() bool {

	if getChar(playerX, playerY) == "x" {

		if getChar(frontX, frontY) == " " || getChar(frontX, frontY) == "o" {

			return true
		}
	}

	return false
}

func isMovable(x int, y int) bool {
	return getChar(x, y) != " " && getChar(x, y) != "S"
}

func getChar(x int, y int) string {
	return LevelMap[int(y)][int(x)]
}

func setChar(x int, y int, char string) {
	LevelMap[int(y)][int(x)] = char
}

func movePlayer(origX int, origY int, newX int, newY int) {
	LevelMap[origY][origX] = " "
	LevelMap[newY][newX] = "S"
}

func push() {

	fmt.Println(holes)
	fmt.Println("pushed")

	setChar(origX, origY, " ")
	setChar(playerX, playerY, "S")

	if getChar(frontX, frontY) == "o" {
		setChar(frontX, frontY, "o")
		holes--
	} else {
		setChar(frontX, frontY, "x")
	}

	if holes == 0 {
		fmt.Println("you win!")
		LoadLevel(Level.GetString("next"))
	}

}

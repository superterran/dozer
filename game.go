package main

import (
	"fmt"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

var unitSize int = 30

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

				frontX = x
				frontY = y

				// char = " "
				isSpawned = true
			}

			switch char {
			case "x":
				cr.SetSourceRGB(255, 0, 0)
			case "*":
				cr.SetSourceRGB(0, 255, 0)
			case "o":
				cr.SetSourceRGB(0, 0, 0)
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

func isPushable() bool {

	// check to see if pushable area is even in bounds
	// if pushX <= 0 && pushY <= 0 {

	// 	if pushX <= width && pushY >= height {
	// 		// if in bounds, check to see if bulldozer is pushing a rock

	fmt.Println("Test")
	fmt.Println(getChar(playerX, playerY))

	if getChar(playerX, playerY) == "x" {
		// if it is a rock, check to see if pushable area is free for a rock
		fmt.Println("pushable?")

		if getChar(frontX, frontY) == " " || getChar(frontX, frontY) == "o" {

			fmt.Println("pushable!")
			return true
		}
	}

	// 	}
	// }
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

	fmt.Println("pushed")
	setChar(playerX, playerY, "S")

	if getChar(frontX, frontY) == "o" {
		setChar(frontX, frontY, "o")
	} else {
		setChar(frontX, frontY, "x")
	}

}

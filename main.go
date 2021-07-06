package main

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

var holes int = 0

// the bulldozer coordinates
var playerX int = 0.00
var playerY int = 0.00

// front of the bulldozer, for comparion purposes
var frontX int = 0.00
var frontY int = 0.00

// orig location
var origX int = 0.00
var origY int = 0.00

var isSpawned bool = false

var currentLevelName string = "level1.yaml"

var app *gtk.Application

var builder *gtk.Builder

const gladeTemplateFilename string = "main.glade"

const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)

func main() {

	const appID = "com.github.superterran.dozer"
	app, _ = gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

	path, _ := os.Getwd()

	if _, err := os.Stat(path + "/levels/" + currentLevelName); !os.IsNotExist(err) {

		LoadLevel(path + "/levels/" + currentLevelName)

	}

	app.Connect("activate", func() {

		builder, _ = gtk.BuilderNewFromFile(gladeTemplateFilename)

		createLoadDialog()
		createRestartDialog()
		createCloseDialog()
		createDrawArea()

		obj, _ := builder.GetObject("window")
		wnd := obj.(*gtk.Window)
		wnd.ShowAll()
		app.AddWindow(wnd)

		wnd.GetAllocation().GetWidth()

	})
	app.Run(os.Args)

}

func createDrawArea() {

	keyMap := map[uint]func(){
		KEY_LEFT: func() {
			playerX--
			frontX = playerX - 1
			frontY = playerY
			frontY = playerY

		},
		KEY_UP: func() {
			playerY--
			frontY = playerY - 1
			frontX = playerX
			frontX = playerX
		},
		KEY_RIGHT: func() {
			playerX++
			frontX = playerX + 1
			frontY = playerY
			frontY = playerY
		},
		KEY_DOWN: func() {
			playerY++
			frontY = playerY + 1
			frontX = playerX
			frontX = playerX
		},
	}

	winObj, _ := builder.GetObject("window")
	window := winObj.(*gtk.Window)

	optionsObj, _ := builder.GetObject("drawarea")
	da := optionsObj.(*gtk.DrawingArea)

	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {

		drawLevel(da, cr)

	})
	window.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		if move, found := keyMap[keyEvent.KeyVal()]; found {

			fmt.Println(playerX, playerY, frontX, frontY)

			origX := playerX
			origY := playerY

			move()

			if !isMovable(int(playerX), int(playerY)) {
				movePlayer(origX, origY, playerX, playerY)
				win.QueueDraw()
			} else {

				if isPushable() {

					push()
					win.QueueDraw()
				}

				playerX = origX
				playerY = origY
			}
		}
	})
}

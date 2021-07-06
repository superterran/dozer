package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

var playerX int = 0.00
var playerY int = 0.00
var isSpawned bool = false

const gladeTemplateFilename string = "main.glade"

const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)

func main() {

	const appID = "com.github.superterran.dozer"
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatalln("Couldn't create app:", err)
	}

	LoadMap("level1.yaml")

	app.Connect("activate", func() {

		builder, err := gtk.BuilderNewFromFile(gladeTemplateFilename)
		if err != nil {
			log.Fatalln("Couldn't make builder:", err)
		}

		createRestartDialog(app, builder)
		createCloseDialog(app, builder)
		createDrawArea(app, builder)

		obj, _ := builder.GetObject("window")
		wnd := obj.(*gtk.Window)
		wnd.ShowAll()
		app.AddWindow(wnd)

		wnd.GetAllocation().GetWidth()

	})
	app.Run(os.Args)

}

func createDrawArea(app *gtk.Application, builder *gtk.Builder) {

	keyMap := map[uint]func(){
		KEY_LEFT:  func() { playerX-- },
		KEY_UP:    func() { playerY-- },
		KEY_RIGHT: func() { playerX++ },
		KEY_DOWN:  func() { playerY++ },
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

			origX := playerX
			origY := playerY

			move()

			if !isPositionOccupied(int(playerX), int(playerY)) {
				movePlayer(origX, origY, playerX, playerY)
				win.QueueDraw()
			} else {
				playerX = origX
				playerY = origY
			}
		}
	})
}

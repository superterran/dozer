package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

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
	})
	app.Run(os.Args)

}

func createDrawArea(app *gtk.Application, builder *gtk.Builder) {

	// Data
	unitSize := 20.0
	x := 0.0
	y := 0.0
	keyMap := map[uint]func(){
		KEY_LEFT:  func() { x-- },
		KEY_UP:    func() { y-- },
		KEY_RIGHT: func() { x++ },
		KEY_DOWN:  func() { y++ },
	}

	winObj, _ := builder.GetObject("window")
	window := winObj.(*gtk.Window)

	optionsObj, _ := builder.GetObject("drawarea")
	da := optionsObj.(*gtk.DrawingArea)

	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
		cr.Fill()
	})
	window.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		if move, found := keyMap[keyEvent.KeyVal()]; found {
			move()
			win.QueueDraw()
		}
	})
}

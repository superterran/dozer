package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

const gladeTemplateFilename string = "main.glade"

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

		obj, _ := builder.GetObject("window")
		wnd := obj.(*gtk.Window)
		wnd.ShowAll()
		app.AddWindow(wnd)
	})
	app.Run(os.Args)
}

func createRestartDialog(app *gtk.Application, builder *gtk.Builder) {
	winObj, _ := builder.GetObject("window")
	window := winObj.(*gtk.Window)

	optionsObj, _ := builder.GetObject("restart")
	options := optionsObj.(*gtk.Button)

	_ = options.Connect("clicked", func() {
		dialog := gtk.MessageDialogNew(
			window,               //Specify the parent window
			gtk.DIALOG_MODAL,     //Modal dialog
			gtk.MESSAGE_QUESTION, //Specify the dialog box type
			gtk.BUTTONS_YES_NO,   //Default button
			"Are you sure you want to restart the level?") //Set content

		dialog.SetTitle("Restart Level") //Dialog box setting title

		flag := dialog.Run() //Run dialog
		if flag == gtk.RESPONSE_YES {
			fmt.Println("Press yes")
		} else if flag == gtk.RESPONSE_NO {
			fmt.Println("Press no")
		} else {
			fmt.Println("Press the close button")
		}

		dialog.Destroy() //Destroy the dialog

	})
}

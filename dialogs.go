package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

func createLoadDialog() {
	winObj, _ := builder.GetObject("window")
	win := winObj.(*gtk.Window)

	optionsObj, _ := builder.GetObject("load")
	options := optionsObj.(*gtk.MenuItem)

	_ = options.Connect("activate", func() {

		openDialog, err := gtk.FileChooserDialogNewWith2Buttons("Select files", win, gtk.FILE_CHOOSER_ACTION_OPEN,
			"Cancel", gtk.RESPONSE_CANCEL, "OK", gtk.RESPONSE_OK)
		if err != nil {
			log.Fatal("Dialog creation failed")
		}

		response := openDialog.Run()
		if response != gtk.RESPONSE_OK {
			log.Fatal("Error getting filename")
		}
		file := openDialog.GetFilename()
		openDialog.Destroy()

		fmt.Println(file)

		LoadLevel(file)

	})

}

func createRestartDialog() {
	winObj, _ := builder.GetObject("window")
	window := winObj.(*gtk.Window)

	optionsObj, _ := builder.GetObject("restart")
	options := optionsObj.(*gtk.MenuItem)

	_ = options.Connect("activate", func() {
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
			LoadLevel(currentLevelName)
		}

		dialog.Destroy() //Destroy the dialog

	})

}

func createCloseDialog() {
	winObj, _ := builder.GetObject("window")
	window := winObj.(*gtk.Window)

	optionsObj, _ := builder.GetObject("exit")
	options := optionsObj.(*gtk.MenuItem)

	_ = options.Connect("activate", func() {
		dialog := gtk.MessageDialogNew(
			window,                                  //Specify the parent window
			gtk.DIALOG_MODAL,                        //Modal dialog
			gtk.MESSAGE_QUESTION,                    //Specify the dialog box type
			gtk.BUTTONS_YES_NO,                      //Default button
			"Are you sure you want to close Dozer?") //Set content

		dialog.SetTitle("Exit Dozer") //Dialog box setting title

		flag := dialog.Run() //Run dialog
		if flag == gtk.RESPONSE_YES {
			os.Exit(0)
		} else if flag == gtk.RESPONSE_NO {
			fmt.Println("Press no")
		} else {
			fmt.Println("Press the close button")
		}

		dialog.Destroy() //Destroy the dialog

	})

}

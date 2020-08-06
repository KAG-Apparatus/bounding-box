package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "com.github.KAG-Apparatus.blob-editor"

var month map[int]string
var application *gtk.Application
var builder *gtk.Builder
var mainWindow *gtk.ApplicationWindow

func init() {
}

func main() {

	// Create a new application.
	var err error
	application, err = gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")

		// Get the GtkBuilder UI definition in the glade file.
		builder, err = gtk.BuilderNewFromFile("res/ui/main.glade")
		errorCheck(err)

		// Get the object with the id of "main_window".
		obj, err := builder.GetObject("main_window")
		errorCheck(err)
		// Verify that the object is a pointer to a gtk.ApplicationWindow.
		mainWindow, err = isWindow(obj)
		errorCheck(err)

		// Map the handlers to callback functions, and connect the signals
		// to the Builder.
		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
		}
		builder.ConnectSignals(signals)

		// Show the Window and all of its components.
		//mainWindow.Maximize()
		mainWindow.Show()
		application.AddWindow(mainWindow)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("Application finished.")
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}

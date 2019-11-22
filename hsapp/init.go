package hsapp

import (
	"log"

	"github.com/OpenDiablo2/HellSpawner/hsproj"
	"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/OpenDiablo2/D2Shared/d2data/d2mpq"

	"gopkg.in/alecthomas/kingpin.v2"
)

var argPath = kingpin.Arg("path", "Project path").String()

func Init() {
	// handle init of non-UI components
	log.Println("Launching HellSpawner...")
	// init project state to an empty state
	hsproj.SetDefaultActiveProject()
	// init cryptographic tables
	d2mpq.InitializeCryptoBuffer()

	// handle kingpin args
	kingpin.Parse()
	// if a path was passed in, load that project
	if *argPath != "" {
		hsproj.ActiveProject.PromptUnsavedChanges()
		hsproj.ActiveProject.Close()

		newproj, err := hsproj.LoadProjectStateFromFolder(*argPath)
		if err != nil {
			hsutil.PopupError(err)
		} else {
			hsproj.ActiveProject = newproj
		}
	}
}
package hswindows

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/OpenDiablo2/HellSpawner/hsproj"
	"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
)

type OpenProjectDialog struct {
	visible     bool
	currentDir  string
	directories []os.FileInfo
	loaded      func()
}

func CreateOpenProjectDialog(callback func()) OpenProjectDialog {
	result := OpenProjectDialog{}
	result.currentDir, _ = os.UserHomeDir()
	result.loaded = callback
	result.RefreshDirs()
	return result
}

func (v *OpenProjectDialog) Show(ctx *nk.Context) {
	v.visible = true
}

func (v *OpenProjectDialog) RefreshDirs() {
	v.directories, _ = ioutil.ReadDir(v.currentDir)
}

func (v *OpenProjectDialog) Render(win *glfw.Window, ctx *nk.Context) {
	if !v.visible {
		return
	}
	dialogWidth := 450
	dialogHeight := 300
	width, height := win.GetSize()
	bounds := nk.NkRect(float32((width/2)-(dialogWidth/2)), float32((height/2)-(dialogHeight/2)), float32(dialogWidth), float32(dialogHeight))
	if nk.NkBegin(ctx, "Open Project", bounds, nk.WindowClosable|nk.WindowBorder|nk.WindowMovable) > 0 {
		nk.NkLayoutRowDynamic(ctx, 18, 1)
		if nk.NkButtonLabel(ctx, "[Go up a folder]") > 0 {
			v.currentDir = path.Join(v.currentDir, "..")
			v.RefreshDirs()
		}
		for _, dir := range v.directories {
			if !dir.IsDir() {
				continue
			}
			dirName := dir.Name()
			if nk.NkButtonLabel(ctx, dirName) > 0 {
				v.currentDir = path.Join(v.currentDir, dirName)
				v.RefreshDirs()
			}
		}
		if nk.NkButtonLabel(ctx, "[Okay]") > 0 {
			hsproj.ActiveProject.PromptUnsavedChanges()
			hsproj.ActiveProject.Close()

			newproj, err := hsproj.LoadProjectStateFromFolder(v.currentDir)
			if err != nil {
				hsutil.PopupError(err)
				return
			}
			
			hsproj.ActiveProject = newproj
			v.visible = false
			v.loaded()
		}
	} else {
		v.visible = false
	}
	nk.NkEnd(ctx)
}

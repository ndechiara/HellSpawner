package hsapp

import (
	"os"
	"path/filepath"

	"github.com/OpenDiablo2/HellSpawner/hsinterface"

	"github.com/inkyblackness/imgui-go"

	"github.com/OpenDiablo2/HellSpawner/hsproj"
	"github.com/OpenDiablo2/HellSpawner/hsutil"
	"github.com/OpenDiablo2/HellSpawner/hsapp/hseditors"
)

var WindowWidth float32
var WindowHeight float32

type MainWindow struct {
	aboutDialog      *AboutDialog
	openFolderDialog *OpenFolderDialog
	projectTreeView  *ProjectTreeView
	dynamicWindows   []hsinterface.UIWindow

	doClose    bool
	menuHeight int
}

func CreateMainWindow() *MainWindow {
	icons := CreateMpqTreeIcons()

	result := &MainWindow{
		aboutDialog:    CreateAboutDialog(),
		dynamicWindows: make([]hsinterface.UIWindow, 0),
	}

	startdir, _ := os.UserHomeDir()
	result.openFolderDialog = CreateOpenFolderDialog("Select Project Folder", startdir, icons, func(folderPath string) {
		hsproj.ActiveProject.PromptUnsavedChanges()
		hsproj.ActiveProject.Close()

		newproj, err := hsproj.LoadProjectStateFromFolder(folderPath)
		if err != nil {
			hsutil.PopupError(err)
			return
		}

		hsproj.ActiveProject = newproj
		result.RefreshProjectLoaded()
	})

	result.projectTreeView = CreateProjectTreeView("ProjectTreeView", icons, result)

	return result
}

func (v MainWindow) DoClose() bool {
	return v.doClose
}

func (v *MainWindow) Render() {
	v.renderMainMenu()
	v.renderProjectTreeView()

	for windowIdx := range v.dynamicWindows {
		v.dynamicWindows[windowIdx].Render()
	}

	for windowIdx := 0; windowIdx < len(v.dynamicWindows); {
		if v.dynamicWindows[windowIdx].IsClosed() {
			v.dynamicWindows = append(v.dynamicWindows[:windowIdx], v.dynamicWindows[windowIdx+1:]...)
			continue
		}
		windowIdx++
	}

	v.aboutDialog.Render()
	v.openFolderDialog.Render()
}

func (v *MainWindow) renderMainMenu() {
	showAboutPopup := false
	showOpenFolderDialog := false
	if imgui.BeginMainMenuBar() {

		if imgui.BeginMenu("File") {
			if imgui.MenuItem("Open Folder") {
				v.openFolderDialog.Show()
				showOpenFolderDialog = true
			}
			if imgui.MenuItem("Exit") {
				v.doClose = true
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Help") {
			if imgui.MenuItem("About HellSpawner...") {
				showAboutPopup = true
			}
			imgui.EndMenu()
		}
		imgui.EndMainMenuBar()
	}

	// popups
	if showAboutPopup {
		imgui.OpenPopup(AboutDialogPopupName)
	}
	if showOpenFolderDialog {
		imgui.OpenPopup(v.openFolderDialog.PopupName)
	}
}

func (v *MainWindow) renderProjectTreeView() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0.0)
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 20})
	imgui.SetNextWindowSize(imgui.Vec2{X: 300, Y: WindowHeight - 20})
	if imgui.BeginV("Project Treeview", nil, imgui.WindowFlagsNoMove|imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoSavedSettings) {
		v.projectTreeView.Render()
		imgui.End()
	}
	imgui.PopStyleVar()
	/*imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0.0)
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 20})
	imgui.SetNextWindowSize(imgui.Vec2{X: 300, Y: WindowHeight - 20})
	if imgui.BeginV("Project Treeview", nil, imgui.WindowFlagsNoMove|imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoSavedSettings) {
		if imgui.TreeNode("Project") {
			imgui.TreePop()
		}

		imgui.End()
	}
	imgui.PopStyleVar()*/
}

// should call sub components and ask them to refresh because a new ActiveProject was loaded
func (v *MainWindow) RefreshProjectLoaded() {
	v.projectTreeView.Refresh()
}

func (v *MainWindow) OnFileSelected(filename string, mpqpath hsutil.MpqPath) {
	// try to open an editor based on the file extension
	var ed hsinterface.UIEditor
	ed = nil

	if filepath.Ext(mpqpath.FilePath) == ".txt" {
		ed = hseditors.CreateTextEditor(filename, mpqpath)
	}

	if ed != nil {
		v.dynamicWindows = append(v.dynamicWindows, hseditors.CreateEditorWindow(&ed))
	}
}

func (v *MainWindow) OnViewMpqFileDetails(mpqpath string) {
	v.dynamicWindows = append(v.dynamicWindows, CreateMpqDetailsDialog(hsproj.ActiveProject.FolderPath+mpqpath))
}

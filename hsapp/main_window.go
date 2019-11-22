package hsapp

import (
	"os"

	"github.com/inkyblackness/imgui-go"

	"github.com/OpenDiablo2/HellSpawner/hsproj"
	"github.com/OpenDiablo2/HellSpawner/hsutil"
)

var WindowWidth float32
var WindowHeight float32

type MainWindow struct {
	aboutDialog      *AboutDialog
	openFolderDialog *OpenFolderDialog
	projectTreeView  *ProjectTreeView

	doClose    bool
	menuHeight int
}

func CreateMainWindow() *MainWindow {
	icons := CreateMpqTreeIcons()

	result := &MainWindow{
		aboutDialog: CreateAboutDialog(),
	}

	startdir, _ := os.UserHomeDir()
	result.openFolderDialog = CreateOpenFolderDialog("Select Project Folder", startdir, icons, func(folderPath string){
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

	result.projectTreeView = CreateProjectTreeView("ProjectTreeView", icons, func(filename string, mpqpath hsutil.MpqPath){
		// do something when a file is selected
	})

	return result
}

func (v MainWindow) DoClose() bool {
	return v.doClose
}

func (v *MainWindow) Render() {
	v.renderMainMenu()
	v.renderProjectTreeView()
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
		v.projectTreeView.Render();
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
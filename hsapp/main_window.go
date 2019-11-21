package hsapp

import "github.com/inkyblackness/imgui-go"

var WindowWidth float32
var WindowHeight float32

type MainWindow struct {
	aboutDialog *AboutDialog
	doClose     bool
	menuHeight  int
}

func CreateMainWindow() *MainWindow {
	result := &MainWindow{
		aboutDialog: CreateAboutDialog(),
	}

	return result
}

func (v MainWindow) DoClose() bool {
	return v.doClose
}

func (v *MainWindow) Render() {
	v.renderMainMenu()
	v.renderProjectTreeView()
	v.aboutDialog.Render()
}

func (v *MainWindow) renderMainMenu() {
	showAboutPopup := false
	if imgui.BeginMainMenuBar() {

		if imgui.BeginMenu("File") {
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
	if showAboutPopup {
		imgui.OpenPopup(AboutDialogPopupName)
	}
}

func (v *MainWindow) renderProjectTreeView() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0.0)
	imgui.SetNextWindowPos(imgui.Vec2{0, 20})
	imgui.SetNextWindowSize(imgui.Vec2{300, WindowHeight - 20})
	if imgui.BeginV("Project Treeview", nil, imgui.WindowFlagsNoMove|imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoSavedSettings) {
		if imgui.TreeNode("Project") {
			imgui.TreePop()
		}

		imgui.End()
	}
	imgui.PopStyleVar()
}

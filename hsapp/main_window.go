package hsapp

import "github.com/inkyblackness/imgui-go"

type MainWindow struct {
	aboutDialog *AboutDialog
	doClose     bool
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
	v.aboutDialog.Render()
}

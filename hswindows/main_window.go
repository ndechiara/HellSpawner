package hswindows

import (
	"github.com/OpenDiablo2/HellSpawner/hsproj"
	"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
)

type MainWindow struct {
	aboutDialog       AboutDialog
	openProjectDialog OpenProjectDialog

	mpqTreeDialog *MpqTreeDialog
}

func CreateMainWindow() MainWindow {
	result := MainWindow{}
	result.mpqTreeDialog = CreateMpqTreeDialog()

	result.aboutDialog = CreateAboutDialog()
	result.openProjectDialog = CreateOpenProjectDialog(func() {
		// callback when a new project is loaded
		result.Refresh()
	})
	return result
}

func (v *MainWindow) Refresh(){
	v.mpqTreeDialog.Refresh()
}

var (
	MenuBarRowHeight int = 32
)

func (v *MainWindow) Render(win *glfw.Window, ctx *nk.Context) {
	width, _ := win.GetSize()
	bounds := nk.NkRect(0, 0, float32(width), 32)
	if nk.NkBegin(ctx, "Bla", bounds, nk.WindowNoScrollbar) > 0 {
		nk.NkMenubarBegin(ctx)
		nk.NkLayoutRowBegin(ctx, nk.LayoutStaticRow, 25, 3)
		nk.NkLayoutRowPush(ctx, 45)
		if nk.NkMenuBeginLabel(ctx, "File", nk.TextAlignLeft, nk.NkVec2(190, 200)) > 0 {
			nk.NkLayoutRowDynamic(ctx, 25, 1)
			if nk.NkMenuItemLabel(ctx, "Open Project Folder...", nk.TextAlignLeft) > 0 {
				v.openProjectDialog.Show(ctx)
			}
			if nk.NkMenuItemLabel(ctx, "Save", nk.TextAlignLeft) > 0 {
				err := hsproj.ActiveProject.Save()
				if err != nil {
					hsutil.PopupError(err)
				}
			}
			if nk.NkMenuItemLabel(ctx, "Save As", nk.TextAlignLeft) > 0 {
				// TODO
			}
			if nk.NkMenuItemLabel(ctx, "Quit", nk.TextAlignLeft) > 0 {
				win.SetShouldClose(true)
			}
			nk.NkMenuEnd(ctx)
		}
		if nk.NkMenuBeginLabel(ctx, "Help", nk.TextAlignLeft, nk.NkVec2(180, 200)) > 0 {
			nk.NkLayoutRowDynamic(ctx, 25, 1)
			if nk.NkMenuItemLabel(ctx, "About HellSpawner...", nk.TextAlignLeft) > 0 {
				v.aboutDialog.Show(ctx)
			}
			nk.NkMenuEnd(ctx)
		}
		nk.NkMenubarEnd(ctx)
	}
	nk.NkEnd(ctx)

	// tree view
	v.mpqTreeDialog.Render(win, ctx)

	// popups
	v.aboutDialog.Render(win, ctx)
	v.openProjectDialog.Render(win, ctx)
}

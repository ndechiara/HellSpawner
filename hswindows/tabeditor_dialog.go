package hswindows

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"

	"github.com/OpenDiablo2/HellSpawner/hswindows/hseditors"
)

type TabEditorDialog struct {
	visible      bool

	Icons *MpqTreeIcons

	Tabs []*EditorTab
}

type EditorTab struct {
	Visible bool
	Editor  *hseditors.Editor
}

func CreateTabEditorDialog() *TabEditorDialog {
	result := TabEditorDialog{}
	result.Tabs = make([]*EditorTab, 0)

	// load images
	result.Icons = CreateMpqTreeIcons()

	result.visible = true
	result.Refresh()
	return &result
}

func (v *TabEditorDialog) AddEditor(editor *hseditors.Editor) {
	tab := EditorTab{true, editor}
	for _, ed := range v.Tabs {
		ed.Visible = false
	}
	v.Tabs = append(v.Tabs, &tab)
}

func (v *TabEditorDialog) Show() {
	v.visible = true
}

func (v *TabEditorDialog) Refresh() {
	
}

func (v *TabEditorDialog) Render(win *glfw.Window, ctx *nk.Context) {
	if !v.visible {
		return
	}
	width, _ := win.GetSize()
	dialogWidth := width - 350
	dialogHeight := MenuBarRowHeight
	x := 353
	y := MenuBarRowHeight
	bounds := nk.NkRect(float32(x), float32(y), float32(dialogWidth), float32(dialogHeight))
	vis := false
	if nk.NkBegin(ctx, "Editor", bounds, nk.WindowNoScrollbar) > 0 {
		vis = true
		// display the tab bar
		nk.NkLayoutRowTemplateBegin(ctx, 18)
		for _, _ = range v.Tabs {
			nk.NkLayoutRowTemplatePushStatic(ctx, 18)
			nk.NkLayoutRowTemplatePushVariable(ctx, 100)
		}
		nk.NkLayoutRowTemplateEnd(ctx)
		for i, ed := range v.Tabs {
			img := v.Icons.GetIcon((*ed.Editor).Name())
			nk.NkImage(ctx, *img)
			if nk.NkSelectText(ctx, (*ed.Editor).Name(), int32(len((*ed.Editor).Name())), nk.TextLeft, 0) > 0 {
				// select this editor
				ed.Visible = true
				// set all other editors to non-visible
				for o, otherEd := range v.Tabs {
					if i != o {
						otherEd.Visible = false
					}
				}
			}
		}
	}
	nk.NkEnd(ctx)

	if vis {
		// actually display whatever tab is visible
		for _, ed := range v.Tabs {
			if ed.Visible {
				(*ed.Editor).Render(win, ctx, x, y + MenuBarRowHeight)
			}
		}
	}
}

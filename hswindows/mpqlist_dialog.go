package hswindows

import (
	//"path"

	"github.com/OpenDiablo2/HellSpawner/hsproj"
	//"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
)

type MpqListDialog struct {
	visible      bool
	names        []string
	AnySelected  bool
	SelectedName string
}

func CreateMpqListDialog() MpqListDialog {
	result := MpqListDialog{}
	result.Refresh()
	return result
}

func (v *MpqListDialog) Show(ctx *nk.Context) {
	v.visible = true
}

func (v *MpqListDialog) Refresh() {
	v.names = make([]string, 0)
	v.SelectedName = ""
	v.AnySelected = false
	if hsproj.ActiveProject.Loaded == false {
		return
	}
	for _, m := range hsproj.ActiveProject.MpqList.Mpqs {
		v.names = append(v.names, m.Data.FileName)
	}
}

func (v *MpqListDialog) Render(win *glfw.Window, ctx *nk.Context) {
	if !v.visible {
		return
	}
	dialogWidth := 450
	dialogHeight := 300
	width, height := win.GetSize()
	bounds := nk.NkRect(float32((width/2)-(dialogWidth/2)), float32((height/2)-(dialogHeight/2)), float32(dialogWidth), float32(dialogHeight))
	if nk.NkBegin(ctx, "MPQs", bounds, nk.WindowClosable|nk.WindowBorder|nk.WindowMovable|nk.WindowBackground) > 0 {
		nk.NkLayoutRowDynamic(ctx, 18, 1)
		for _, name := range v.names {
			if name == v.SelectedName {
				if nk.NkButtonLabel(ctx, "* " + name) > 0 {
					v.SelectedName = ""
					v.AnySelected = false
				}
				continue
			}
			if nk.NkButtonLabel(ctx, "  " + name) > 0 {
				v.SelectedName = name
				v.AnySelected = true
			}
		}
	} else {
		v.visible = false
	}
	nk.NkEnd(ctx)
}

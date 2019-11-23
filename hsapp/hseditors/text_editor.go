package hseditors

import (
	"github.com/OpenDiablo2/HellSpawner/hsutil"
	"github.com/OpenDiablo2/HellSpawner/hsproj"

	"github.com/inkyblackness/imgui-go"
)

type TextEditor struct {
	FileName string 
	MpqPath  hsutil.MpqPath

	Text *string

	unsavedChanges bool

	renderName string
}

func CreateTextEditor(filename string, mpqpath hsutil.MpqPath) *TextEditor {
	ed := TextEditor{}
	ed.FileName = filename
	ed.MpqPath = mpqpath
	ed.renderName = "TextEditor" + ed.MpqPath.MpqName + ":" + ed.MpqPath.FilePath + ":";

	// load the text itself
	bytes, err := hsproj.ActiveProject.MpqList.LoadFile(mpqpath)
	var str string
	if err == nil {
		str = string(bytes)
	} else {
		hsutil.PopupError(err)
		str = "Failed to load file"
	}
	ed.Text = &str

	return &ed
}

func (v *TextEditor) Name() string {
	return v.FileName
}

func (v *TextEditor) LongName() string {
	return v.MpqPath.MpqName + " :: " + v.MpqPath.FilePath
}

func (v *TextEditor) RenderName() string {
	return v.renderName
}

func (v *TextEditor) HasUnsavedChanges() bool {
	return v.unsavedChanges
}

func (v *TextEditor) Save() {
	// TODO
	v.unsavedChanges = false
}

func (v *TextEditor) editCallback(d imgui.InputTextCallbackData) int32 {
	v.unsavedChanges = true
	return 0
}

func (v *TextEditor) Render(size imgui.Vec2) {
	//if imgui.BeginChildV(v.renderName + "Child", size, false, imgui.WindowFlagsNoMove|imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoSavedSettings|imgui.WindowFlagsNoTitleBar) {
		imgui.InputTextMultilineV(v.renderName + "InputText", v.Text,
			imgui.Vec2{X: size.X - 10, Y: size.Y - 30},
			imgui.InputTextFlagsAllowTabInput|imgui.InputTextFlagsCallbackCharFilter,
			v.editCallback)
	//}
	//imgui.EndChild()
}
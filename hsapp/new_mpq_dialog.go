package hsapp

import "github.com/inkyblackness/imgui-go"

type NewMpqDialog struct {
	MpqName          string
	CompressListFile bool
}

var NewMpqDialogPopupName = "New MPQ##dialog"

func CreateNewMpqDialog() *NewMpqDialog {
	result := &NewMpqDialog{
		MpqName:          "Untitled.mpq",
		CompressListFile: true,
	}

	return result
}

func (v *NewMpqDialog) Render() {
	t := true

	imgui.SetNextWindowSize(imgui.Vec2{300, 0})
	if imgui.BeginPopupModalV(NewMpqDialogPopupName, &t, imgui.WindowFlagsNoResize) {
		imgui.AlignTextToFramePadding()
		imgui.Text("File Name")
		imgui.SameLine()
		imgui.InputText("##NewMpqDialog.MpqName", &v.MpqName)
		imgui.Checkbox("Compress ListFile##NewMpqDialog.CompressListFile", &v.CompressListFile)
		imgui.Separator()
		if imgui.Button("Create MPQ##NewMpqDialog.Create") {

		}
		if imgui.Button("Cancel##NewMpqDialog.Close") {
			imgui.CloseCurrentPopup()
		}
		imgui.EndPopup()
	}
}

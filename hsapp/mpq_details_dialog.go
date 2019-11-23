package hsapp

import (
	"path"

	"github.com/OpenDiablo2/D2Shared/d2data/d2mpq"
	"github.com/inkyblackness/imgui-go"
)

type MpqDetailsDialog struct {
	mpq    *d2mpq.MPQ
	isOpen bool
}

func CreateMpqDetailsDialog(mpqPath string) *MpqDetailsDialog {
	result := &MpqDetailsDialog{
		isOpen: true,
	}
	mpq, err := d2mpq.Load(mpqPath)
	if err != nil {
		panic(err)
	}
	result.mpq = mpq
	return result
}

func (v *MpqDetailsDialog) Render() {
	imgui.SetNextWindowSize(imgui.Vec2{300, 300})
	if imgui.BeginV(path.Base(v.mpq.FileName)+" Details##MpqDetailsDialog", &v.isOpen, imgui.WindowFlagsNoResize) {

	}
	imgui.End()
}

func (v MpqDetailsDialog) IsClosed() bool {
	return !v.isOpen
}

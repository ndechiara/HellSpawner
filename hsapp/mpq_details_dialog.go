package hsapp

import (
	"path"
	"strconv"

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
	imgui.SetNextWindowSize(imgui.Vec2{275, 0})
	if imgui.BeginV(path.Base(v.mpq.FileName)+" Details##MpqDetailsDialog", &v.isOpen, imgui.WindowFlagsNoResize) {
		imgui.Columns(2, "Details")

		imgui.Text("Hash Table Entries")
		imgui.NextColumn()
		imgui.Text(strconv.Itoa(len(v.mpq.HashTableEntries)))
		imgui.NextColumn()

		imgui.Text("Block Table Entries")
		imgui.NextColumn()
		imgui.Text(strconv.Itoa(len(v.mpq.BlockTableEntries)))
		imgui.NextColumn()

		imgui.Text("Archive Size")
		imgui.NextColumn()
		imgui.Text(strconv.Itoa(int(v.mpq.Data.ArchiveSize)))
		imgui.NextColumn()

		imgui.Text("Block Size")
		imgui.NextColumn()
		imgui.Text(strconv.Itoa(int(v.mpq.Data.BlockSize)))
		imgui.NextColumn()

		imgui.Text("Block Size")
		imgui.NextColumn()
		imgui.Text(strconv.Itoa(int(v.mpq.Data.BlockSize)))
		imgui.NextColumn()

		imgui.Text("Format Version")
		imgui.NextColumn()
		imgui.Text(strconv.Itoa(int(v.mpq.Data.FormatVersion)))
		imgui.NextColumn()

		imgui.Text("Block Table Offset")
		imgui.NextColumn()
		imgui.Text(strconv.Itoa(int(v.mpq.Data.BlockTableOffset)))
		imgui.NextColumn()

		imgui.Text("Hash Table Offset")
		imgui.NextColumn()
		imgui.Text(strconv.Itoa(int(v.mpq.Data.HashTableOffset)))
		imgui.NextColumn()

		imgui.Columns(1, "")
	}
	imgui.End()
}

func (v MpqDetailsDialog) IsClosed() bool {
	return !v.isOpen
}

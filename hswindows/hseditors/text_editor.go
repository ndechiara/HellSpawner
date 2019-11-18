package hseditors

import (
	"github.com/OpenDiablo2/HellSpawner/hsutil"
	"github.com/OpenDiablo2/HellSpawner/hsproj"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
)

type TextEditor struct {
	FileName string 
	MpqPath hsutil.MpqPath

	TextBytes []byte
}

func CreateTextEditor(filename string, mpqpath hsutil.MpqPath) *TextEditor {
	ed := TextEditor{}
	ed.FileName = filename
	ed.MpqPath = mpqpath

	// load the text itself
	bytes, err := hsproj.ActiveProject.MpqList.LoadFile(mpqpath)
	if err == nil {
		ed.TextBytes = bytes
	} else {
		hsutil.PopupError(err)
		ed.TextBytes = []byte("Failed to load file")
	}

	return &ed
}

func (v *TextEditor) Name() string {
	return v.FileName
}

func (v *TextEditor) Refresh() {

}
	
func (v *TextEditor) Render(win *glfw.Window, ctx *nk.Context, x int, y int) {
	width, height := win.GetSize()
	dialogWidth := width - 350 - 3
	dialogHeight := height - y - 1
	bounds := nk.NkRect(float32(x), float32(y), float32(dialogWidth), float32(dialogHeight))
	if nk.NkBegin(ctx, v.MpqPath.MpqName + " : " + v.MpqPath.FilePath, bounds, nk.WindowTitle) > 0 {
		nk.NkLayoutRowDynamic(ctx, float32(dialogHeight - 32), 1)
		nk.NkEditStringZeroTerminated(
			ctx,
			nk.EditEditor, // |nk.EditSelectable|nk.EditClipboard|nk.TextEditMultiLine, // nk.EditNoHorizontalScroll|
			v.TextBytes,
			int32(len(v.TextBytes)),
			nk.NkFilterAscii)
	}
	nk.NkEnd(ctx)
}
	
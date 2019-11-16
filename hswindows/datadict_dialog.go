package hswindows

import (
	"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
)

type DataDictDialog struct {
	visible bool
	data    hsutil.DataDictionary
}


func CreateDataDictDialog(path hsutil.MpqPath) DataDictDialog {
	result := DataDictDialog{}
	// todo: here we would set result.data by loading some passed in mpq name + filepath pair
	return result
}

func (v *DataDictDialog) Show(ctx *nk.Context) {
	v.visible = true
}

func (v *DataDictDialog) Render(win *glfw.Window, ctx *nk.Context) {
	// todo
}

package hsapp

import (
	"log"

	"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/inkyblackness/imgui-go"
)

type AboutDialog struct {
	logo *hsutil.Texture
}

var AboutDialogPopupName = "AboutHellspawnerDialog"

func CreateAboutDialog() *AboutDialog {
	result := &AboutDialog{}
	tex, err := hsutil.NewTextureFromFile("d2logo.png", gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		log.Fatal(err.Error())
	}
	result.logo = tex
	return result
}

func (v *AboutDialog) Render() {
	t := true
	if imgui.BeginPopupModalV(AboutDialogPopupName, &t, imgui.WindowFlagsNoResize) {
		imgui.Image(imgui.TextureID(v.logo.GetHandle()), imgui.Vec2{X: 256, Y: 256})
		imgui.Separator()
		imgui.Text("HellSpawner v0.0.1")
		imgui.Text("https://github.com/OpenDiablo2")
		imgui.PushItemWidth(256)
		imgui.EndPopup()
	}
}

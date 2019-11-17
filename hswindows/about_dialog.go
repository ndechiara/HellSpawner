package hswindows

import (
	"log"

	"github.com/OpenDiablo2/HellSpawner/hsutil"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
)

type AboutDialog struct {
	visible bool
	d2Logo  nk.Image
	texture *hsutil.Texture
}

func CreateAboutDialog() AboutDialog {
	result := AboutDialog{}
	tex, err := hsutil.NewTextureFromFile("d2logo.png", gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		log.Fatal(err.Error())
	}
	result.texture = tex
	result.d2Logo = nk.NkImageId(int32(tex.GetHandle()))
	return result
}

func (v *AboutDialog) Show(ctx *nk.Context) {
	v.visible = true
}

func (v *AboutDialog) Render(win *glfw.Window, ctx *nk.Context) {
	if !v.visible {
		return
	}
	dialogWidth := 320
	dialogHeight := 400
	width, height := win.GetSize()
	bounds := nk.NkRect(float32((width/2)-(dialogWidth/2)), float32((height/2)-(dialogHeight/2)), float32(dialogWidth), float32(dialogHeight))
	if nk.NkBegin(ctx, "About HellSpawner", bounds, nk.WindowClosable|nk.WindowBorder|nk.WindowMovable) > 0 {
		nk.NkLayoutRowDynamic(ctx, 256, 1)
		nk.NkImage(ctx, v.d2Logo)
		nk.NkLayoutRowDynamic(ctx, 18, 1)
		nk.NkLabelColored(ctx, "HellSpawner - OpenDiablo 2 toolkit", nk.TextAlignCentered, nk.NkRgb(255, 255, 255))
		nk.NkLabel(ctx, "Version 0.0.01", nk.TextAlignCentered)
		nk.NkLabel(ctx, "https://opendiablo2.com", nk.TextAlignCentered)
	} else {
		v.visible = false
	}
	nk.NkEnd(ctx)
}

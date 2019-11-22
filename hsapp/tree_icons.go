package hsapp

import (
	"log"
	"path/filepath"

	"github.com/go-gl/gl/v2.1/gl"

	"github.com/inkyblackness/imgui-go"

	"github.com/OpenDiablo2/HellSpawner/hsutil"
)

type MpqTreeIcons struct {
	ImageW float32
	ImageH float32

	unknownTex *hsutil.Texture
	txtTex     *hsutil.Texture
	binTex     *hsutil.Texture
	dccTex     *hsutil.Texture
	dc6Tex     *hsutil.Texture
	ds1Tex     *hsutil.Texture

	dirTex     *hsutil.Texture
	openDirTex *hsutil.Texture
	mpqTex     *hsutil.Texture
}

func CreateMpqTreeIcons() *MpqTreeIcons {
	icons := MpqTreeIcons{}
	icons.unknownTex = icons.loadImage(filepath.Join("icons","mpqtree_unknown.png"))
	icons.txtTex = icons.loadImage(filepath.Join("icons","mpqtree_txt.png"))
	icons.binTex = icons.loadImage(filepath.Join("icons","mpqtree_bin.png"))
	icons.dccTex = icons.loadImage(filepath.Join("icons","mpqtree_dcc.png"))
	icons.dc6Tex = icons.loadImage(filepath.Join("icons","mpqtree_dc6.png"))
	icons.ds1Tex = icons.loadImage(filepath.Join("icons","mpqtree_ds1.png"))

	icons.dirTex = icons.loadImage(filepath.Join("icons","mpqtree_dir.png"))
	icons.openDirTex = icons.loadImage(filepath.Join("icons","mpqtree_opendir.png"))
	icons.mpqTex = icons.loadImage(filepath.Join("icons","mpqtree_mpq.png"))

	icons.ImageW = 16
	icons.ImageH = 16
	return &icons
}

func (v *MpqTreeIcons) Size() imgui.Vec2 {
	return  imgui.Vec2{X: v.ImageW, Y: v.ImageH}
}

func (v *MpqTreeIcons) loadImage(path string) *hsutil.Texture {
	tex, err := hsutil.NewTextureFromFile(path, gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		hsutil.PopupError(err)
		if v.unknownTex != nil {
			return v.unknownTex
		}
		log.Fatal(err.Error())
	}
	return tex
}

func (v *MpqTreeIcons) GetIcon(path string) imgui.TextureID {
	if filepath.Ext(path) == ".txt" {
		return v.Txt()
	}
	if filepath.Ext(path) == ".bin" {
		return v.Bin()
	}
	if filepath.Ext(path) == ".dcc" {
		return v.Dcc()
	}
	if filepath.Ext(path) == ".dc6" {
		return v.Dc6()
	}
	if filepath.Ext(path) == ".ds1" {
		return v.Ds1()
	}
	if filepath.Ext(path) == ".mpq" {
		return v.Mpq()
	}

	return v.Unknown()
}

func (v *MpqTreeIcons) Unknown() imgui.TextureID {
	return imgui.TextureID(int32(v.unknownTex.GetHandle()))
}

func (v *MpqTreeIcons) Txt() imgui.TextureID {
	return imgui.TextureID(int32(v.txtTex.GetHandle()))
}

func (v *MpqTreeIcons) Bin() imgui.TextureID {
	return imgui.TextureID(int32(v.binTex.GetHandle()))
}

func (v *MpqTreeIcons) Dcc() imgui.TextureID {
	return imgui.TextureID(int32(v.dccTex.GetHandle()))
}

func (v *MpqTreeIcons) Dc6() imgui.TextureID {
	return imgui.TextureID(int32(v.dc6Tex.GetHandle()))
}

func (v *MpqTreeIcons) Ds1() imgui.TextureID {
	return imgui.TextureID(int32(v.ds1Tex.GetHandle()))
}

func (v *MpqTreeIcons) Dir() imgui.TextureID {
	return imgui.TextureID(int32(v.dirTex.GetHandle()))
}

func (v *MpqTreeIcons) OpenDir() imgui.TextureID {
	return imgui.TextureID(int32(v.openDirTex.GetHandle()))
}

func (v *MpqTreeIcons) Mpq() imgui.TextureID {
	return imgui.TextureID(int32(v.mpqTex.GetHandle()))
}
package hswindows

// todo: display list of MPQs, when expanded display items from listfile 
// associated with that MPQ. Allow user to add items to the listfile


import (
	"path/filepath"
	"log"

	"github.com/OpenDiablo2/HellSpawner/hsproj"
	"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/golang-ui/nuklear/nk"
)

type MpqTreeDialog struct {
	visible      bool
	names        []string
	AnySelected  bool
	SelectedName string
	fileTree     *hsutil.FileTreeNode

	Icons *MpqTreeIcons
}

type MpqTreeIcons struct {
	Unknown    *nk.Image
	unknownTex *hsutil.Texture
	Txt        *nk.Image
	txtTex     *hsutil.Texture
	Bin        *nk.Image
	binTex     *hsutil.Texture
	Dcc        *nk.Image
	dccTex     *hsutil.Texture
	Dc6        *nk.Image
	dc6Tex     *hsutil.Texture
	Ds1        *nk.Image
	ds1Tex     *hsutil.Texture
}

func CreateMpqTreeDialog() *MpqTreeDialog {
	result := MpqTreeDialog{}

	// load images
	result.Icons = CreateMpqTreeIcons()

	result.visible = true
	result.Refresh()
	return &result
}

func CreateMpqTreeIcons() *MpqTreeIcons {
	icons := MpqTreeIcons{}
	icons.unknownTex, icons.Unknown = icons.loadImage(filepath.Join("icons","mpqtree_unknown.png"))
	icons.txtTex, icons.Txt = icons.loadImage(filepath.Join("icons","mpqtree_txt.png"))
	icons.binTex, icons.Bin = icons.loadImage(filepath.Join("icons","mpqtree_bin.png"))
	icons.dccTex, icons.Dcc = icons.loadImage(filepath.Join("icons","mpqtree_dcc.png"))
	icons.dc6Tex, icons.Dc6 = icons.loadImage(filepath.Join("icons","mpqtree_dc6.png"))
	icons.ds1Tex, icons.Ds1 = icons.loadImage(filepath.Join("icons","mpqtree_ds1.png"))

	return &icons
}

func (v *MpqTreeIcons) loadImage(path string) (*hsutil.Texture, *nk.Image) {
	tex, err := hsutil.NewTextureFromFile(path, gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		hsutil.PopupError(err)
		if v.Unknown != nil {
			return v.unknownTex, v.Unknown
		}
		log.Fatal(err.Error())
	}
	result := nk.NkImageId(int32(tex.GetHandle()))
	return tex, &result
}

func (v *MpqTreeIcons) GetIcon(path string) *nk.Image {
	if filepath.Ext(path) == ".txt" {
		return v.Txt
	}
	if filepath.Ext(path) == ".bin" {
		return v.Bin
	}
	if filepath.Ext(path) == ".dcc" {
		return v.Dcc
	}
	if filepath.Ext(path) == ".dc6" {
		return v.Dc6
	}
	if filepath.Ext(path) == ".ds1" {
		return v.Ds1
	}

	return v.Unknown
}

func (v *MpqTreeDialog) Show(ctx *nk.Context) {
	v.visible = true
}

func (v *MpqTreeDialog) Refresh() {
	v.names = make([]string, 0)
	v.SelectedName = ""
	v.AnySelected = false
	if hsproj.ActiveProject.Loaded == false {
		return
	}
	// put every filename in a big list
	for _, m := range hsproj.ActiveProject.MpqList.Mpqs {
		for _, n := range m.ListFile.Files {
			v.names = append(v.names, filepath.Join(m.Name, n.Path))
		}
	}

	// build the tree
	v.fileTree = hsutil.BuildFileTreeFromFileList(v.names)
}

func (v *MpqTreeDialog) Render(win *glfw.Window, ctx *nk.Context) {
	if !v.visible {
		return
	}
	_, height := win.GetSize()
	dialogWidth := 350
	dialogHeight := height - MenuBarRowHeight - 1
	bounds := nk.NkRect(1, float32(MenuBarRowHeight), float32(dialogWidth), float32(dialogHeight))
	if nk.NkBegin(ctx, "MPQ Files", bounds, nk.WindowTitle) > 0 {
		nk.NkLayoutRowDynamic(ctx, 12, 1)
		v.RenderTree(ctx, v.fileTree)
	} else {
		v.visible = false
	}
	nk.NkEnd(ctx)
}

func (v *MpqTreeDialog) RenderTree(ctx *nk.Context, node *hsutil.FileTreeNode) {
	if node == nil {
		return
	}
	if node.IsFile {
		img := v.Icons.GetIcon(node.Name)
		nk.NkLayoutRowTemplateBegin(ctx, 18)
		nk.NkLayoutRowTemplatePushStatic(ctx, 18)
		nk.NkLayoutRowTemplatePushVariable(ctx, 100)
		nk.NkLayoutRowTemplateEnd(ctx)
		nk.NkImage(ctx, *img)
		if nk.NkSelectText(ctx, node.Name, int32(len(node.Name)), nk.TextLeft, 0) > 0 {
			// do something when file is selected
		}
	} else {
		if nk.NkTreePushHashed(ctx, nk.TreeTab, node.Name, nk.Minimized, node.FullPath, int32(len(node.Name)), 0) > 0 {
			for _, c := range node.Children {
				v.RenderTree(ctx, c)
			}
			nk.NkTreePop(ctx)
		}
	}
}

package hswindows

// todo: display list of MPQs, when expanded display items from listfile 
// associated with that MPQ. Allow user to add items to the listfile


import (
	"path/filepath"

	"github.com/OpenDiablo2/HellSpawner/hsproj"
	"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
)

type MpqTreeDialog struct {
	visible      bool
	names        []string
	AnySelected  bool
	SelectedName string
	fileTree     *hsutil.FileTreeNode
}

func CreateMpqTreeDialog() *MpqTreeDialog {
	result := MpqTreeDialog{}
	result.visible = true
	result.Refresh()
	return &result
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
	dialogWidth := 400
	dialogHeight := 600
	width, height := win.GetSize()
	bounds := nk.NkRect(float32((width/2)-(dialogWidth/2)), float32((height/2)-(dialogHeight/2)), float32(dialogWidth), float32(dialogHeight))
	if nk.NkBegin(ctx, "MPQ Tree", bounds, nk.WindowClosable|nk.WindowBorder|nk.WindowMovable|nk.WindowBackground) > 0 {
		nk.NkLayoutRowDynamic(ctx, 18, 1)
		RenderTree(ctx, v.fileTree)
	} else {
		v.visible = false
	}
	nk.NkEnd(ctx)
}

func RenderTree(ctx *nk.Context, node *hsutil.FileTreeNode) {
	if node == nil {
		return
	}
	if node.IsFile {
		if nk.NkButtonLabel(ctx, node.Name) > 0 {
			// do something when file is selected
		}
	} else {
		if nk.NkTreePushHashed(ctx, nk.TreeTab, node.Name, nk.Minimized, node.Name, int32(len(node.Name)), 0) > 0 {
			for _, c := range node.Children {
				RenderTree(ctx, c)
			}
			nk.NkTreePop(ctx)
		}
	}
}

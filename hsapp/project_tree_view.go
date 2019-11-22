package hsapp

import (
	"path/filepath"
	"strconv"

	"github.com/inkyblackness/imgui-go"

	"github.com/OpenDiablo2/HellSpawner/hsutil"
	"github.com/OpenDiablo2/HellSpawner/hsproj"
)

type ProjectTreeView struct {
	ViewName string // must have a unique popupName from other places this dialog is used
	callback *func(filename string, mpqpath hsutil.MpqPath)
	fileTree *hsutil.FileTreeNode
	names    []string
	show     bool

	Icons *MpqTreeIcons
}

func CreateProjectTreeView(viewName string, icons *MpqTreeIcons, callback func(filename string, mpqpath hsutil.MpqPath)) *ProjectTreeView {
	result := &ProjectTreeView{}
	result.ViewName = viewName
	result.callback = &callback
	result.show = true
	result.Icons = icons
	result.Refresh()
	return result
}

func (v *ProjectTreeView) Render() {
	if v.show {
		imgui.BeginChild(v.ViewName) //, false, imgui.WindowFlagsNoMove|imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoSavedSettings)
		imgui.PushStyleColor(imgui.StyleColorButton, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		v.renderTree(v.fileTree, 0)
		imgui.PopStyleColor()
		imgui.EndChild()
	}
}

func (v *ProjectTreeView) renderTree(node *hsutil.FileTreeNode, level int) {
	if node == nil {
		return
	}
	for i := 0; i < level; i++ {
		imgui.Spacing()
		imgui.SameLine()
	}
	if node.IsFile {
		img := v.Icons.GetIcon(node.Name)
		imgui.PushStyleColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		imgui.PushID(v.ViewName + strconv.Itoa(node.Id) + "img")
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:0,Y:0})
		if imgui.ImageButton(img, v.Icons.Size()) {
			if v.callback != nil {
				(*v.callback)(node.Name, node.GetMpqPath())
			}
		}
		imgui.PopStyleVar()
		imgui.PopID()
		imgui.PopStyleColor()
		imgui.PopStyleColor()

		imgui.SameLine()
		imgui.PushID(v.ViewName + strconv.Itoa(node.Id) + "btn")
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:0,Y:3})
		if imgui.Button(node.Name) {
			if v.callback != nil {
				(*v.callback)(node.Name, node.GetMpqPath())
			}
		}
		imgui.PopStyleVar()
		imgui.PopID()
	} else {
		var img imgui.TextureID
		if node.Open {
			img = v.Icons.OpenDir()
		} else {
			img = v.Icons.Dir()
		}

		imgui.PushStyleColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		imgui.PushID(v.ViewName + strconv.Itoa(node.Id) + "img")
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:0,Y:0})
		if imgui.ImageButton(img, v.Icons.Size()) {
			node.Open = !node.Open
		}
		imgui.PopStyleVar()
		imgui.PopID()
		imgui.PopStyleColor()
		imgui.PopStyleColor()

		imgui.SameLine()
		imgui.PushID(v.ViewName + strconv.Itoa(node.Id) + "btn")
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:0,Y:3})
		if imgui.Button(node.Name) {
			node.Open = !node.Open
		}
		imgui.PopStyleVar()
		imgui.PopID()

		if node.Open {
			level++
			for _, c := range node.Children {
				v.renderTree(c, level)
			}
		}

	}
}

func (v *ProjectTreeView) Refresh() {
	v.names = make([]string, 0)
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

func (v *ProjectTreeView) Show() {
	v.show = true
}


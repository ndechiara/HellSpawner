package hsapp

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/HellSpawner/hsinterface"

	"github.com/inkyblackness/imgui-go"

	"github.com/OpenDiablo2/HellSpawner/hsproj"
	"github.com/OpenDiablo2/HellSpawner/hsutil"
)

type ProjectTreeView struct {
	ViewName         string // must have a unique popupName from other places this dialog is used
	Icons            *MpqTreeIcons
	fileTree         *hsutil.FileTreeNode
	names            []string
	show             bool
	newMpqDialog     *NewMpqDialog
	showNewMpqDialog bool
	projectManager   hsinterface.ProjectManager
}

func CreateProjectTreeView(viewName string, icons *MpqTreeIcons, projectManager hsinterface.ProjectManager) *ProjectTreeView {
	result := &ProjectTreeView{
		projectManager: projectManager,
		newMpqDialog:   CreateNewMpqDialog(),
	}
	result.ViewName = viewName
	result.show = true
	result.Icons = icons
	result.Refresh()
	return result
}

func (v *ProjectTreeView) Render() {
	v.showNewMpqDialog = false
	if v.show {
		imgui.BeginChild(v.ViewName) //, false, imgui.WindowFlagsNoMove|imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoSavedSettings)
		imgui.PushStyleColor(imgui.StyleColorButton, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		v.renderTree(v.fileTree, 0)
		imgui.PopStyleColor()
		imgui.EndChild()
	}
	v.newMpqDialog.Render()
	// Dialogs
	if v.showNewMpqDialog {
		imgui.OpenPopup(NewMpqDialogPopupName)
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
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 0})
		if imgui.ImageButton(img, v.Icons.Size()) {
			v.projectManager.OnFileSelected(node.Name, node.GetMpqPath())
		}
		imgui.PopStyleVar()
		imgui.PopID()
		imgui.PopStyleColor()
		imgui.PopStyleColor()

		imgui.SameLine()
		imgui.PushID(v.ViewName + strconv.Itoa(node.Id) + "btn")
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 3})
		if imgui.Button(node.Name) {
			v.projectManager.OnFileSelected(node.Name, node.GetMpqPath())
		}
		imgui.PopStyleVar()
		imgui.PopID()

		return
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
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 0})
		if imgui.ImageButton(img, v.Icons.Size()) {
			node.Open = !node.Open
		}
		imgui.PopStyleVar()
		imgui.PopID()
		imgui.PopStyleColor()
		imgui.PopStyleColor()

		imgui.SameLine()
		imgui.PushID(v.ViewName + strconv.Itoa(node.Id) + "btn")
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 3})
		if imgui.Button(node.Name) {
			node.Open = !node.Open
		}
		imgui.PopStyleVar()
		imgui.PopID()

		imgui.PushID(v.ViewName + strconv.Itoa(node.Id) + "menuNew")
		if level == 0 && imgui.BeginPopupContextItem() {
			if imgui.BeginMenu("New") {
				if imgui.MenuItem("MPQ Package...") {
					v.showNewMpqDialog = true
				}
				imgui.EndMenu()
			}
			imgui.EndPopup()
		}
		imgui.PopID()

		imgui.PushID(v.ViewName + strconv.Itoa(node.Id) + "menuMpqDetails")
		if strings.HasSuffix(strings.ToLower(node.Name), ".mpq") && imgui.BeginPopupContextItem() {
			if imgui.MenuItem("MPQ Details...") {
				v.projectManager.OnViewMpqFileDetails(node.FullPath)
			}
			imgui.EndPopup()
		}
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

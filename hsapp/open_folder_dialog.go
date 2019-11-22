package hsapp

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"path/filepath"

	"github.com/inkyblackness/imgui-go"
)

type OpenFolderDialog struct {
	PopupName  string // must have a unique popupName from other places this dialog is used
	folderPath string
	callback   func(folderPath string)
	folders    []os.FileInfo
	show       bool
	Icons      *MpqTreeIcons
}

func CreateOpenFolderDialog(popupName string, startPath string, icons *MpqTreeIcons, callback func(folderPath string)) *OpenFolderDialog {
	result := &OpenFolderDialog{}
	result.PopupName = popupName
	result.folderPath = startPath
	result.callback = callback
	result.show = false
	result.Icons = icons
	result.Refresh()
	return result
}

func (v *OpenFolderDialog) Render() {
	if imgui.BeginPopupModalV(v.PopupName, &v.show, imgui.WindowFlagsNoResize) {
		imgui.BeginChildV(v.PopupName + "_FoldersList", imgui.Vec2{X: 250, Y: 300}, false, imgui.WindowFlagsNoMove|imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoSavedSettings)
		imgui.PushStyleColor(imgui.StyleColorButton, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})

		opendirimg := v.Icons.OpenDir()
		imgui.PushStyleColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		imgui.PushID("uponelevelimg")
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:0,Y:0})
		if imgui.ImageButton(opendirimg, v.Icons.Size()) {
			v.folderPath = path.Join(v.folderPath, "..")
			v.Refresh()
		}
		imgui.PopStyleVar()
		imgui.PopID()
		imgui.PopStyleColor()
		imgui.PopStyleColor()

		imgui.SameLine()
		imgui.PushID("uponelevelbtn")
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:0,Y:3})
		if imgui.Button("[Up One Level]") {
			v.folderPath = path.Join(v.folderPath, "..")
			v.Refresh()
		}
		imgui.PopStyleVar()
		imgui.PopID()

		dirimg := v.Icons.Dir()
		for i, folder := range v.folders {
			if !folder.IsDir() {
				continue
			}
			name := folder.Name()

			imgui.Spacing()

			imgui.SameLine()
			imgui.PushStyleColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
			imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
			imgui.PushID(strconv.Itoa(i) + "img")
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:0,Y:0})
			if imgui.ImageButton(dirimg, v.Icons.Size()) {
				v.folderPath = path.Join(v.folderPath, name)
				v.Refresh()
			}
			imgui.PopStyleVar()
			imgui.PopID()
			imgui.PopStyleColor()
			imgui.PopStyleColor()
			
			imgui.SameLine()
			imgui.PushID(name + strconv.Itoa(i) + "btn")
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:1,Y:3})
			if imgui.Button(name) {
				v.folderPath = path.Join(v.folderPath, name)
				v.Refresh()
			}
			imgui.PopStyleVar()
			imgui.PopID()
		}

		mpqicon := v.Icons.Mpq()
		for _, folder := range v.folders {
			if folder.IsDir() {
				continue
			}
			if strings.ToLower(filepath.Ext(folder.Name())) != ".mpq" {
				continue
			}
			imgui.Spacing()

			imgui.SameLine()
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:0,Y:0})
			imgui.Image(mpqicon, v.Icons.Size())
			imgui.PopStyleVar()

			imgui.SameLine()
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X:1,Y:3})
			imgui.PushStyleColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
			imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
			imgui.Button(folder.Name())
			imgui.PopStyleVar()
			imgui.PopStyleColor()
			imgui.PopStyleColor()
		}

		imgui.PopStyleColor()
		imgui.EndChild()
		imgui.Separator()
		
		if imgui.Button("Okay") {
			v.callback(v.folderPath)
			v.show = false
		}
		imgui.SameLine()
		if imgui.Button("Cancel") {
			v.show = false
		}
		imgui.EndPopup()
	}
}

func (v *OpenFolderDialog) Refresh() {
	v.folders, _ = ioutil.ReadDir(v.folderPath)
}

func (v *OpenFolderDialog) Show() {
	v.show = true
}
package hseditors

import (
	"github.com/OpenDiablo2/HellSpawner/hsinterface"

	"github.com/inkyblackness/imgui-go"
)

type EditorWindow struct {
	Editor *hsinterface.UIEditor

	isOpen     bool
	renderName string
}


func CreateEditorWindow(editor *hsinterface.UIEditor) *EditorWindow {
	result := &EditorWindow{
		isOpen: true,
	}
	result.Editor = editor
	result.renderName = "##EditorWindow" + (*editor).RenderName();
	return result
}

func (v *EditorWindow) Render() {
	if imgui.BeginV((*v.Editor).Name()+v.renderName, &v.isOpen, imgui.WindowFlagsNoScrollbar|imgui.WindowFlagsNoScrollWithMouse) {
		(*v.Editor).Render(imgui.Vec2{X: imgui.WindowWidth(), Y: imgui.WindowHeight()})
	}
	imgui.End()
}

func (v *EditorWindow) IsClosed() bool {
	return !v.isOpen
}
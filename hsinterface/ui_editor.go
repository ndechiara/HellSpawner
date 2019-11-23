package hsinterface

import "github.com/inkyblackness/imgui-go"

type UIEditor interface {
	Name() string // e.g. for tab display
	LongName() string // e.g. for window display
	RenderName() string
	Render(size imgui.Vec2)
	Save()
	HasUnsavedChanges() bool
}
package hseditors

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
)

type Editor interface {
	Render(win *glfw.Window, ctx *nk.Context, x int, y int)
	Refresh()
	Name() string
}
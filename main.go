package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/OpenDiablo2/HellSpawner/hsapp"

	"github.com/inkyblackness/imgui-go"
)

func main() {
	runtime.LockOSThread()
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := NewGLFW(io, GLFWClientAPIOpenGL2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := NewOpenGL2(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	hsapp.Run(platform, renderer)
}

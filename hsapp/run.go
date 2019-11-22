package hsapp

import (
	"time"

	"github.com/inkyblackness/imgui-go"
)

// Platform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the begin of a render pass. It must update the imgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically this causes the display buffer to be swapped.
	PostRender()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)
}

type clipboard struct {
	platform Platform
}

func (board clipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board clipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

// Renderer covers rendering imgui draw data.
type Renderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [4]float32)
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
}

// Run implements the main program loop of the demo. It returns when the platform signals to stop.
// This demo application shows some basic features of ImGui, as well as exposing the standard demo window.
func Run(p Platform, r Renderer) {
	Init() // call initialization of the HellSpawner system

	mainWindow := CreateMainWindow()
	imgui.CurrentIO().SetClipboard(clipboard{platform: p})
	clearColor := [4]float32{0.1, 0.1, 0.1, 1.0}
	for !p.ShouldStop() {
		p.ProcessEvents()
		p.NewFrame()
		imgui.NewFrame()
		//imgui.ShowDemoWindow(true)
		mainWindow.Render()
		if mainWindow.DoClose() {
			break
		}
		imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.
		r.PreRender(clearColor)
		// A this point, the application could perform its own rendering...
		r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
		p.PostRender()
		// sleep to avoid 100% CPU usage for this demo
		<-time.After(time.Millisecond * 10)
	}
}

package main

// choco install mingw

import (
	"runtime"
	"time"

	"github.com/golang-ui/nuklear/nk"
	"github.com/xlab/closer"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	maxVertexBuffer  = 512 * 1024
	maxElementBuffer = 128 * 1024
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(1280, 720, "OpenDiablo 2 HellSpawner", nil, nil)
	if err != nil {
		closer.Fatalln(err)
	}
	win.MakeContextCurrent()
	width, height := win.GetSize()
	if err := gl.Init(); err != nil {
		closer.Fatalln("opengl: init failed:", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))
	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)

	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	sansFont := nk.NkFontAtlasAddFromFile(atlas, "Cascadia.ttf", 17, nil)
	config := nk.NkFontConfig(17)
	config.SetOversample(0, 0)
	//config.SetRange(nk.NkFontChineseGlyphRanges())
	// simsunFont := nk.NkFontAtlasAddFromFile(atlas, "/Library/Fonts/Microsoft/SimHei.ttf", 14, &config)
	nk.NkFontStashEnd()
	nk.NkStyleSetFont(ctx, sansFont.Handle())
	// if simsunFont != nil {
	// 	nk.NkStyleSetFont(ctx, simsunFont.Handle())
	// }
	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	closer.Bind(func() {
		close(exitC)
		<-doneC
	})

	state := &State{
		bgColor: nk.NkRgba(10, 10, 10, 255),
	}

	fpsTicker := time.NewTicker(time.Second / 30)
	for {
		select {
		case <-exitC:
			nk.NkPlatformShutdown()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if win.ShouldClose() {
				close(exitC)
				continue
			}
			glfw.PollEvents()
			gfxMain(win, ctx, state)
		}
	}
}

func gfxMain(win *glfw.Window, ctx *nk.Context, state *State) {
	width, height := win.GetSize()
	nk.NkPlatformNewFrame()
	bounds := nk.NkRect(0, 0, float32(width), 32)
	if nk.NkBegin(ctx, "Bla", bounds, nk.WindowNoScrollbar|nk.WindowBackground) > 0 {
		nk.NkMenubarBegin(ctx)
		nk.NkLayoutRowBegin(ctx, nk.LayoutStaticRow, 25, 3)
		nk.NkLayoutRowPush(ctx, 45)
		if nk.NkMenuBeginLabel(ctx, "File", nk.TextAlignLeft, nk.NkVec2(120, 200)) > 0 {
			nk.NkLayoutRowDynamic(ctx, 25, 1)
			if nk.NkMenuItemLabel(ctx, "Quit", nk.TextAlignLeft) > 0 {
				win.SetShouldClose(true)
			}
			nk.NkMenuEnd(ctx)
		}
		nk.NkMenubarEnd(ctx)
	}
	nk.NkEnd(ctx)
	// Render
	bg := make([]float32, 4)
	nk.NkColorFv(bg, state.bgColor)
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(bg[0], bg[1], bg[2], bg[3])
	nk.NkPlatformRender(nk.AntiAliasingOff, maxVertexBuffer, maxElementBuffer)
	win.SwapBuffers()
}

type State struct {
	bgColor nk.Color
}

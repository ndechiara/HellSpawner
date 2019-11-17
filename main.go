package main

// choco install mingw

import (
	"runtime"
	"time"
	"log"

	"github.com/OpenDiablo2/HellSpawner/hswindows"
	"github.com/OpenDiablo2/HellSpawner/hsproj"

	"github.com/golang-ui/nuklear/nk"
	"github.com/xlab/closer"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/OpenDiablo2/D2Shared/d2data/d2mpq"
)

const (
	maxVertexBuffer  = 512 * 1024
	maxElementBuffer = 128 * 1024
)

func init() {
	runtime.LockOSThread()
}

var mainWindow hswindows.MainWindow

// please see https://github.com/golang-ui/nuklear/issues/48
// it is important that we prevent fonts / font handles from being garbage collected
// otherwise we will see obscure "signal arrived during external code execution" panics
// when you create a font, store it and its handle in a var as such to preven this issue
var (
	MainFont *nk.Font

	MainFontHandle *nk.UserFont
)

func main() {
	// startup GLFW 
	// (don't let anything come before this)
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

	// handle init of non-UI components
	log.Println("Launching HellSpawner...")
	// init project state to an empty state
	hsproj.SetDefaultActiveProject()
	// init cryptographic tables
	d2mpq.InitializeCryptoBuffer()

	// finish UI init
	win.MakeContextCurrent()
	width, height := win.GetSize()
	if err := gl.Init(); err != nil {
		closer.Fatalln("opengl: init failed:", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))
	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)
	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	MainFont = nk.NkFontAtlasAddFromFile(atlas, "Roboto-Regular.ttf", 18, nil)
	config := nk.NkFontConfig(18)
	config.SetPixelSnap(false)
	config.SetOversample(4, 4)
	//config.SetRange(nk.NkFontChineseGlyphRanges())
	// simsunFont := nk.NkFontAtlasAddFromFile(atlas, "/Library/Fonts/Microsoft/SimHei.ttf", 14, &config)
	nk.NkFontStashEnd()

	MainFontHandle = MainFont.Handle()
	nk.NkStyleSetFont(ctx, MainFontHandle)
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
		bgColor: nk.NkRgba(20, 20, 20, 255),
	}

	fpsTicker := time.NewTicker(time.Second / 30)
	mainWindow = hswindows.CreateMainWindow()
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
	nk.NkPlatformNewFrame()
	mainWindow.Render(win, ctx)
	// Render
	width, height := win.GetSize()
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

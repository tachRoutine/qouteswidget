package main

import (
	"embed"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	//setup auto start on device startup
}

func main() {
	app := NewApp()

	// systray := app.NewSystemTray()
	// systray.SetLabel("My App")
	// systray.SetIcon(iconBytes)

	err := wails.Run(&options.App{
		Title:         "HeyMe",
		Width:         360,
		Height:        160,
		DisableResize: true,
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			Appearance: mac.NSAppearanceNameAqua,
			About: &mac.AboutInfo{
				Title: "HeyMe",
				Message: strings.Join([]string{
					"© 2025 Tachera Sasi",
				}, "\n"),
			},
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    true,
			DisablePinchZoom:     true,
			BackdropType:         windows.Acrylic,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 250, G: 249, B: 238, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

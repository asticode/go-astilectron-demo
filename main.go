package main

import (
	"flag"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron/bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Flags
var (
	debug = flag.Bool("d", false, "if yes, the app is in debug mode")
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            "Demo",
			AppIconDarwinPath:  "resources/gopher.icns",
			AppIconDefaultPath: "resources/gopher.png",
		},
		Debug:          *debug,
		Homepage:       "index.html",
		MessageHandler: handleMessages,
		OnWait:         onWait,
		// RestoreAssets:  RestoreAssets,
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#333"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(600),
			Width:           astilectron.PtrInt(600),
		},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}

func onWait(a *astilectron.Astilectron, w *astilectron.Window) (err error) {
	// Init menu
	m := a.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astilectron.PtrStr("Demo"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Role: astilectron.MenuItemRoleClose,
				},
			},
		},
		{
			Label: astilectron.PtrStr("Style"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Checked: astilectron.PtrBool(true),
					Label:   astilectron.PtrStr("Dark"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						// Send
						if err = w.Send(bootstrap.MessageOut{Name: "set.style", Payload: "dark"}); err != nil {
							astilog.Error(errors.Wrap(err, "setting dark style failed"))
							return
						}
						return
					},
					Type: astilectron.MenuItemTypeRadio,
				},
				{
					Label: astilectron.PtrStr("Light"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						// Send
						if err = w.Send(bootstrap.MessageOut{Name: "set.style", Payload: "light"}); err != nil {
							astilog.Error(errors.Wrap(err, "setting dark style failed"))
							return
						}
						return
					},
					Type: astilectron.MenuItemTypeRadio,
				},
			},
		},
	})

	// Create menu
	if err = m.Create(); err != nil {
		err = errors.Wrap(err, "creating menu failed")
		return
	}
	return
}

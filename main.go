package main

import (
	"flag"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "if yes, the app is in debug mode")
	window  *astilectron.Window
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{AppName: AppName},
		Debug:              *debug,
		Homepage:           "index.html",
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astilectron.PtrStr(AppName),
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
							if err := window.Send(bootstrap.MessageOut{Name: "set.style", Payload: "dark"}); err != nil {
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
							if err := window.Send(bootstrap.MessageOut{Name: "set.style", Payload: "light"}); err != nil {
								astilog.Error(errors.Wrap(err, "setting dark style failed"))
								return
							}
							return
						},
						Type: astilectron.MenuItemTypeRadio,
					},
				},
			},
		},
		MessageHandler: handleMessages,
		OnWait: func(_ *astilectron.Astilectron, w *astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			window = w
			return nil
		},
		RestoreAssets: RestoreAssets,
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

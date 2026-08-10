package main

import (
	"log"

	"mme/state"
	"mme/ui"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		th := material.NewTheme()
		store := state.New()

		var ops op.Ops

		for {
			switch e := w.Event().(type) {

			case app.DestroyEvent:
				log.Fatal(e.Err)

			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				ui.RootLayout(gtx, th, store, w)

				e.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}

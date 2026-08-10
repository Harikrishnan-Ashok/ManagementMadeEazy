package ui

import (
	"strconv"

	"mme/state"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func RootLayout(
	gtx layout.Context,
	th *material.Theme,
	store *state.UIState,
	w *app.Window,
) layout.Dimensions {
	for store.CloseBtn.Clicked(gtx) {
		store.Num++
	}

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Button(th, &store.CloseBtn, strconv.Itoa(store.Num)).Layout(gtx)
	})
}


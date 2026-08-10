package state

import "gioui.org/widget"

type TestButton struct {
	CloseBtn widget.Clickable
	Num      int
}

type UIState struct {
	TestButton
}

func New() *UIState {
	return &UIState{}
}

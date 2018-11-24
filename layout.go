package main

import (
	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

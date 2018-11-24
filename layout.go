package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	width, height := g.Size()

	if v, err := g.SetView("treeview", 2, 1, width/4, height-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}

	if v, err := g.SetView("filelist", 2+width/4, 1, width-2, height-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	// Quit
	return gocui.ErrQuit
}

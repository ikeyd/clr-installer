// Copyright © 2018 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/gotk3/gotk3/gtk"
)

// Window provides management of the underlying GtkWindow and
// associated windows to provide a level of OOP abstraction.
type Window struct {
	handle *gtk.Window  // Abstract the underlying GtkWindow
}

// New creates a new Window instance
func NewWindow() (*Window, error) {
	window := &Window{}

	// Construct main window
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	window.handle = win

	// Headerbar for visual consistency
	hbar, err := gtk.HeaderBarNew()
	hbar.SetShowCloseButton(true)
	window.handle.SetTitlebar(hbar)

	// Set up basic window attributes
	window.handle.SetTitle("Install Clear Linux")
	window.handle.SetPosition(gtk.WIN_POS_CENTER)
	window.handle.SetDefaultSize(800, 600)

	// Temporary for development testing: Close window when asked
	window.handle.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Show it
	window.handle.ShowAll()

	return window, nil
}

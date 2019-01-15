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
	handle *gtk.Window     // Abstract the underlying GtkWindow
	header *gtk.HeaderBar  // Headerbar for navigation
	stack  *gtk.Stack      // Hold all of our pages
	top    *gtk.Box        // Top box for the main labels
	layout *gtk.Box        // Main layout (vertical)
}

// New creates a new Window instance
func NewWindow() (*Window, error) {
	window := &Window{}
	var err error

	// Construct main window
	window.handle, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	// Headerbar for visual consistency
	window.header, err = gtk.HeaderBarNew()
	if err != nil {
		return nil, err
	}
	window.header.SetShowCloseButton(true)
	window.handle.SetTitlebar(window.header)

	// Set up basic window attributes
	window.handle.SetTitle("Install Clear Linux")
	window.handle.SetPosition(gtk.WIN_POS_CENTER)
	window.handle.SetDefaultSize(800, 600)
	// Temporary icon: Need .desktop file + icon asset
	window.handle.SetIconName("system-software-install")

	// Set up the main layout
	window.layout, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	window.handle.Add(window.layout)

	// Set up the top box
	window.top, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	window.layout.PackStart(window.top, false, false, 0)

	// Set up the content stack
	window.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}
	window.layout.PackStart(window.stack, true, true, 0)

	// Temporary for development testing: Close window when asked
	window.handle.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Remove in future
	window.UglyDemoCode()

	// Show it
	window.handle.ShowAll()

	return window, nil
}

// AddPage will add the given page to this window
func (win *Window) AddPage(page *Page) {
}

func (window *Window) UglyDemoCode() {
	button, _ := gtk.ButtonNewWithLabel("Required")
	button.SetRelief(gtk.RELIEF_NONE)
	window.top.PackStart(button, false, false, 0)

	button, _ = gtk.ButtonNewWithLabel("Advanced")
	button.SetRelief(gtk.RELIEF_NONE)
	window.top.PackStart(button, false, false, 0)
}

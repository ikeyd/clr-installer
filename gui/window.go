// Copyright © 2018 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/gtk"
)

// Window provides management of the underlying GtkWindow and
// associated windows to provide a level of OOP abstraction.
type Window struct {
	handle   *gtk.Window        // Abstract the underlying GtkWindow
	header   *gtk.HeaderBar     // Headerbar for navigation
	stack    *gtk.Stack         // Hold primary switcher content
	switcher *gtk.StackSwitcher // Allow switching between main components
	top      *gtk.Box           // Top box for the main labels
	layout   *gtk.Box           // Main layout (vertical)

	screens map[bool]*ContentView // Mapping to content views
}

// ConstructHeaderBar attempts creation of the headerbar
func (window *Window) ConstructHeaderBar() error {
	var err error

	// Headerbar for visual consistency
	window.header, err = gtk.HeaderBarNew()
	if err != nil {
		return err
	}

	window.header.SetShowCloseButton(true)
	window.handle.SetTitlebar(window.header)

	return nil
}

// NewWindow creates a new Window instance
func NewWindow() (*Window, error) {
	window := &Window{}
	var err error

	// Set up screen mapping
	window.screens = make(map[bool]*ContentView)

	// Construct main window
	window.handle, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	// Need HeaderBar ?
	if err = window.ConstructHeaderBar(); err != nil {
		return nil, err
	}

	// Set up basic window attributes
	window.handle.SetTitle("Clear Linux* OS Installer [" + model.Version + "]")
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

	// Set up the stack switcher
	window.switcher, err = gtk.StackSwitcherNew()
	if err != nil {
		return nil, err
	}
	window.layout.PackStart(window.switcher, false, false, 0)
	window.switcher.SetHAlign(gtk.ALIGN_CENTER)

	// Set up the content stack
	window.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}
	window.layout.PackStart(window.stack, true, true, 0)
	window.switcher.SetStack(window.stack)

	// Temporary for development testing: Close window when asked
	window.handle.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Set up primary content views
	if err = window.InitScreens(); err != nil {
		return nil, err
	}

	// Remove in future
	window.UglyDemoCode()
	window.handle.SetBorderWidth(4)

	// Show it
	window.handle.ShowAll()

	return window, nil
}

// InitScreens will set up the content views
func (window *Window) InitScreens() error {
	var err error

	// Set up required screen
	if window.screens[true], err = NewContentView(); err != nil {
		return err
	}
	window.stack.AddTitled(window.screens[ContentViewRequired].GetRootWidget(), "required", "Required options")

	// Set up non required screen
	if window.screens[false], err = NewContentView(); err != nil {
		return err
	}
	window.stack.AddTitled(window.screens[ContentViewAdvanced].GetRootWidget(), "advanced", "Advanced options")

	return nil
}

// AddPage will add the given page to this window
func (win *Window) AddPage(page *Page) {
}

func (window *Window) UglyDemoCode() {
	// Set up nav buttons
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	window.layout.PackEnd(box, false, false, 0)
	box.SetHAlign(gtk.ALIGN_END)

	button, _ := gtk.ButtonNewWithLabel("Cancel")
	box.PackStart(button, false, false, 2)

	button, _ = gtk.ButtonNewWithLabel("Install")
	st, _ := button.GetStyleContext()
	st.AddClass("suggested-action")
	box.PackStart(button, false, false, 2)
}

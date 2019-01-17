// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/clearlinux/clr-installer/gui/pages"
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// PageConstructor is a typedef of the constructors for our pages
type PageConstructor func() (pages.Page, error)

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
	window.handle.SetResizable(false)
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
	window.top.SetMarginTop(24)
	window.top.SetMarginBottom(24)
	window.layout.PackStart(window.top, false, false, 0)

	// Set up the stack switcher
	window.switcher, err = gtk.StackSwitcherNew()
	if err != nil {
		return nil, err
	}

	// Stick the switcher into the headerbar
	window.header.SetCustomTitle(window.switcher)

	// Create the header
	window.CreateHeader()

	// Set up the content stack
	window.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}
	window.stack.SetTransitionType(gtk.STACK_TRANSITION_TYPE_CROSSFADE)
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

	// Create footer area now
	window.CreateFooter()

	// Our pages
	pageCreators := []PageConstructor{
		pages.NewTimezonePage,
		pages.NewLanguagePage,
		pages.NewKeyboardPage,
		pages.NewBundlePage,
	}

	// Create all pages
	for _, f := range pageCreators {
		page, err := f()
		if err != nil {
			return nil, err
		}
		window.AddPage(page)
	}

	// Show the whole window now
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

// AddPage will add the page to the relevant screen
func (window *Window) AddPage(page pages.Page) {
	window.screens[page.IsRequired()].AddPage(page)
}

// MakeHeader constructs the header component
func (window *Window) CreateHeader() {
	img, _ := gtk.ImageNew()
	filePath := "themes/clr.png"
	pbuf, _ := gdk.PixbufNewFromFileAtSize(filePath, 128, 128)
	img.SetFromPixbuf(pbuf)
	img.SetPixelSize(64)
	img.SetHAlign(gtk.ALIGN_START)
	window.top.PackStart(img, false, false, 0)
	window.top.SetHAlign(gtk.ALIGN_CENTER)

	label, _ := gtk.LabelNew("<span font-size='xx-large'>Install Clear Linux* OS</span>\n\nTODO: Insert awesome header widget in this general region.\nKinda show off why this is an awesome decision.")
	label.SetUseMarkup(true)
	label.SetMarginStart(40)
	label.SetHAlign(gtk.ALIGN_START)
	label.SetVAlign(gtk.ALIGN_CENTER)
	window.top.PackStart(label, true, true, 0)
}

func (window *Window) CreateFooter() {
	// Store components
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	box.SetMarginTop(4)
	box.SetMarginBottom(6)
	box.SetMarginEnd(6)
	box.SetMarginStart(6)
	window.layout.PackEnd(box, false, false, 0)
	box.SetHAlign(gtk.ALIGN_FILL)

	// Version label
	label, _ := gtk.LabelNew("Clear Linux* OS Installer [" + model.Version + "]")
	label.SetHAlign(gtk.ALIGN_START)
	st, _ := label.GetStyleContext()
	st.AddClass("dim-label")
	box.PackStart(label, false, false, 0)

	// Set up nav buttons
	button, _ := gtk.ButtonNewWithLabel("Install")
	button.SetHAlign(gtk.ALIGN_END)
	button.SetRelief(gtk.RELIEF_NONE)
	st, _ = button.GetStyleContext()
	st.AddClass("suggested-action")
	box.PackEnd(button, false, false, 2)

	button, _ = gtk.ButtonNewWithLabel("Cancel")
	button.SetHAlign(gtk.ALIGN_END)
	button.SetRelief(gtk.RELIEF_NONE)
	box.PackEnd(button, false, false, 2)
}

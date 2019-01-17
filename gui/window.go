// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"fmt"
	"github.com/clearlinux/clr-installer/gui/pages"
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// PageConstructor is a typedef of the constructors for our pages
type PageConstructor func() (pages.Page, error)

// Window provides management of the underlying GtkWindow and
// associated windows to provide a level of OOP abstraction.
type Window struct {
	handle    *gtk.Window        // Abstract the underlying GtkWindow
	header    *gtk.HeaderBar     // Headerbar for navigation
	stack     *gtk.Stack         // Hold primary switcher content
	rootStack *gtk.Stack         // Root-level stack
	switcher  *gtk.StackSwitcher // Allow switching between main components
	layout    *gtk.Box           // Main layout (vertical)
	banner    *Banner            // Top banner

	didInit bool // Whether we've inited the view animation

	screens map[bool]*ContentView // Mapping to content views
	pages   map[int]gtk.IWidget   // Mapping to each root page
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
	var err error

	// Construct basic window
	window := &Window{
		didInit: false,
		screens: make(map[bool]*ContentView),
		pages:   make(map[int]gtk.IWidget),
	}

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

	// Create the banner
	if window.banner, err = NewBanner(); err != nil {
		return nil, err
	}
	window.layout.PackStart(window.banner.GetRootWidget(), false, false, 0)

	// Set up the stack switcher
	window.switcher, err = gtk.StackSwitcherNew()
	if err != nil {
		return nil, err
	}

	// Stick the switcher into the headerbar
	window.header.SetCustomTitle(window.switcher)

	// Set up the root stack
	window.rootStack, err = gtk.StackNew()
	window.rootStack.SetTransitionType(gtk.STACK_TRANSITION_TYPE_CROSSFADE)
	if err != nil {
		return nil, err
	}
	window.layout.PackStart(window.rootStack, true, true, 0)

	// Set up the content stack
	window.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}
	window.stack.SetTransitionType(gtk.STACK_TRANSITION_TYPE_CROSSFADE)
	window.switcher.SetStack(window.stack)

	// Add menu stack to root stack
	window.rootStack.AddTitled(window.stack, "menu", "Menu")

	// Temporary for development testing: Close window when asked
	window.handle.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// On map, expose the revealer
	window.handle.Connect("map", window.handleMap)

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
	if window.screens[true], err = NewContentView(window); err != nil {
		return err
	}
	window.stack.AddTitled(window.screens[ContentViewRequired].GetRootWidget(), "required", "Required options")

	// Set up non required screen
	if window.screens[false], err = NewContentView(window); err != nil {
		return err
	}
	window.stack.AddTitled(window.screens[ContentViewAdvanced].GetRootWidget(), "advanced", "Advanced options")

	return nil
}

// AddPage will add the page to the relevant screen
func (window *Window) AddPage(page pages.Page) {
	id := page.GetID()

	// Add to the required or advanced(optional) screen
	window.screens[page.IsRequired()].AddPage(page)

	// Store root widget too
	root := page.GetRootWidget()
	window.pages[id] = root
	if root != nil {
		window.rootStack.AddNamed(root, "page:"+string(id))
	}
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

// We've been mapped on screen
func (window *Window) handleMap() {
	if window.didInit {
		return
	}
	glib.TimeoutAdd(200, func() bool {
		if !window.didInit {
			window.banner.ShowFirst()
			window.stack.SetVisibleChildName("required")
			window.didInit = true
		}
		return false
	})
}

// ActivatePage will set the view as visible.
func (window *Window) ActivatePage(page pages.Page) {
	fmt.Println("Activating: " + page.GetTitle())

	// Hide banner so we can get more room
	window.banner.Hide()

	id := page.GetID()
	root := window.pages[id]
	if root != nil {
		// Set the root stack to show the new page
		window.rootStack.SetVisibleChild(window.pages[id])
	}

}

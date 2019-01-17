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
	handle        *gtk.Window // Abstract the underlying GtkWindow
	rootStack     *gtk.Stack  // Root-level stack
	layout        *gtk.Box    // Main layout (vertical)
	contentLayout *gtk.Box    // content layout (horizontal)
	banner        *Banner     // Top banner

	// Wrap titlebar access
	title struct {
		bar   *gtk.HeaderBar // Headerbar for navigation
		stack *gtk.Stack     // Allow switching between title and switcher
		label *gtk.Label     // Label to display in title mode
	}

	// Menus
	menu struct {
		stack    *gtk.Stack            // Menu switching
		switcher *gtk.StackSwitcher    // Allow switching between main menu
		screens  map[bool]*ContentView // Mapping to content views
	}

	didInit bool // Whether we've inited the view animation

	pages map[int]gtk.IWidget // Mapping to each root page
}

// ConstructHeaderBar attempts creation of the headerbar
func (window *Window) ConstructHeaderBar() error {
	var err error

	// Headerbar for visual consistency
	window.title.bar, err = gtk.HeaderBarNew()
	if err != nil {
		return err
	}

	// Set up titlebar
	window.title.bar.SetShowCloseButton(true)
	window.handle.SetTitlebar(window.title.bar)

	// Set up stack
	window.title.stack, err = gtk.StackNew()
	if err != nil {
		return err
	}
	window.title.stack.SetTransitionType(gtk.STACK_TRANSITION_TYPE_CROSSFADE)
	window.title.bar.SetCustomTitle(window.title.stack)

	// Add custom label
	window.title.label, err = gtk.LabelNew("")
	if err != nil {
		return err
	}

	// Set up title styling
	st, err := window.title.label.GetStyleContext()
	if err != nil {
		return err
	}
	st.AddClass("title")
	window.title.label.ShowAll()

	window.title.stack.AddNamed(window.title.label, "title")

	return nil
}

// NewWindow creates a new Window instance
func NewWindow() (*Window, error) {
	var err error

	// Construct basic window
	window := &Window{
		didInit: false,
		pages:   make(map[int]gtk.IWidget),
	}
	window.menu.screens = make(map[bool]*ContentView)

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
	window.handle.SetDefaultSize(800, 500)
	window.handle.SetResizable(false)
	// Temporary icon: Need .desktop file + icon asset
	window.handle.SetIconName("system-software-install")

	// Set up the main layout
	window.layout, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	window.handle.Add(window.layout)

	// To add the *main* content
	window.contentLayout, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	window.layout.PackStart(window.contentLayout, true, true, 0)

	// Create the banner
	if window.banner, err = NewBanner(); err != nil {
		return nil, err
	}
	window.contentLayout.PackStart(window.banner.GetRootWidget(), false, false, 0)

	// Set up the stack switcher
	window.menu.switcher, err = gtk.StackSwitcherNew()
	if err != nil {
		return nil, err
	}

	// Stick the switcher into the headerbar
	window.title.stack.AddNamed(window.menu.switcher, "switcher")

	// Set up the root stack
	window.rootStack, err = gtk.StackNew()
	window.rootStack.SetTransitionType(gtk.STACK_TRANSITION_TYPE_CROSSFADE)
	if err != nil {
		return nil, err
	}
	window.contentLayout.PackStart(window.rootStack, true, true, 0)

	// Set up the content stack
	window.menu.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}
	window.menu.stack.SetTransitionType(gtk.STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT)
	window.menu.switcher.SetStack(window.menu.stack)

	// Add menu stack to root stack
	window.rootStack.AddTitled(window.menu.stack, "menu", "Menu")

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

	// Show the default view
	window.ShowDefaultView()

	return window, nil
}

func (window *Window) ShowDefaultView() {
	// Ensure switcher is the guy showing now
	window.title.stack.SetVisibleChildName("switcher")

	// Ensure menu page is set
	window.menu.stack.SetVisibleChildName("required")

	// And root stack
	window.rootStack.SetVisibleChildName("menu")
}

// InitScreens will set up the content views
func (window *Window) InitScreens() error {
	var err error

	// Set up required screen
	if window.menu.screens[true], err = NewContentView(window); err != nil {
		return err
	}
	window.menu.stack.AddTitled(window.menu.screens[ContentViewRequired].GetRootWidget(), "required", "Required options")

	// Set up non required screen
	if window.menu.screens[false], err = NewContentView(window); err != nil {
		return err
	}
	window.menu.stack.AddTitled(window.menu.screens[ContentViewAdvanced].GetRootWidget(), "advanced", "Advanced options")

	return nil
}

// AddPage will add the page to the relevant screen
func (window *Window) AddPage(page pages.Page) {
	id := page.GetID()

	// Add to the required or advanced(optional) screen
	window.menu.screens[page.IsRequired()].AddPage(page)

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

	// Set up nav buttons
	button, _ := gtk.ButtonNewWithLabel("Install")
	button.SetHAlign(gtk.ALIGN_END)
	button.SetRelief(gtk.RELIEF_NONE)
	st, _ := button.GetStyleContext()
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
			window.menu.stack.SetVisibleChildName("required")
			window.didInit = true
		}
		return false
	})
}

// ActivatePage will set the view as visible.
func (window *Window) ActivatePage(page pages.Page) {
	fmt.Println("Activating: " + page.GetSummary())

	// Hide banner so we can get more room
	window.banner.Hide()

	id := page.GetID()
	root := window.pages[id]
	if root != nil {
		// Set the root stack to show the new page
		window.rootStack.SetVisibleChild(window.pages[id])
	}

	// Update title for display
	window.title.label.SetLabel(page.GetSummary())
	window.title.stack.SetVisibleChildName("title")
}

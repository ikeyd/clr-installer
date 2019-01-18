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

	// Menus
	menu struct {
		stack    *gtk.Stack            // Menu switching
		switcher *Switcher             // Allow switching between main menu
		screens  map[bool]*ContentView // Mapping to content views
	}

	// Buttons
	buttons struct {
		stack        *gtk.Stack // Storage for buttons
		boxPrimary   *gtk.Box   // Storage for main buttons (install/quit)
		boxSecondary *gtk.Box   // Storage for secondary buttons (confirm/cancel)

		confirm *gtk.Button // Apply changes
		cancel  *gtk.Button // Cancel changes
		install *gtk.Button // Install Clear Linux
		quit    *gtk.Button // Quit the installer
	}

	didInit bool // Whether we've inited the view animation

	pages map[int]gtk.IWidget // Mapping to each root page
}

// ConstructHeaderBar attempts creation of the headerbar
func (window *Window) ConstructHeaderBar() error {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return err
	}

	st, err := box.GetStyleContext()
	if err != nil {
		return err
	}

	window.handle.SetTitlebar(box)
	st.RemoveClass("titlebar")
	st.RemoveClass("headerbar")
	st.RemoveClass("header")
	st.AddClass("invisible-titlebar")

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

	// Set up the stack switcher
	window.menu.switcher, err = NewSwitcher(nil)
	if err != nil {
		return nil, err
	}

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

	// Set up the root stack
	window.rootStack, err = gtk.StackNew()
	window.rootStack.SetTransitionType(gtk.STACK_TRANSITION_TYPE_CROSSFADE)
	if err != nil {
		return nil, err
	}

	// We want vertical layout here with buttons above the rootstack
	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	window.contentLayout.PackStart(vbox, true, true, 0)
	vbox.PackStart(window.menu.switcher.GetRootWidget(), false, false, 0)
	vbox.PackStart(window.rootStack, true, true, 0)

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
	if err = window.CreateFooter(vbox); err != nil {
		return nil, err
	}

	// Our pages
	pageCreators := []PageConstructor{
		// required
		pages.NewTimezonePage,
		pages.NewLanguagePage,
		pages.NewKeyboardPage,
		pages.NewDiskConfigPage,
		pages.NewTelemetryPage,

		// advanced
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
	window.menu.stack.AddTitled(window.menu.screens[ContentViewRequired].GetRootWidget(), "required", "REQUIRED OPTIONS\nTakes approximately 2 minutes")

	// Set up non required screen
	if window.menu.screens[false], err = NewContentView(window); err != nil {
		return err
	}
	window.menu.stack.AddTitled(window.menu.screens[ContentViewAdvanced].GetRootWidget(), "advanced", "ADVANCED OPTIONS\nCustomize setup")

	return nil
}

// AddPage will add the page to the relevant screen
func (window *Window) AddPage(page pages.Page) {
	id := page.GetID()

	// Add to the required or advanced(optional) screen
	window.menu.screens[page.IsRequired()].AddPage(page)

	// Store root widget too
	root := page.GetRootWidget()
	if root == nil {
		window.pages[id] = root
		return
	}

	ebox, _ := gtk.EventBoxNew()
	st, _ := ebox.GetStyleContext()
	st.AddClass("installer-header-box")
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	box2, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	box2.SetBorderWidth(6)
	ebox.Add(box2)
	box.PackStart(ebox, false, false, 0)
	img, _ := gtk.ImageNewFromIconName(page.GetIcon()+"-symbolic", gtk.ICON_SIZE_INVALID)
	img.SetPixelSize(48)
	img.SetMarginStart(6)
	img.SetMarginEnd(12)
	img.SetMarginTop(4)
	img.SetMarginBottom(4)
	box2.PackStart(img, false, false, 0)

	lab, _ := gtk.LabelNew("<big>" + page.GetTitle() + "</big>")
	lab.SetUseMarkup(true)
	box2.PackStart(lab, false, false, 0)
	box.ShowAll()
	ebox.SetMarginBottom(6)

	box.PackStart(root, true, true, 0)
	window.pages[id] = box
	window.rootStack.AddNamed(box, "page:"+string(id))
}

// createNavButton creates specialised navigation button
func createNavButton(label string) (*gtk.Button, error) {
	var st *gtk.StyleContext
	button, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		return nil, err
	}

	st, err = button.GetStyleContext()
	if err != nil {
		return nil, err
	}
	st.AddClass("nav-button")
	return button, nil
}

// CreateFooter creates our navigation footer area
func (window *Window) CreateFooter(store *gtk.Box) error {
	var err error

	// Create stack for buttons
	if window.buttons.stack, err = gtk.StackNew(); err != nil {
		return err
	}

	// Set alignment up
	window.buttons.stack.SetMarginTop(4)
	window.buttons.stack.SetMarginBottom(6)
	window.buttons.stack.SetMarginEnd(24)
	window.buttons.stack.SetHAlign(gtk.ALIGN_END)
	window.buttons.stack.SetTransitionType(gtk.STACK_TRANSITION_TYPE_CROSSFADE)
	store.PackEnd(window.buttons.stack, false, false, 0)

	// Create box for primary buttons
	if window.buttons.boxPrimary, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0); err != nil {
		return err
	}

	// Install button
	if window.buttons.install, err = createNavButton("INSTALL"); err != nil {
		return err
	}

	// Exit button
	if window.buttons.quit, err = createNavButton("EXIT"); err != nil {
		return err
	}
	window.buttons.quit.Connect("clicked", func() {
		gtk.MainQuit()
	})

	// Create box for secondary buttons
	if window.buttons.boxSecondary, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0); err != nil {
		return err
	}

	// Confirm button
	if window.buttons.confirm, err = createNavButton("CONFIRM"); err != nil {
		return err
	}

	// Cancel button
	if window.buttons.cancel, err = createNavButton("CANCEL"); err != nil {
		return err
	}

	// Pack the buttons
	window.buttons.boxPrimary.PackEnd(window.buttons.install, false, false, 4)
	window.buttons.boxPrimary.PackEnd(window.buttons.quit, false, false, 4)
	window.buttons.boxSecondary.PackEnd(window.buttons.confirm, false, false, 4)
	window.buttons.boxSecondary.PackEnd(window.buttons.cancel, false, false, 4)

	// Add the boxes
	window.buttons.stack.AddNamed(window.buttons.boxPrimary, "primary")
	window.buttons.stack.AddNamed(window.buttons.boxSecondary, "secondary")

	return nil
}

// We've been mapped on screen
func (window *Window) handleMap() {
	if window.didInit {
		return
	}
	glib.TimeoutAdd(200, func() bool {
		if !window.didInit {
			window.banner.ShowFirst()
			window.menu.switcher.Show()
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
	window.menu.switcher.Hide()

	// Show secondary controls
	window.buttons.stack.SetVisibleChildName("secondary")

	id := page.GetID()
	root := window.pages[id]
	if root != nil {
		// Set the root stack to show the new page
		window.rootStack.SetVisibleChild(window.pages[id])
	}
}

// SetButtonState is called by the pages to enable/disable certain buttons.
func (window *Window) SetButtonState(flags pages.Button, enabled bool) {
	if flags&pages.ButtonCancel == pages.ButtonCancel {
		window.buttons.cancel.SetSensitive(enabled)
	}
	if flags&pages.ButtonConfirm == pages.ButtonConfirm {
		window.buttons.confirm.SetSensitive(enabled)
	}
}

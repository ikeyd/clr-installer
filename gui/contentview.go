// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/clearlinux/clr-installer/gui/pages"
	"github.com/gotk3/gotk3/gtk"
)

// ContentView is used to encapsulate the Required/Advanced options view
// by wrapping them into simple styled lists
type ContentView struct {
	views      map[int]pages.Page // Mapping of row to page
	controller pages.Controller

	scroll *gtk.ScrolledWindow
	list   *gtk.ListBox
}

// NewContentView will attempt creation of a new ContentView
func NewContentView(controller pages.Controller) (*ContentView, error) {
	var err error

	// Init the struct
	view := &ContentView{
		controller: controller,
		views:      make(map[int]pages.Page),
	}

	// Set up the scroller
	if view.scroll, err = gtk.ScrolledWindowNew(nil, nil); err != nil {
		return nil, err
	}

	// Set the scroll policy
	view.scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)

	// Set shadow type
	view.scroll.SetShadowType(gtk.SHADOW_NONE)

	// Set up the list
	if view.list, err = gtk.ListBoxNew(); err != nil {
		return nil, err
	}

	// Ensure navigation works properly
	view.list.SetSelectionMode(gtk.SELECTION_SINGLE)
	view.list.SetActivateOnSingleClick(true)
	view.list.Connect("row-activated", view.onRowActivated)
	view.scroll.Add(view.list)

	return view, nil
}

// GetRootWidget will return the root widget for embedding
func (view *ContentView) GetRootWidget() gtk.IWidget {
	return view.scroll
}

// AddPage will add the relevant page to this content view.
// Right now it does nothing.
func (view *ContentView) AddPage(page pages.Page) {
	// TESTING CODE ONLY!
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	box.SetHAlign(gtk.ALIGN_START)
	box.SetMarginStart(18)
	box.SetMarginEnd(18)
	box.SetMarginTop(6)
	box.SetMarginBottom(6)

	// image
	img, _ := gtk.ImageNewFromIconName(page.GetIcon(), gtk.ICON_SIZE_DIALOG)
	img.SetMarginEnd(6)
	box.PackStart(img, false, false, 0)

	// label
	wid, _ := gtk.LabelNew("<big>" + page.GetTitle() + "</big>")
	wid.SetUseMarkup(true)
	wid.SetHAlign(gtk.ALIGN_START)
	box.PackStart(wid, false, false, 0)

	wrap, _ := gtk.ListBoxRowNew()
	wrap.Add(box)
	wrap.ShowAll()
	view.list.Add(wrap)
	view.views[wrap.GetIndex()] = page
}

func (view *ContentView) onRowActivated(box *gtk.ListBox, row *gtk.ListBoxRow) {
	if row == nil {
		return
	}
	// Go activate this.
	view.controller.ActivatePage(view.views[row.GetIndex()])
}

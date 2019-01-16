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
	scroll *gtk.ScrolledWindow
	list   *gtk.ListBox
}

// NewContentView will attempt creation of a new ContentView
func NewContentView() (*ContentView, error) {
	var err error

	// Init the struct
	view := &ContentView{}

	// Set up the scroller
	if view.scroll, err = gtk.ScrolledWindowNew(nil, nil); err != nil {
		return nil, err
	}

	// Set the scroll policy
	view.scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)

	// Set shadow type
	view.scroll.SetShadowType(gtk.SHADOW_ETCHED_IN)

	// Set up the list
	if view.list, err = gtk.ListBoxNew(); err != nil {
		return nil, err
	}
	view.list.SetSelectionMode(gtk.SELECTION_NONE)
	view.scroll.Add(view.list)

	return view, nil
}

// GetRootWidget will return the root widget for embedding
func (view *ContentView) GetRootWidget() *gtk.ScrolledWindow {
	return view.scroll
}

// AddPage will add the relevant page to this content view.
// Right now it does nothing.
func (view *ContentView) AddPage(page pages.Page) {
	// TESTING CODE ONLY!
	wid, _ := gtk.LabelNew(page.GetTitle())
	wid.SetHAlign(gtk.ALIGN_START)
	wid.SetMarginStart(18)
	wid.SetMarginEnd(18)
	wid.SetMarginTop(6)
	wid.SetMarginBottom(6)
	view.list.Add(wid)
	wid.ShowAll()
}

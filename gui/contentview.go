// Copyright © 2018 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
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

	// Set up the list
	if view.list, err = gtk.ListBoxNew(); err != nil {
		return nil, err
	}

	return view, nil
}

// GetRootWidget will return the root widget for embedding
func (view *ContentView) GetRootWidget() (*gtk.ScrolledWindow) {
	return view.scroll
}

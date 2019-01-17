// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/gotk3/gotk3/gtk"
)

// Switcher is used to switch between main installer sections
type Switcher struct {
	box      *gtk.Box    // Main layout
	stack    *gtk.Stack  // Stack to control
}

// MakeHeader constructs the header component
func NewSwitcher(stack *gtk.Stack) (*Switcher, error) {
	var err error
	var st *gtk.StyleContext

	// Create switcher
	switcher := &Switcher{
		stack: stack,
	}

	// Create main layout
	switcher.box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		return nil, err
	}

	// Set styling
	st, err = switcher.box.GetStyleContext()
	if err != nil {
		return nil, err
	}
	st.AddClass("installer-switcher")
	st.AddClass("linked")

	// Required options
	button, err = gtk.ButtonNewWithLabel("REQUIRED OPTIONS")
	if err != nil {
		return nil, err
	}
	switcher.box.PackStart(button, true, true, 0)

	// Advanced options
	button, err = gtk.ButtonNewWithLabel("ADVANCED OPTIONS")
	if err != nil {
		return nil, err
	}
	switcher.box.PackStart(button, true, true, 0)

	return switcher, nil
}

// GetRootWidget returns the embeddable root widget
func (Switcher *Switcher) GetRootWidget() gtk.IWidget {
	return switcher.box
}


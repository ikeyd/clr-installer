// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/gotk3/gotk3/gtk"
)

// Switcher is used to switch between main installer sections
type Switcher struct {
	box   *gtk.Box   // Main layout
	stack *gtk.Stack // Stack to control
}

// MakeHeader constructs the header component
func NewSwitcher(stack *gtk.Stack) (*Switcher, error) {
	var err error
	var st *gtk.StyleContext
	var button *gtk.Button

	// Create switcher
	switcher := &Switcher{
		stack: stack,
	}

	// Create main layout
	switcher.box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
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
	button, err = createFancyButton("<b>REQUIRED OPTIONS</b>\n<small>Takes approximately 2 minutes</small>")
	if err != nil {
		return nil, err
	}
	button.Connect("clicked", func() {
		if switcher.stack == nil {
			return
		}
		switcher.stack.SetVisibleChildName("required")
	})

	switcher.box.PackStart(button, true, true, 0)

	// Advanced options
	button, err = createFancyButton("<b>ADVANCED OPTIONS</b>\n<small>Customize setup</small>")
	if err != nil {
		return nil, err
	}
	button.Connect("clicked", func() {
		if switcher.stack == nil {
			return
		}
		switcher.stack.SetVisibleChildName("advanced")
	})
	switcher.box.PackStart(button, true, true, 0)

	return switcher, nil
}

func createFancyButton(text string) (*gtk.Button, error) {
	button, err := gtk.ButtonNew()
	if err != nil {
		return nil, err
	}
	label, err := gtk.LabelNew(text)
	if err != nil {
		return nil, err
	}
	label.SetUseMarkup(true)
	button.Add(label)
	return button, nil
}

// GetRootWidget returns the embeddable root widget
func (switcher *Switcher) GetRootWidget() gtk.IWidget {
	return switcher.box
}

// SetStack updates the associated stack
func (switcher *Switcher) SetStack(stack *gtk.Stack) {
	switcher.stack = stack
}

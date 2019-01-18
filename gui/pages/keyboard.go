// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/keyboard"
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/gtk"
)

// Keyboard is a simple page to help with Keyboard settings
type Keyboard struct {
	controller Controller
	model      *model.SystemInstall
	keymaps    []*keyboard.Keymap
	box        *gtk.Box
	scroll     *gtk.ScrolledWindow
	list       *gtk.ListBox
}

// NewKeyboardPage returns a new KeyboardPage
func NewKeyboardPage(controller Controller, model *model.SystemInstall) (Page, error) {
	keymaps, err := keyboard.LoadKeymaps()
	if err != nil {
		return nil, err
	}

	keyboard := &Keyboard{
		controller: controller,
		model:      model,
		keymaps:    keymaps,
	}

	keyboard.box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	keyboard.box.SetBorderWidth(8)

	// Build storage for listbox
	keyboard.scroll, err = gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	keyboard.box.PackStart(keyboard.scroll, true, true, 0)
	keyboard.scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)

	// Build listbox
	keyboard.list, err = gtk.ListBoxNew()
	if err != nil {
		return nil, err
	}
	keyboard.list.SetSelectionMode(gtk.SELECTION_SINGLE)
	keyboard.list.SetActivateOnSingleClick(true)
	// keyboard.list.Connect("row-activated", keyboard.onRowActivated)
	keyboard.scroll.Add(keyboard.list)
	// Remove background
	st, _ := keyboard.list.GetStyleContext()
	st.AddClass("scroller-special")

	return keyboard, nil
}

// IsRequired will return true as we always need a Keyboard
func (t *Keyboard) IsRequired() bool {
	return true
}

func (t *Keyboard) GetID() int {
	return PageIDKeyboard
}

func (t *Keyboard) GetIcon() string {
	return "preferences-desktop-keyboard-shortcuts"
}

func (t *Keyboard) GetRootWidget() gtk.IWidget {
	return t.box
}

func (t *Keyboard) GetSummary() string {
	return "Configure the Keyboard"
}

func (t *Keyboard) GetTitle() string {
	return t.GetSummary()
}

func (t *Keyboard) StoreChanges() {}
func (t *Keyboard) ResetChanges() {}

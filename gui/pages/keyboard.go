// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/gotk3/gotk3/gtk"
)

// Keyboard is a simple page to help with Keyboard settings
type Keyboard struct {
}

// NewKeyboardPage returns a new KeyboardPage
func NewKeyboardPage() *Keyboard {
	return &Keyboard{}
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

func (t *Keyboard) GetRootWidget() *gtk.Widget {
	return nil
}

func (t *Keyboard) GetTitle() string {
	return "Configure the Keyboard"
}

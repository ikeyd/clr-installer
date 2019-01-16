// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/gotk3/gotk3/gtk"
)

// Language is a simple page to help with Language settings
type Language struct {
}

// NewLanguagePage returns a new LanguagePage
func NewLanguagePage() *Language {
	return &Language{}
}

// IsRequired will return true as we always need a Language
func (t *Language) IsRequired() bool {
	return true
}

func (t *Language) GetID() int {
	return PageIDLanguage
}

func (t *Language) GetIcon() string {
	return "preferences-desktop-locale"
}

func (t *Language) GetRootWidget() *gtk.Widget {
	return nil
}

func (t *Language) GetTitle() string {
	return "Choose Language"
}

// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/gtk"
)

// Language is a simple page to help with Language settings
type Language struct {
	controller Controller
	model      *model.SystemInstall
	box        *gtk.Box
	scroll     *gtk.ScrolledWindow
	list       *gtk.ListBox
}

// NewLanguagePage returns a new LanguagePage
func NewLanguagePage(controller Controller, model *model.SystemInstall) (Page, error) {
	var err error

	language := &Language{
		controller: controller,
		model:      model,
	}

	language.box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	language.box.SetBorderWidth(8)

	// Build storage for listbox
	language.scroll, err = gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	language.box.PackStart(language.scroll, true, true, 0)
	language.scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)

	// Build listbox
	language.list, err = gtk.ListBoxNew()
	if err != nil {
		return nil, err
	}
	language.list.SetSelectionMode(gtk.SELECTION_SINGLE)
	language.list.SetActivateOnSingleClick(true)
	// language.list.Connect("row-activated", language.onRowActivated)
	language.scroll.Add(language.list)
	// Remove background
	st, _ := language.list.GetStyleContext()
	st.AddClass("scroller-special")

	return language, nil

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

func (t *Language) GetRootWidget() gtk.IWidget {
	return t.box
}

func (t *Language) GetSummary() string {
	return "Choose Language"
}

func (t *Language) GetTitle() string {
	return t.GetSummary()
}

func (t *Language) StoreChanges() {}
func (t *Language) ResetChanges() {}

// GetConfiguredValue returns our current config
func (t *Language) GetConfiguredValue() string {
	return ""
}

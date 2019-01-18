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
}

// NewLanguagePage returns a new LanguagePage
func NewLanguagePage(controller Controller, model *model.SystemInstall) (Page, error) {
	return &Language{
		controller: controller,
		model:      model,
	}, nil
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
	return nil
}

func (t *Language) GetSummary() string {
	return "Choose Language"
}

func (t *Language) GetTitle() string {
	return t.GetSummary()
}

func (t *Language) StoreChanges() {}
func (t *Language) ResetChanges() {}

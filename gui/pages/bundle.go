// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/gotk3/gotk3/gtk"
)

// Bundle is a simple page to help with Bundle settings
type Bundle struct {
}

// NewBundlePage returns a new BundlePage
func NewBundlePage() (Page, error) {
	return &Bundle{}, nil
}

// IsRequired will return false as we have default values
func (t *Bundle) IsRequired() bool {
	return false
}

func (t *Bundle) GetID() int {
	return PageIDBundle
}

func (t *Bundle) GetIcon() string {
	return "preferences-desktop-applications"
}

func (t *Bundle) GetRootWidget() gtk.IWidget {
	return nil
}

func (t *Bundle) GetTitle() string {
	return "Bundle selection"
}

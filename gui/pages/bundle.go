// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"fmt"
	"github.com/clearlinux/clr-installer/swupd"
	"github.com/gotk3/gotk3/gtk"
)

// Bundle is a simple page to help with Bundle settings
type Bundle struct {
	bundles []*swupd.Bundle     // Known bundles
	box     *gtk.Box            // Main layout
	checks  *gtk.Box            // Where to store checks
	scroll  *gtk.ScrolledWindow // Scroll the checks
}

// NewBundlePage returns a new BundlePage
func NewBundlePage() (Page, error) {
	var err error
	bundle := &Bundle{}

	// Load our bundles
	bundle.bundles, err = swupd.LoadBundleList()
	if err != nil {
		return nil, err
	}

	// main layout
	bundle.box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	bundle.box.SetBorderWidth(8)

	// check list
	bundle.checks, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	bundle.scroll, err = gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	// no horizontal scrolling
	bundle.scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	bundle.scroll.Add(bundle.checks)
	bundle.box.PackStart(bundle.scroll, true, true, 0)

	for _, b := range bundle.bundles {
		lab := fmt.Sprintf("%s - %s", b.Name, b.Desc)
		check, err := gtk.CheckButtonNewWithLabel(lab)
		if err != nil {
			return nil, err
		}
		bundle.checks.PackStart(check, false, false, 0)
	}

	return bundle, nil
}

// IsRequired will return false as we have default values
func (bundle *Bundle) IsRequired() bool {
	return false
}

func (bundle *Bundle) GetID() int {
	return PageIDBundle
}

func (bundle *Bundle) GetIcon() string {
	return "preferences-desktop-applications"
}

func (bundle *Bundle) GetRootWidget() gtk.IWidget {
	return bundle.box
}

func (bundle *Bundle) GetSummary() string {
	return "Bundle selection"
}

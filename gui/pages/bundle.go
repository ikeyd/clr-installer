// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"fmt"
	"github.com/clearlinux/clr-installer/swupd"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"os"
	"path/filepath"
)

const (
	// IconDirectory is where we can find bundle icons
	IconDirectory = "/usr/share/clear/bundle-icons"
)

var (
	// IconSuffixes is the supported set of suffixes for the
	// current Clear Bundles
	IconSuffixes = []string{
		".svg",
		".png",
	}
)

// Bundle is a simple page to help with Bundle settings
type Bundle struct {
	bundles []*swupd.Bundle     // Known bundles
	box     *gtk.Box            // Main layout
	checks  *gtk.FlowBox        // Where to store checks
	scroll  *gtk.ScrolledWindow // Scroll the checks
}

// LookupBundleIcon attempts to find the icon for the given bundle.
// If it is found, we'll return true and the icon path, otherwise
// we'll return false with an empty string.
func LookupBundleIcon(bundle *swupd.Bundle) (string, bool) {
	for _, suffix := range IconSuffixes {
		path := filepath.Join(IconDirectory, fmt.Sprintf("%s%s", bundle.Name, suffix))
		if _, err := os.Stat(path); err == nil {
			return path, true
		}
	}
	return "", false
}

// createBundleWidget creates new displayable widget for the given bundle
func createBundleWidget(bundle *swupd.Bundle) (gtk.IWidget, error) {
	// Create the root layout
	root, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}

	// Create display check
	check, err := gtk.CheckButtonNew()
	if err != nil {
		return nil, err
	}
	check.SetMarginTop(6)
	check.SetMarginStart(12)

	// Create display image
	img, err := gtk.ImageNew()
	img.SetMarginStart(6)
	img.SetMarginEnd(6)
	if err != nil {
		return nil, err
	}
	icon, set := LookupBundleIcon(bundle)
	if set {
		pbuf, err := gdk.PixbufNewFromFileAtSize(icon, 48, 48)
		if err != nil {
			icon = ""
			set = false
		} else {
			img.SetFromPixbuf(pbuf)
		}
	}

	// Still not set? Fallback.
	if !set {
		img.SetFromIconName("package-x-generic", gtk.ICON_SIZE_INVALID)
	}
	img.SetPixelSize(48)
	root.PackStart(img, false, false, 0)

	txt := fmt.Sprintf("<b>%s</b>\n%s", bundle.Name, bundle.Desc)
	label, err := gtk.LabelNew(txt)
	if err != nil {
		return nil, err
	}
	label.SetMarginStart(6)
	label.SetMarginEnd(12)
	label.SetXAlign(0.0)
	root.PackStart(label, false, false, 0)
	label.SetUseMarkup(true)

	check.Add(root)
	return check, nil
}

// NewBundlePage returns a new BundlePage
func NewBundlePage() (Page, error) {
	var err error
	var label *gtk.Label
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

	// label
	label, err = gtk.LabelNew("<big>Select bundles to install</big>")
	label.SetMarginTop(16)
	label.SetMarginBottom(16)
	label.SetHAlign(gtk.ALIGN_START)
	if err != nil {
		return nil, err
	}
	label.SetUseMarkup(true)
	bundle.box.PackStart(label, false, false, 0)

	// check list
	bundle.checks, err = gtk.FlowBoxNew()
	if err != nil {
		return nil, err
	}
	bundle.checks.SetSelectionMode(gtk.SELECTION_NONE)
	bundle.scroll, err = gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	// no horizontal scrolling
	bundle.scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	bundle.scroll.Add(bundle.checks)
	bundle.box.PackStart(bundle.scroll, true, true, 0)

	for _, b := range bundle.bundles {
		wid, err := createBundleWidget(b)
		if err != nil {
			return nil, err
		}
		bundle.checks.Add(wid)
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

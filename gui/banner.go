// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const (
	clearLinuxImage = "themes/clr.png"
)

// Banner is used to add a nice banner widget to the front of the installer
type Banner struct {
	revealer *gtk.Revealer // For animations
	box      *gtk.Box      // Main layout
	img      *gtk.Image    // Our image widget
	label    *gtk.Label    // Display label
}

// MakeHeader constructs the header component
func NewBanner() (*Banner, error) {
	var err error
	var pbuf *gdk.Pixbuf
	banner := &Banner{}

	// Create the "holder" (revealer)
	if banner.revealer, err = gtk.RevealerNew(); err != nil {
		return nil, err
	}

	// Create the root box
	if banner.box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0); err != nil {
		return nil, err
	}
	banner.revealer.Add(banner.box)
	banner.revealer.SetTransitionType(gtk.REVEALER_TRANSITION_TYPE_CROSSFADE)

	// Set the margins up
	banner.box.SetMarginTop(24)
	banner.box.SetMarginBottom(24)

	// Construct the image
	if banner.img, err = gtk.ImageNew(); err != nil {
		return nil, err
	}
	if pbuf, err = gdk.PixbufNewFromFileAtSize(clearLinuxImage, 128, 128); err != nil {
		return nil, err
	}
	banner.img.SetFromPixbuf(pbuf)
	banner.img.SetPixelSize(64)
	banner.img.SetHAlign(gtk.ALIGN_START)
	banner.box.PackStart(banner.img, false, false, 0)
	banner.box.SetHAlign(gtk.ALIGN_CENTER)

	// Sort the label out
	labelText := "<span font-size='xx-large'>Install Clear Linux* OS</span>\n\nTODO: Insert awesome header widget in this general region.\nKinda show off why this is an awesome decision."
	if banner.label, err = gtk.LabelNew(labelText); err != nil {
		return nil, err
	}
	banner.label.SetUseMarkup(true)
	banner.label.SetMarginStart(40)
	banner.label.SetHAlign(gtk.ALIGN_START)
	banner.label.SetVAlign(gtk.ALIGN_CENTER)
	banner.box.PackStart(banner.label, true, true, 0)

	return banner, nil
}

// GetRootWidget returns the embeddable root widget
func (banner *Banner) GetRootWidget() gtk.IWidget {
	return banner.revealer
}

// ShowFirst will display the banner for the first time during an intro sequence
func (banner *Banner) ShowFirst() {
	banner.revealer.SetTransitionType(gtk.REVEALER_TRANSITION_TYPE_CROSSFADE)
	banner.revealer.SetTransitionDuration(3000)
	banner.revealer.SetRevealChild(true)
}

// Show will animate the banner into view, showing the content
func (banner *Banner) Show() {
	banner.revealer.SetTransitionType(gtk.REVEALER_TRANSITION_TYPE_SLIDE_DOWN)
	banner.revealer.SetTransitionDuration(250)
	banner.revealer.SetRevealChild(true)
}

// Hide will animate the banner out of view, hiding the content
func (banner *Banner) Hide() {
	banner.revealer.SetTransitionType(gtk.REVEALER_TRANSITION_TYPE_SLIDE_UP)
	banner.revealer.SetTransitionDuration(250)
	banner.revealer.SetRevealChild(false)
}

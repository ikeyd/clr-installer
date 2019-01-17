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

// MakeHeader constructs the header component
func CreateBanner() (*gtk.Box, error) {
	var (
		err   error
		box   *gtk.Box
		img   *gtk.Image
		pbuf  *gdk.Pixbuf
		label *gtk.Label
	)

	// Create the root box
	if box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0); err != nil {
		return nil, err
	}
	// Set the margins up
	box.SetMarginTop(24)
	box.SetMarginBottom(24)

	// Construct the image
	if img, err = gtk.ImageNew(); err != nil {
		return nil, err
	}
	if pbuf, err = gdk.PixbufNewFromFileAtSize(clearLinuxImage, 128, 128); err != nil {
		return nil, err
	}
	img.SetFromPixbuf(pbuf)
	img.SetPixelSize(64)
	img.SetHAlign(gtk.ALIGN_START)
	box.PackStart(img, false, false, 0)
	box.SetHAlign(gtk.ALIGN_CENTER)

	// Sort the label out
	labelText := "<span font-size='xx-large'>Install Clear Linux* OS</span>\n\nTODO: Insert awesome header widget in this general region.\nKinda show off why this is an awesome decision."
	if label, err = gtk.LabelNew(labelText); err != nil {
		return nil, err
	}
	label.SetUseMarkup(true)
	label.SetMarginStart(40)
	label.SetHAlign(gtk.ALIGN_START)
	label.SetVAlign(gtk.ALIGN_CENTER)
	box.PackStart(label, true, true, 0)

	return box, nil
}

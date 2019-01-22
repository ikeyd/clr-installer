// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/clearlinux/clr-installer/gui/pages"
	"github.com/gotk3/gotk3/gtk"
)

// SummaryWidget is used within the ContentView to represent
// individual steps within the installer.
// It provides a description of the step, as well as a brief
// summary of the current selection/state (if set)
//
// In combination with the ContentView, this widget allows selection
// of each 'page' within the installer in a condensed form.
type SummaryWidget struct {
	handle *gtk.ListBoxRow
	box    *gtk.Box
	image  *gtk.Image
	label  *gtk.Label
	page   pages.Page
}

// NewSummaryWidget will construct a new SummaryWidget for the given page.
func NewSummaryWidget(page pages.Page) (*SummaryWidget, error) {
	// Create our root widget
	handle, err := gtk.ListBoxRowNew()
	if err != nil {
		return nil, err
	}

	// Create SummaryWidget
	s := &SummaryWidget{
		handle: handle,
		page:   page,
	}

	// Create layout box
	s.box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	s.box.SetHAlign(gtk.ALIGN_START)
	s.box.SetMarginStart(18)
	s.box.SetMarginEnd(18)
	s.box.SetMarginTop(6)
	s.box.SetMarginBottom(6)

	// Create image
	s.image, err = gtk.ImageNewFromIconName(page.GetIcon()+"-symbolic", gtk.ICON_SIZE_DIALOG)
	if err != nil {
		return nil, err
	}
	s.image.SetMarginEnd(12)
	s.box.PackStart(s.image, false, false, 0)

	// label
	s.label, err = gtk.LabelNew("<big>" + page.GetSummary() + "</big>")
	if err != nil {
		return nil, err
	}
	s.label.SetUseMarkup(true)
	s.label.SetHAlign(gtk.ALIGN_START)
	s.box.PackStart(s.label, false, false, 0)

	// Add to row and show it
	s.handle.Add(s.box)
	s.handle.ShowAll()

	return s, nil
}

// GetRootWidget returns the root embeddable widget for the SummaryWidget
func (s *SummaryWidget) GetRootWidget() *gtk.ListBoxRow {
	return s.handle
}

// GetRowIndex returns the row index of our internal GtkListBoxRow
func (s *SummaryWidget) GetRowIndex() int {
	return s.handle.GetIndex()
}

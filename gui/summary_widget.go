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
	layout *gtk.Box
	box    *gtk.Box
	image  *gtk.Image
	label  *gtk.Label
	value  *gtk.Label
	page   pages.Page
}

// NewSummaryWidget will construct a new SummaryWidget for the given page.
func NewSummaryWidget(page pages.Page) (*SummaryWidget, error) {
	var st *gtk.StyleContext

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

	// Add styling
	st, err = s.handle.GetStyleContext()
	if err != nil {
		return nil, err
	}
	st.AddClass("installer-summary-widget")

	// Create layout box
	s.layout, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	s.layout.SetHAlign(gtk.ALIGN_START)
	s.layout.SetMarginStart(18)
	s.layout.SetMarginEnd(18)
	s.layout.SetMarginTop(6)
	s.layout.SetMarginBottom(6)

	// Create image
	s.image, err = gtk.ImageNewFromIconName(page.GetIcon()+"-symbolic", gtk.ICON_SIZE_DIALOG)
	if err != nil {
		return nil, err
	}
	s.image.SetMarginEnd(12)
	s.layout.PackStart(s.image, false, false, 0)

	// Label box
	s.box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	s.layout.PackStart(s.box, false, false, 0)

	// title label
	s.label, err = gtk.LabelNew("<big>" + page.GetSummary() + "</big>")
	if err != nil {
		return nil, err
	}
	s.label.SetUseMarkup(true)
	s.label.SetHAlign(gtk.ALIGN_START)
	s.box.PackStart(s.label, false, false, 0)

	// value label
	s.value, err = gtk.LabelNew("")
	if err != nil {
		return nil, err
	}
	st, err = s.value.GetStyleContext()
	if err != nil {
		return nil, err
	}
	st.AddClass("configured-value")
	s.value.SetUseMarkup(false)
	s.value.SetHAlign(gtk.ALIGN_START)
	s.box.PackStart(s.value, false, false, 0)

	// Do not show by ShowAll() or by default, to allow hiding.
	s.value.ShowAll()
	s.value.SetNoShowAll(true)
	s.value.Hide()

	// Add to row and show it
	s.handle.Add(s.layout)
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

// Update will alter the view to show the currently configured values
func (s *SummaryWidget) Update() {
	value := s.page.GetConfiguredValue()
	if value == "" {
		s.value.Hide()
		return
	}
	s.value.SetText(value)
	s.value.Show()
}

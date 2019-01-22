// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/model"
	"github.com/clearlinux/clr-installer/timezone"
	"github.com/gotk3/gotk3/gtk"
)

// Timezone is a simple page to help with timezone settings
type Timezone struct {
	controller Controller
	model      *model.SystemInstall
	timezones  []*timezone.TimeZone
	box        *gtk.Box
	scroll     *gtk.ScrolledWindow
	list       *gtk.ListBox
}

// NewTimezonePage returns a new TimezonePage
func NewTimezonePage(controller Controller, model *model.SystemInstall) (Page, error) {
	tzones, err := timezone.Load()
	if err != nil {
		return nil, err
	}

	t := &Timezone{
		controller: controller,
		model:      model,
		timezones:  tzones,
	}

	t.box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	t.box.SetBorderWidth(8)

	// Build storage for listbox
	t.scroll, err = gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	t.box.PackStart(t.scroll, true, true, 0)
	t.scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)

	// Build listbox
	t.list, err = gtk.ListBoxNew()
	if err != nil {
		return nil, err
	}
	t.list.SetSelectionMode(gtk.SELECTION_SINGLE)
	t.list.SetActivateOnSingleClick(true)
	// t.list.Connect("row-activated", timezone.onRowActivated)
	t.scroll.Add(t.list)
	// Remove background
	st, _ := t.list.GetStyleContext()
	st.AddClass("scroller-special")

	for _, zone := range t.timezones {
		lab, err := gtk.LabelNew("<big>" + zone.Code + "</big>")
		if err != nil {
			return nil, err
		}
		lab.SetUseMarkup(true)
		lab.SetHAlign(gtk.ALIGN_START)
		lab.SetXAlign(0.0)
		lab.ShowAll()
		t.list.Add(lab)
	}

	return t, nil
}

// IsRequired will return true as we always need a timezone
func (t *Timezone) IsRequired() bool {
	return true
}

func (t *Timezone) GetID() int {
	return PageIDTimezone
}

func (t *Timezone) GetIcon() string {
	return "preferences-system-time"
}

func (t *Timezone) GetRootWidget() gtk.IWidget {
	return t.box
}

func (t *Timezone) GetSummary() string {
	return "Choose Timezone"
}

func (t *Timezone) GetTitle() string {
	return t.GetSummary()
}

func (t *Timezone) StoreChanges() {}

// ResetChanges will find the default model selection and set the
// timezone as appropriate in the view
func (t *Timezone) ResetChanges() {
	code := timezone.DefaultTimezone
	if t.model.Timezone.Code != "" {
		code = t.model.Timezone.Code
	}

	// Preselect the timezone here
	for n, tz := range t.timezones {
		if tz.Code != code {
			continue
		}
		row := t.list.GetRowAtIndex(n)
		t.list.SelectRow(row)
		scrollToView(t.scroll, t.list, &row.Widget)
	}
}

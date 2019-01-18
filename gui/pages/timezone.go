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
}

// NewTimezonePage returns a new TimezonePage
func NewTimezonePage(controller Controller, model *model.SystemInstall) (Page, error) {
	tzones, err := timezone.Load()
	if err != nil {
		return nil, err
	}
	return &Timezone{
		controller: controller,
		model:      model,
		timezones:  tzones,
	}, nil
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
	wid, _ := gtk.LabelNew("I am the timezone page :O")
	return wid
}

func (t *Timezone) GetSummary() string {
	return "Choose Timezone"
}

func (t *Timezone) GetTitle() string {
	return t.GetSummary()
}

func (t *Timezone) StoreChanges() {}
func (t *Timezone) ResetChanges() {}

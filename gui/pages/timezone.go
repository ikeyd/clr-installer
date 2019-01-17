// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/timezone"
	"github.com/gotk3/gotk3/gtk"
)

// Timezone is a simple page to help with timezone settings
type Timezone struct {
	timezones []*timezone.TimeZone
}

// NewTimezonePage returns a new TimezonePage
func NewTimezonePage() (Page, error) {
	tzones, err := timezone.Load()
	if err != nil {
		return nil, err
	}
	return &Timezone{
		timezones: tzones,
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
	return nil
}

func (t *Timezone) GetTitle() string {
	return "Choose Timezone"
}

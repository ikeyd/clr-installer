// Copyright © 2018 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/timezone"
)

// Timezone is a simple page to help with timezone settings
type Timezone struct {
	Page

	timezones []*timezone.TimeZone
}

// NewTimezonePage returns a new TimezonePage
func NewTimezonePage() *Timezone {
	return &Timezone{}
}

// IsRequired will return true as we always need a timezone
func (t *Timezone) IsRequired() bool {
	return true
}

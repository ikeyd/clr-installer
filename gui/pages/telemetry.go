// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/gotk3/gotk3/gtk"
)

// Telemetry is a simple page to help with Telemetry settings
type Telemetry struct {
}

// NewTelemetryPage returns a new TelemetryPage
func NewTelemetryPage() (Page, error) {
	return &Telemetry{}, nil
}

// IsRequired will return true as we always need a Telemetry
func (t *Telemetry) IsRequired() bool {
	return true
}

func (t *Telemetry) GetID() int {
	return PageIDTelemetry
}

func (t *Telemetry) GetIcon() string {
	return "web-browser"
}

func (t *Telemetry) GetRootWidget() gtk.IWidget {
	return nil
}

func (t *Telemetry) GetSummary() string {
	return "Telemetry"
}

func (t *Telemetry) GetTitle() string {
	return "Enable Telemetry"
}

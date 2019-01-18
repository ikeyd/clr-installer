// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/telemetry"
	"github.com/gotk3/gotk3/gtk"
)

// Telemetry is a simple page to help with Telemetry settings
type Telemetry struct {
	box *gtk.Box
	sw  *gtk.CheckButton
}

// NewTelemetryPage returns a new TelemetryPage
func NewTelemetryPage() (Page, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	box.SetVAlign(gtk.ALIGN_CENTER)

	lab, err := gtk.LabelNew(telemetry.HelpMarkdown)
	if err != nil {
		return nil, err
	}
	lab.SetUseMarkup(true)
	box.PackStart(lab, false, false, 0)

	sw, err := gtk.CheckButtonNewWithLabel("Enable telemetry")
	if err != nil {
		return nil, err
	}
	sw.SetHAlign(gtk.ALIGN_CENTER)
	box.PackStart(sw, false, false, 0)

	return &Telemetry{box: box, sw: sw}, nil
}

// IsRequired will return true as we always need a Telemetry
func (t *Telemetry) IsRequired() bool {
	return true
}

func (t *Telemetry) GetID() int {
	return PageIDTelemetry
}

func (t *Telemetry) GetIcon() string {
	return "network-transmit-receive"
}

func (t *Telemetry) GetRootWidget() gtk.IWidget {
	return t.box
}

func (t *Telemetry) GetSummary() string {
	return "Telemetry"
}

func (t *Telemetry) GetTitle() string {
	return telemetry.Title
}

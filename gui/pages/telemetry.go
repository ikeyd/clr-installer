// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/model"
	"github.com/clearlinux/clr-installer/telemetry"
	"github.com/gotk3/gotk3/gtk"
)

// Telemetry is a simple page to help with Telemetry settings
type Telemetry struct {
	model      *model.SystemInstall
	controller Controller
	box        *gtk.Box
	check      *gtk.CheckButton
}

// NewTelemetryPage returns a new TelemetryPage
func NewTelemetryPage(controller Controller, model *model.SystemInstall) (Page, error) {
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

	check, err := gtk.CheckButtonNewWithLabel("Enable telemetry")
	if err != nil {
		return nil, err
	}
	check.SetHAlign(gtk.ALIGN_CENTER)
	box.PackStart(check, false, false, 0)

	return &Telemetry{
		controller: controller,
		model:      model,
		box:        box,
		check:      check}, nil
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

func (t *Telemetry) StoreChanges() {
	t.model.EnableTelemetry(t.check.GetActive())
}

func (t *Telemetry) ResetChanges() {
	t.controller.SetButtonState(ButtonConfirm, true)
	t.check.SetActive(t.model.IsTelemetryEnabled())
}

// GetConfiguredValue returns our current config
func (t *Telemetry) GetConfiguredValue() string {
	if t.model.IsTelemetryEnabled() {
		return "Enabled"
	}
	return "Disabled"
}

// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"fmt"
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/gtk"
)

// InstallPage is a specialised page type with no corresponding
// ContentView summary. It handles the actual install routine.
type InstallPage struct {
	controller Controller
	model      *model.SystemInstall
	layout     *gtk.Box
}

// NewInstallPage constructs a new InstallPage.
func NewInstallPage(controller Controller, model *model.SystemInstall) (Page, error) {
	layout, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	return &InstallPage{
		controller: controller,
		model:      model,
		layout:     layout,
	}, nil
}

func (install *InstallPage) IsRequired() bool {
	return true
}

func (install *InstallPage) IsDone() bool {
	return false
}

func (install *InstallPage) GetID() int {
	return PageIDInstall
}

func (install *InstallPage) GetSummary() string {
	return "Installing Clear Linux OS"
}

func (install *InstallPage) GetTitle() string {
	return "Installing Clear Linux OS"
}

func (install *InstallPage) GetIcon() string {
	return "system-software-install-symbolic"
}

func (install *InstallPage) GetConfiguredValue() string {
	return ""
}

func (install *InstallPage) GetRootWidget() gtk.IWidget {
	return install.layout
}

func (install *InstallPage) StoreChanges() {}

// ResetChanges begins as our initial execution point as we're only going
// to get called when showing our page.
func (install *InstallPage) ResetChanges() {
	// Disable quit button
	install.controller.SetButtonState(ButtonQuit, false)

	// Validate the model
	err := install.model.Validate()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Validation passed")
}

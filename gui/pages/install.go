// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/gtk"
)

// InstallPage is a specialised page type with no corresponding
// ContentView summary. It handles the actual install routine.
type InstallPage struct {
	controller Controller
	model      *model.SystemInstall
}

// NewInstallPage constructs a new InstallPage.
func NewInstallPage(controller Controller, model *model.SystemInstall) (Page, error) {
	return &InstallPage{
		controller: controller,
		model:      model,
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
	return nil
}

func (install *InstallPage) StoreChanges() {}
func (install *InstallPage) ResetChanges() {}

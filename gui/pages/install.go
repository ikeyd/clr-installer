// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"fmt"
	ctrl "github.com/clearlinux/clr-installer/controller"
	"github.com/clearlinux/clr-installer/model"
	"github.com/clearlinux/clr-installer/progress"
	"github.com/gotk3/gotk3/gtk"
	"time"
)

var (
	loopWaitDuration = 2 * time.Second
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

	// TODO: Disable closing of the installer
	go func() {
		progress.Set(install)
		err := ctrl.Install(install.controller.GetRootDir(),
			install.model,
			install.controller.GetOptions(),
		)
		panic(err)
	}()

}

// Following methods are for the progress.Client API
func (install *InstallPage) Desc(desc string) {
	fmt.Println(desc)
}

// Failure handles failure to install
func (install *InstallPage) Failure() {
	fmt.Println("Failure")
}

// LoopWaitDuration will return the duration for step-waits
func (install *InstallPage) LoopWaitDuration() time.Duration {
	return loopWaitDuration
}

// Partial handles an actual progress update
func (install *InstallPage) Partial(total int, step int) {
}

// Step will step the progressbar in indeterminate mode
func (install *InstallPage) Step() {
}

// Success notes the install was successful
func (install *InstallPage) Success() {
	fmt.Println("Success")
}

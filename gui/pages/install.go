// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"fmt"
	_ "github.com/clearlinux/clr-installer/controller"
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

	pbar *gtk.ProgressBar // Progress bar
}

// NewInstallPage constructs a new InstallPage.
func NewInstallPage(controller Controller, model *model.SystemInstall) (Page, error) {
	var err error

	// Create page
	page := &InstallPage{
		controller: controller,
		model:      model,
	}

	// Create layout
	page.layout, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}

	// Create progressbar
	page.pbar, err = gtk.ProgressBarNew()
	if err != nil {
		return nil, err
	}

	// Sort out padding
	page.pbar.SetHAlign(gtk.ALIGN_FILL)
	page.pbar.SetMarginStart(24)
	page.pbar.SetMarginEnd(24)
	page.pbar.SetMarginBottom(12)
	page.pbar.SetMarginTop(12)

	// Throw it on the bottom of the page
	page.layout.PackEnd(page.pbar, false, false, 0)

	return page, nil
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
		// Become the progress hook
		progress.Set(install)

		// Go install it
		//err := ctrl.Install(install.controller.GetRootDir(),
		//	install.model,
		//	install.controller.GetOptions(),
		//)

		// TODO: Handle this moar better.
		//if err != nil {
		//	panic(err)
		//}
		fmt.Println("Installation completed")
		install.controller.SetButtonState(ButtonQuit, true)
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
	install.pbar.Pulse()
}

// Success notes the install was successful
func (install *InstallPage) Success() {
	fmt.Println("Success")
}

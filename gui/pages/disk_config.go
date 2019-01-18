// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/model"
	"github.com/clearlinux/clr-installer/storage"
	"github.com/gotk3/gotk3/gtk"
)

// DiskConfig is a simple page to help with DiskConfig settings
type DiskConfig struct {
	model *model.SystemInstall
}

// NewDiskConfigPage returns a new DiskConfigPage
func NewDiskConfigPage(model *model.SystemInstall) (Page, error) {
	return &DiskConfig{model: model}, nil
}

// IsRequired will return true as we always need a DiskConfig
func (t *DiskConfig) IsRequired() bool {
	return true
}

func (t *DiskConfig) GetID() int {
	return PageIDDiskConfig
}

func (t *DiskConfig) GetIcon() string {
	return "media-removable"
}

func (t *DiskConfig) GetRootWidget() gtk.IWidget {
	return nil
}

func (t *DiskConfig) GetSummary() string {
	return "Configure Media"
}

func (t *DiskConfig) GetTitle() string {
	return t.GetSummary() + " - WARNING: SUPER EXPERIMENTAL"
}

func (t *DiskConfig) StoreChanges() {}
func (t *DiskConfig) ResetChanges() {
	store, _ := storage.RescanBlockDevices(t.model.TargetMedias)
	for _, device := range store {
		print(device.Name + " - " + device.FsType)
	}
}

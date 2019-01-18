// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/gtk"
)

// DiskConfig is a simple page to help with DiskConfig settings
type DiskConfig struct {
}

// NewDiskConfigPage returns a new DiskConfigPage
func NewDiskConfigPage() (Page, error) {
	return &DiskConfig{}, nil
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

func (t *DiskConfig) StoreChanges(model *model.SystemInstall) {}
func (t *DiskConfig) ResetChanges(model *model.SystemInstall) {}

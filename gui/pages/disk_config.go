// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"fmt"
	"github.com/clearlinux/clr-installer/model"
	"github.com/clearlinux/clr-installer/storage"
	"github.com/gotk3/gotk3/gtk"
)

// DiskConfig is a simple page to help with DiskConfig settings
type DiskConfig struct {
	devs  []*storage.BlockDevice
	model *model.SystemInstall
}

// NewDiskConfigPage returns a new DiskConfigPage
func NewDiskConfigPage(model *model.SystemInstall) (Page, error) {
	disk := &DiskConfig{
		model: model,
	}
	devs, err := disk.buildDisks()
	if err != nil {
		return nil, err
	}
	disk.devs = devs
	return disk, nil
}

func (disk *DiskConfig) buildDisks() ([]*storage.BlockDevice, error) {
	//return storage.RescanBlockDevices(disk.model.TargetMedias)
	devices, err := storage.RescanBlockDevices(nil)
	for _, device := range devices {
		storage.NewStandardPartitions(device)
	}
	for _, device := range devices {
		fmt.Println(device.GetDeviceFile())
		for _, device := range device.Children {
			fmt.Println("\t" + device.GetDeviceFile() + " - " + device.FsType)
		}
	}
	return devices, err
}

// IsRequired will return true as we always need a DiskConfig
func (disk *DiskConfig) IsRequired() bool {
	return true
}

func (disk *DiskConfig) GetID() int {
	return PageIDDiskConfig
}

func (disk *DiskConfig) GetIcon() string {
	return "media-removable"
}

func (disk *DiskConfig) GetRootWidget() gtk.IWidget {
	return nil
}

func (disk *DiskConfig) GetSummary() string {
	return "Configure Media"
}

func (disk *DiskConfig) GetTitle() string {
	return disk.GetSummary() + " - WARNING: SUPER EXPERIMENTAL"
}

func (disk *DiskConfig) StoreChanges() {}
func (disk *DiskConfig) ResetChanges() {
}

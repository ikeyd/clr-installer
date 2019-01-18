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
	devs   []*storage.BlockDevice
	model  *model.SystemInstall
	box    *gtk.Box
	scroll *gtk.ScrolledWindow
	list   *gtk.ListBox
}

// NewDiskConfigPage returns a new DiskConfigPage
func NewDiskConfigPage(model *model.SystemInstall) (Page, error) {
	disk := &DiskConfig{
		model: model,
	}
	var placeholder *gtk.Label

	devs, err := disk.buildDisks()
	if err != nil {
		return nil, err
	}
	disk.devs = devs

	disk.box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	disk.box.SetBorderWidth(8)

	// Build storage for listbox
	disk.scroll, err = gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	disk.box.PackStart(disk.scroll, true, true, 0)
	disk.scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)

	// Build listbox
	disk.list, err = gtk.ListBoxNew()
	if err != nil {
		return nil, err
	}
	disk.scroll.Add(disk.list)
	// Remove background
	st, _ := disk.list.GetStyleContext()
	st.AddClass("scroller-special")

	// Set placeholder
	placeholder, err = gtk.LabelNew("No usable devices found")
	if err != nil {
		return nil, err
	}

	placeholder.ShowAll()
	disk.list.SetPlaceholder(placeholder)

	if err = disk.buildList(); err != nil {
		return nil, err
	}

	return disk, nil
}

func (disk *DiskConfig) buildDisks() ([]*storage.BlockDevice, error) {
	//return storage.RescanBlockDevices(disk.model.TargetMedias)
	devices, err := storage.RescanBlockDevices(nil)
	for _, device := range devices {
		storage.NewStandardPartitions(device)
	}
	return devices, err
}

// buildList populates the ListBox with usable widget things
func (disk *DiskConfig) buildList() error {
	for _, device := range disk.devs {
		box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		if err != nil {
			return err
		}
		img, err := gtk.ImageNewFromIconName("drive-harddisk-system-symbolic", gtk.ICON_SIZE_DIALOG)
		if err != nil {
			return err
		}
		img.SetMarginEnd(12)
		img.SetMarginStart(12)
		box.PackStart(img, false, false, 0)
		text := fmt.Sprintf("<big>Wipe all data on: <b>%s</b></big>\n", device.GetDeviceFile())
		for _, child := range device.Children {
			text += fmt.Sprintf(" - Create partition <b>%s</b> with type <b>%s</b> (%s)\n", child.GetDeviceFile(), child.FsType, child.Label)
		}

		label, err := gtk.LabelNew(text)
		if err != nil {
			return err
		}
		label.SetXAlign(0.0)
		label.SetHAlign(gtk.ALIGN_START)
		label.SetUseMarkup(true)
		box.PackStart(label, false, false, 0)
		box.ShowAll()
		disk.list.Add(box)
	}
	return nil
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
	return disk.box
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

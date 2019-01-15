// Copyright © 2018 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package gui

import (
	"github.com/clearlinux/clr-installer/args"
	"github.com/clearlinux/clr-installer/model"
	"github.com/gotk3/gotk3/gtk"
)

// Gui is the main tui data struct and holds data about the higher level data for this
// front end, it also implements the Frontend interface
type Gui struct {
	window        *Window
	model         *model.SystemInstall
	options       args.Args
	rootDir       string
	installReboot bool
}

// New creates a new Gui frontend instance
func New() *Gui {
	return &Gui{}
}

// MustRun is part of the Frontend interface implementation and tells the core that this
// frontend wants/must run.
func (gui *Gui) MustRun(args *args.Args) bool {
	if args.ForceGUI {
		return true
	}

	return gtk.InitCheck(nil) == nil
}

// Run is part of the Frontend interface implementation and is the gui frontend main entry point
func (gui *Gui) Run(md *model.SystemInstall, rootDir string, options args.Args) (bool, error) {
	gui.model = md
	gui.options = options
	gui.rootDir = rootDir
	gui.installReboot = false

	// Construct main window
	win, err := NewWindow()
	if err != nil {
		return gui.installReboot, err
	}
	gui.window = win

	// Main loop
	gtk.Main()

	return gui.installReboot, nil
}

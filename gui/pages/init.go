// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/gotk3/gotk3/gtk"
)

// Button allows us to flag up different buttons
type Button uint

const (
	// ButtonCancel enables the cancel button
	ButtonCancel Button = 1 << iota

	// ButtonConfirm enables the confirm button
	ButtonConfirm Button = 1 << iota
)

// Page interface provides a common definition that other
// pages can share to give a standard interface for the
// main controller, the Window
type Page interface {
	IsRequired() bool
	GetID() int
	GetSummary() string
	GetTitle() string
	GetIcon() string
	GetRootWidget() gtk.IWidget
	StoreChanges() // Store changes in the model
	ResetChanges() // Reset data to model
}

// PageController is implemented by the Window struct, and
// is used by pages and ContentView to exert some control
// over workflow.
type Controller interface {
	ActivatePage(Page)
	SetButtonState(flags Button, enabled bool)
}

const (
	PageIDTimezone   = iota
	PageIDLanguage   = iota
	PageIDKeyboard   = iota
	PageIDBundle     = iota
	PageIDTelemetry  = iota
	PageIDDiskConfig = iota
)

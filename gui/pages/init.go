// Copyright © 2018-2019 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package pages

import (
	"github.com/gotk3/gotk3/gtk"
)

// Page interface provides a common definition that other
// pages can share to give a standard interface for the
// main controller, the Window
type Page interface {
	IsRequired() bool
	GetID() int
	GetTitle() string
	GetIcon() string
	GetRootWidget() gtk.IWidget
}

// PageController is implemented by the Window struct, and
// is used by pages and ContentView to exert some control
// over workflow.
type Controller interface {
	ActivatePage(Page)
}

const (
	PageIDTimezone = iota
	PageIDLanguage = iota
	PageIDKeyboard = iota
	PageIDBundle   = iota
)

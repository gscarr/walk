// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type GroupBox struct {
	AssignTo           **walk.GroupBox
	Name               string
	Enabled            Property
	Visible            Property
	Font               Font
	ToolTipText        Property
	MinSize            Size
	MaxSize            Size
	StretchFactor      int
	Row                int
	RowSpan            int
	Column             int
	ColumnSpan         int
	ContextMenuActions []*walk.Action
	OnKeyDown          walk.KeyEventHandler
	OnMouseDown        walk.MouseEventHandler
	OnMouseMove        walk.MouseEventHandler
	OnMouseUp          walk.MouseEventHandler
	OnSizeChanged      walk.EventHandler
	Title              string
	DataBinder         DataBinder
	Layout             Layout
	Children           []Widget
}

func (gb GroupBox) Create(builder *Builder) error {
	w, err := walk.NewGroupBox(builder.Parent())
	if err != nil {
		return err
	}

	return builder.InitWidget(gb, w, func() error {
		if err := w.SetTitle(gb.Title); err != nil {
			return err
		}

		if gb.AssignTo != nil {
			*gb.AssignTo = w
		}

		return nil
	})
}

func (w GroupBox) WidgetInfo() (name string, disabled, hidden bool, font *Font, toolTipText string, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, contextMenuActions []*walk.Action, OnKeyDown walk.KeyEventHandler, OnMouseDown walk.MouseEventHandler, OnMouseMove walk.MouseEventHandler, OnMouseUp walk.MouseEventHandler, OnSizeChanged walk.EventHandler) {
	return w.Name, false, false, &w.Font, "", w.MinSize, w.MaxSize, w.StretchFactor, w.Row, w.RowSpan, w.Column, w.ColumnSpan, w.ContextMenuActions, w.OnKeyDown, w.OnMouseDown, w.OnMouseMove, w.OnMouseUp, w.OnSizeChanged
}

func (gb GroupBox) ContainerInfo() (DataBinder, Layout, []Widget) {
	return gb.DataBinder, gb.Layout, gb.Children
}

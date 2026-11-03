package view

import "github.com/egoist/mygo/ui"

// Ext is a Raycast command shown in the panel (raycast.go).
type Ext struct{ Open bool }

func (e *Ext) Height() float32                                { return 97 }
func (e *Ext) view(c *ui.Context, p *Panel, shell ui.Element) {}

package view

import "github.com/egoist/mygo/ui"

// Welcome is the first-run welcome (welcome.go).
type Welcome struct{ Active bool }

func (w *Welcome) Height() float32              { return 330 }
func (w *Welcome) view(c *ui.Context, p *Panel) {}

// Ext is a Raycast command shown in the panel (raycast.go).
type Ext struct{ Open bool }

func (e *Ext) Height() float32                                { return 97 }
func (e *Ext) view(c *ui.Context, p *Panel, shell ui.Element) {}

// Settings is the Settings screen (settings.go).
type Settings struct{ Open bool }

func (s *Settings) init()                                          {}
func (s *Settings) Show()                                          { s.Open = true }
func (s *Settings) Hide()                                          { s.Open = false }
func (s *Settings) view(c *ui.Context, p *Panel, shell ui.Element) {}

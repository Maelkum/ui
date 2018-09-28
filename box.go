// 12 december 2015

package ui

import (
	"unsafe"
)

// #include "pkgui.h"
import "C"

// Box is a Control that holds a group of Controls horizontally
// or vertically. If horizontally, then all controls have the same
// height. If vertically, then all controls have the same width.
// By default, each control has its preferred width (horizontal)
// or height (vertical); if a control is marked "stretchy", it will
// take whatever space is left over. If multiple controls are marked
// stretchy, they will be given equal shares of the leftover space.
// There can also be space between each control ("padding").
type Box struct {
	ControlBase
	b	*C.uiBox
	children	[]Control
}

// NewHorizontalBox creates a new horizontal Box.
func NewHorizontalBox() *Box {
	b := new(Box)

	b.b = C.uiNewHorizontalBox()

	b.ControlBase = NewControlBase(b, uintptr(unsafe.Pointer(b.b)))
	return b
}

// NewVerticalBox creates a new vertical Box.
func NewVerticalBox() *Box {
	b := new(Box)

	b.b = C.uiNewVerticalBox()

	b.ControlBase = NewControlBase(b, uintptr(unsafe.Pointer(b.b)))
	return b
}

// Destroy destroys the Box. If the Box has children,
// Destroy calls Destroy on those Controls as well.
func (b *Box) Destroy() {
	for len(b.children) != 0 {
		c := b.children[0]
		b.Delete(0)
		c.Destroy()
	}
	b.ControlBase.Destroy()
}

// Append adds the given control to the end of the Box.
func (b *Box) Append(child Control, stretchy bool) {
	b.AppendMultipleEx(stretchy, child)
}

// Append multiple (non-stretchy) controls to the end of the Box.
func (b *Box) AppendMultiple(children ...Control) {
	b.AppendMultipleEx(false, children...)
}

// AppendMultiple but allows "stretchy" parameter
func (b *Box) AppendMultipleEx(stretchy bool, children ...Control) {
	c := (*C.uiControl)(nil)

	for _, child := range children {
		if nil == child {
			continue
		}

		c = touiControl(child.LibuiControl())

		C.uiBoxAppend(b.b, c, frombool(stretchy))
		b.children = append(b.children, child)
	}
}

// Delete deletes the nth control of the Box.
func (b *Box) Delete(n int) {
	b.children = append(b.children[:n], b.children[n + 1:]...)
	C.uiBoxDelete(b.b, C.int(n))
}

// Padded returns whether there is space between each control
// of the Box.
func (b *Box) Padded() bool {
	return tobool(C.uiBoxPadded(b.b))
}

// SetPadded controls whether there is space between each control
// of the Box. The size of the padding is determined by the OS and
// its best practices.
func (b *Box) SetPadded(padded bool) {
	C.uiBoxSetPadded(b.b, frombool(padded))
}

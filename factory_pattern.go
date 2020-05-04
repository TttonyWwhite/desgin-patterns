package desgin_patterns

import "fmt"

// 工厂模式

type Button interface {
	render()
	onClick()
}

type Dialog interface {
	render()
	createButton() Button
}

type WindowsButton struct {
}

func (w *WindowsButton) render() {
	// windows button's render function
}

func (w *WindowsButton) onClick() {
	fmt.Println("this is a windows button")
}

type HTMLButton struct {
}

func (h *HTMLButton) render() {
	// HTML button's render function
}

func (h *HTMLButton) onClick() {
	fmt.Println("this is a html button")
}

type WindowsDialog struct {
}

func (w *WindowsDialog) render() {
	// windows dialog's render function
}

func (w *WindowsDialog) createButton() Button {
	return &WindowsButton{}
}

type HTMLDialog struct {
}

func (h *HTMLDialog) render() {
	// html dialog's render function
}

func (h *HTMLDialog) createButton() Button {
	return &HTMLButton{}
}

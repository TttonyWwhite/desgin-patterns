package desgin_patterns

import "fmt"

type Checkbox interface {
	render()
	check()
}

type GUIFactory interface {
	createButton() Button
	createCheckBox() Checkbox
}

type WinButton struct {
	// windows风格的按钮
}

func (w *WinButton) render() {
	// windows button's render function
}

func (w *WinButton) onClick() {
	fmt.Println("click on WinButton")
}

type WinCheckbox struct {
	// windows风格的checkbox
}

func (w *WinCheckbox) render() {
	// windows checkbox's render function
}

func (w *WinCheckbox) check() {
	fmt.Println("check on WinCheckbox")
}

type WinFactory struct {
}

func (w *WinFactory) createButton() Button {
	return &WinButton{}
}

func (w *WinFactory) createCheckBox() Checkbox {
	return &WinCheckbox{}
}

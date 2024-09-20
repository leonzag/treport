package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewHBoxLeading(lead fyne.CanvasObject, obj ...fyne.CanvasObject) *fyne.Container {
	box := container.NewHBox(obj...)
	return container.NewBorder(nil, nil, nil, box, lead)
}

func NewHBoxTrailing(trail fyne.CanvasObject, obj ...fyne.CanvasObject) *fyne.Container {
	box := container.NewHBox(obj...)
	return container.NewBorder(nil, nil, box, nil, trail)
}

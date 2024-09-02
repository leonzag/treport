package interfaces

import "fyne.io/fyne/v2"

type ProgressBarInfinite interface {
	Content() fyne.CanvasObject
	Show()
	Hide()
	Running() bool
}

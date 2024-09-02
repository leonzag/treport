package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui/interfaces"
)

type progressBarInfinite struct {
	bar *widget.ProgressBarInfinite
}

var _ interfaces.ProgressBarInfinite = new(progressBarInfinite)

func NewProgressBarInfinite() *progressBarInfinite {
	bar := widget.NewProgressBarInfinite()
	bar.Hide()
	return &progressBarInfinite{bar: bar}
}

func (b *progressBarInfinite) Content() fyne.CanvasObject {
	return b.bar
}

func (b *progressBarInfinite) Show() {
	b.bar.Start()
	b.bar.Show()
}

func (b *progressBarInfinite) Hide() {
	b.bar.Hide()
	b.bar.Stop()
}

func (b *progressBarInfinite) Running() bool {
	return b.bar.Running()
}

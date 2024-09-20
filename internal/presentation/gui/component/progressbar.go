package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ProgressBarInfinite struct {
	widget.BaseWidget
	bar *widget.ProgressBarInfinite
}

func NewProgressBarInfinite() *ProgressBarInfinite {
	bar := widget.NewProgressBarInfinite()
	bar.Hide()
	w := &ProgressBarInfinite{bar: bar}
	w.ExtendBaseWidget(w)

	return w
}

func (w *ProgressBarInfinite) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(w.bar)
}

func (b *ProgressBarInfinite) Show() {
	b.bar.Start()
	b.bar.Show()
}

func (b *ProgressBarInfinite) Hide() {
	b.bar.Hide()
	b.bar.Stop()
}

func (b *ProgressBarInfinite) Running() bool {
	return b.bar.Running()
}

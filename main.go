//go:generate fyne bundle -o data.go data

package main

import (
	"fmt"
	"image/color"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func makeUI() fyne.CanvasObject {
	merry := canvas.NewText("Merry Christmas", color.White)
	merry.Alignment = fyne.TextAlignCenter
	merry.TextSize = 42
	to := widget.NewLabel("To our customers and the Fyne community")
	to.Alignment = fyne.TextAlignCenter
	texts := container.NewVBox(merry,
		container.NewHBox(widget.NewLabel("From the Fyne Labs team"),
			layout.NewSpacer(),
			widget.NewButton("Donate", func() {
				u, _ := url.Parse("https://github.com/sponsors/fyne-io")
				_ = fyne.CurrentApp().OpenURL(u)
			})))

	flake := newSnowFlake(3)
	label := widget.NewLabel(fmt.Sprintf("%d", flake.Count))
	down := widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
		if flake.Count <= 1 {
			return
		}

		flake.Count--
		label.SetText(fmt.Sprintf("%d", flake.Count))
		flake.Refresh()
	})
	down.Importance = widget.LowImportance
	up := widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
		if flake.Count >= 6 {
			return
		}

		flake.Count++
		label.SetText(fmt.Sprintf("%d", flake.Count))
		flake.Refresh()
	})
	up.Importance = widget.LowImportance
	return container.NewBorder(to, texts, nil, nil, flake,
		container.NewCenter(container.NewHBox(
			down, label, up)))
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&cardTheme{})
	w := a.NewWindow("Snowflake")

	w.SetPadded(true)
	w.SetContent(container.NewPadded(makeUI()))
	w.Resize(fyne.NewSize(360, 520))
	w.ShowAndRun()
}

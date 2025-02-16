package page

import "github.com/rivo/tview"

func NewAlertModal() *tview.Modal {
	return tview.NewModal().
		AddButtons([]string{"Close"})
}

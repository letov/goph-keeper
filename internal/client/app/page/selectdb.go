package page

import (
	"os"

	"github.com/rivo/tview"
)

func NewSelectDBModal(
	pages *tview.Pages,
	selectDbCallback func(action int) error,
) *tview.Modal {
	return tview.NewModal().
		SetText("Select DB action").
		AddButtons([]string{"Open DB", "Load remote DB", "Save to remote DB", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonIndex {
			case 0:
				if err := selectDbCallback(buttonIndex); err == nil {
					pages.SwitchToPage(PrivateListPage)
				}
			case 1:
				if err := selectDbCallback(buttonIndex); err == nil {
					pages.SwitchToPage(PrivateListPage)
				}
			case 2:
				if err := selectDbCallback(buttonIndex); err == nil {
					pages.SwitchToPage(PrivateListPage)
				}
			case 3:
				os.Exit(1)
			}
		})
}

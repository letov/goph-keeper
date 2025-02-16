package page

import (
	"os"

	"github.com/rivo/tview"
)

func NewWelcomeModal(pages *tview.Pages) *tview.Modal {
	return tview.NewModal().
		SetText("Welcome to Goph Password Keeper").
		AddButtons([]string{"Login", "Register", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonIndex {
			case 0:
				pages.SwitchToPage(LoginPage)
			case 1:
				pages.SwitchToPage(RegisterPage)
			case 2:
				os.Exit(1)
			}
		})
}

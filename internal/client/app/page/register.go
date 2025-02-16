package page

import (
	"github.com/rivo/tview"
)

func NewRegisterForm(
	pages *tview.Pages,
	registerCallback func(email string, password string) error,
) *tview.Form {
	email := ""
	password := ""

	form := tview.NewForm().
		AddInputField("Email", "", 30, nil, func(text string) { email = text }).
		AddPasswordField("Password", "", 30, '*', func(text string) { password = text })

	form = form.AddButton("Register", func() {
		if err := registerCallback(email, password); err == nil {
			clearForm(form, []string{"Email", "Password"})
			pages.SwitchToPage(WelcomePage)
		}
	}).AddButton("Back", func() {
		clearForm(form, []string{"Email", "Password"})
		pages.SwitchToPage(WelcomePage)
	})
	form.SetBorder(true).SetTitle("Register").SetTitleAlign(tview.AlignLeft)

	return form
}

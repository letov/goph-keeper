package page

import (
	"github.com/rivo/tview"
)

func NewLoginForm(
	pages *tview.Pages,
	loginCallback func(email string, password string) error,
) *tview.Form {
	email := ""
	password := ""

	form := tview.NewForm().
		AddInputField("Email", "", 30, nil, func(text string) { email = text }).
		AddPasswordField("Password", "", 30, '*', func(text string) { password = text })

	form = form.
		AddButton("Login", func() {
			if err := loginCallback(email, password); err == nil {
				clearForm(form, []string{"Email", "Password"})
				pages.SwitchToPage(SelectDBPage)
			}
		}).
		AddButton("Back", func() {
			clearForm(form, []string{"Email", "Password"})
			pages.SwitchToPage(WelcomePage)
		})
	form.SetBorder(true).SetTitle("Login").SetTitleAlign(tview.AlignLeft)

	return form
}

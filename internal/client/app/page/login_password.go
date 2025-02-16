package page

import (
	"GophKeeper/internal/common/dto"

	"github.com/rivo/tview"
)

func NewLoginPassword(
	pages *tview.Pages,
	e dto.LoginPassword,
) {
	pageName := "login_password"
	form := tview.NewForm().
		AddInputField("Meta", string(e.Meta), 30, nil, nil).
		AddInputField("Login", string(e.Login), 30, nil, nil).
		AddInputField("Password", string(e.Password), 30, nil, nil)

	form = form.
		AddButton("Save", func() {}).
		AddButton("Back", func() {
			pages.RemovePage(pageName).SwitchToPage(PrivateListPage)
		})
	form.SetBorder(true).SetTitle("Login / Password").SetTitleAlign(tview.AlignLeft)

	pages.AddAndSwitchToPage(
		pageName,
		form,
		true,
	)

	return
}

package page

import (
	"GophKeeper/internal/common/dto"

	"github.com/rivo/tview"
)

func NewBankCard(
	pages *tview.Pages,
	e dto.BankCard,
) {
	pageName := "bank_card"
	form := tview.NewForm().
		AddInputField("Meta", string(e.Meta), 30, nil, nil).
		AddInputField("Number", string(e.Number), 30, nil, nil).
		AddInputField("Date", string(e.Date), 30, nil, nil).
		AddInputField("Cvv", string(e.Cvv), 30, nil, nil)

	form = form.
		AddButton("Save", func() {}).
		AddButton("Back", func() {
			pages.RemovePage(pageName).SwitchToPage(PrivateListPage)
		})
	form.SetBorder(true).SetTitle("Bank Card").SetTitleAlign(tview.AlignLeft)

	pages.AddAndSwitchToPage(
		pageName,
		form,
		true,
	)

	return
}

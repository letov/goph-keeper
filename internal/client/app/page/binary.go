package page

import (
	"GophKeeper/internal/common/dto"

	"github.com/rivo/tview"
)

func NewBinary(
	pages *tview.Pages,
	e dto.Binary,
) {
	pageName := "binary"
	form := tview.NewForm().
		AddInputField("Meta", string(e.Meta), 30, nil, nil).
		AddInputField("Binary", string(e.Binary), 30, nil, nil)

	form = form.
		AddButton("Save", func() {}).
		AddButton("Back", func() {
			pages.RemovePage(pageName).SwitchToPage(PrivateListPage)
		})
	form.SetBorder(true).SetTitle("Binary").SetTitleAlign(tview.AlignLeft)

	pages.AddAndSwitchToPage(
		pageName,
		form,
		true,
	)

	return
}

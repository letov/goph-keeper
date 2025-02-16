package page

import "github.com/rivo/tview"

type PageName = string

const (
	WelcomePage     PageName = "welcome"
	LoginPage       PageName = "login"
	RegisterPage    PageName = "register"
	ErrorPage       PageName = "error"
	SelectDBPage    PageName = "select_db"
	PrivateListPage PageName = "private_list_page"
)

func clearForm(form *tview.Form, labels []string) {
	for _, l := range labels {
		i := form.GetFormItemIndex(l)
		if i == -1 {
			continue
		}
		item, ok := form.GetFormItem(i).(*tview.InputField)
		if ok {
			item.SetText("")
		}
	}

}

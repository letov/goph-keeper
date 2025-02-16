package page

import (
	"GophKeeper/internal/client/app/repo"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewPrivateList(
	pages *tview.Pages,
	sr repo.Snapshot,
) *tview.TreeView {
	root := tview.NewTreeNode("Root").
		SetSelectable(false).
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	ss, _ := sr.GetSnapshot()

	node := tview.NewTreeNode("Login / Password").
		SetSelectable(false).
		SetColor(tcell.ColorDarkGreen)
	for _, e := range ss.LoginPasswordList {
		n := tview.NewTreeNode(string(e.Meta)).
			SetSelectedFunc(func() {
				NewLoginPassword(pages, e)
			})
		node.AddChild(n)
	}
	root.AddChild(node)

	node = tview.NewTreeNode("Binary").
		SetSelectable(false).
		SetColor(tcell.ColorDarkGreen)
	for _, e := range ss.BinaryList {
		n := tview.NewTreeNode(string(e.Meta)).
			SetSelectedFunc(func() {
				NewBinary(pages, e)
			})
		node.AddChild(n)
	}
	root.AddChild(node)

	node = tview.NewTreeNode("Bank Card").
		SetSelectable(false).
		SetColor(tcell.ColorDarkGreen)
	for _, e := range ss.BankCardList {
		n := tview.NewTreeNode(string(e.Meta)).
			SetSelectedFunc(func() {
				NewBankCard(pages, e)
			})
		node.AddChild(n)
	}
	root.AddChild(node)

	return tree
}

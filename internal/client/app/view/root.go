package view

import (
	"GophKeeper/internal/client/app/handler"
	"GophKeeper/internal/client/app/page"
	"GophKeeper/internal/client/app/repo"
	"context"

	"github.com/rivo/tview"
	"go.uber.org/zap"
)

type Root struct {
	Root *tview.Application
}

func NewRoot(
	log zap.SugaredLogger,

	sr repo.Snapshot,

	lh *handler.Login,
	rh *handler.Register,
	sdbh *handler.SelectDB,
) *Root {
	root := tview.NewApplication()

	ctx := context.Background()
	pages := tview.NewPages()
	errorModal := page.NewAlertModal()

	getErrorCallback := func(p page.PageName) func(errorMsg string) {
		return func(errorMsg string) {
			errorModal.SetText(errorMsg).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonIndex == 0 {
						pages.SwitchToPage(p)
					}
				})
			pages.SwitchToPage(page.ErrorPage)
		}
	}

	pages.AddPage(
		page.WelcomePage,
		page.NewWelcomeModal(pages),
		true,
		true,
	)

	pages.AddPage(
		page.LoginPage,
		page.NewLoginForm(pages, lh.GetHandler(ctx, getErrorCallback(page.LoginPage))),
		true,
		false,
	)

	pages.AddPage(
		page.RegisterPage,
		page.NewRegisterForm(pages, rh.GetHandler(ctx, getErrorCallback(page.RegisterPage))),
		true,
		false,
	)

	pages.AddPage(
		page.SelectDBPage,
		page.NewSelectDBModal(pages, sdbh.GetHandler(ctx, getErrorCallback(page.SelectDBPage))),
		true,
		false,
	)

	pages.AddPage(
		page.ErrorPage,
		errorModal,
		true,
		false,
	)

	pages.AddPage(
		page.PrivateListPage,
		page.NewPrivateList(pages, sr),
		true,
		false,
	)

	if err := root.SetRoot(pages, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		log.Error(err)
	}

	return &Root{
		Root: root,
	}
}

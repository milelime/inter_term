package ui

import (
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/milelime/inter_term/auth"
	"github.com/rivo/tview"
)

const (
	IS_DEBUG = false
)

func IsConnected() (bool, error) {
	if auth.IS_DEBUG {
		return true, nil
	}
	resp, err := http.Get(auth.API_URL + "/health")
	if err != nil {
		return false, err
	}
	return resp.Status == "200", nil
}

func createConnectionRetryModal(app *tview.Application) *tview.Modal {
	return tview.NewModal().
		SetText("You are not connected to the API. Would you like to retry?").
		AddButtons([]string{"Retry", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Retry" {
				tryConnection(app)
			} else {
				app.Stop()
			}
		})
}

func createWelcomeScreen(app *tview.Application) *tview.Flex {
	// Main container
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	// Title
	title := tview.NewTextView().
		SetText("interterm").
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorGreen)

	// Form
	form := tview.NewForm().
		AddInputField("Full Name", "", 30, nil, nil).
		AddPasswordField("Interviewer Passkey", "", 20, '*', nil).
		AddButton("Submit", func() {
			// TODO: Handle authentication
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	form.SetBorder(true).
		SetTitle(" Login ").
		SetTitleAlign(tview.AlignLeft)

	// Layout
	flex.AddItem(nil, 0, 1, false). // Spacing
					AddItem(title, 3, 1, false).
					AddItem(nil, 0, 1, false). // Spacing
					AddItem(form, 0, 2, true)

	return flex
}

func tryConnection(app *tview.Application) {
	isConnected, err := IsConnected()
	if !isConnected || err != nil {
		app.SetRoot(createConnectionRetryModal(app), true)
		return
	}
	app.SetRoot(createWelcomeScreen(app), true)
}

func Start() {
	app := tview.NewApplication()
	tryConnection(app)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

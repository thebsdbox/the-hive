package challenges

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SelectChallenge() (*Challenge, error) {
	var app = tview.NewApplication()
	var flex = tview.NewFlex()
	var contactsList = tview.NewList().ShowSecondaryText(false).SetMainTextColor(tcell.ColorGreen)
	var contactText = tview.NewTextView()
	var pages = tview.NewPages()

	var text = tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("(Enter) to select a challenge")

	contactsList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		setSelection(contactText, index)
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(contactsList, 0, 1, true).
			AddItem(contactText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	pages.AddPage("Menu", flex, true, true)

	contactsList.Clear()

	for index, challenge := range Challenges {
		contactsList.AddItem(challenge.Name, "", rune(49+index), nil)
	}

	setSelection(contactText, 0)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		currentSelection := contactsList.GetCurrentItem()
		switch event.Key() {
		case tcell.KeyUp:
			if currentSelection != 0 {
				contactsList.SetCurrentItem(currentSelection - 1)
			}
		case tcell.KeyDown:
			if currentSelection < len(Challenges) {
				contactsList.SetCurrentItem(currentSelection + 1)
			}

		case tcell.KeyEnter:
			app.Stop()
			// case tcell.KeyCtrlC:
			// 	app.Stop()
			//cancel = true
		}
		setSelection(contactText, contactsList.GetCurrentItem())
		return event
	})
	// This blocks
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		return nil, err
	}
	return &Challenges[contactsList.GetCurrentItem()], nil
}

func setSelection(contactText *tview.TextView, index int) {
	contactText.Clear()
	text := Challenges[index].Description + "\n " + "Time to complete:" + Challenges[index].AllowedTime.String()
	contactText.SetText(text)
}

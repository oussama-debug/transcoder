package graphics

import (
	"github.com/rivo/tview"
)

var (
	app  *tview.Application
	flex *tview.Flex
)

func Start() {
	app = tview.NewApplication()
	DrawHome()
}

func DrawHome() {
	flex = tview.NewFlex().SetDirection(tview.FlexRow).AddItem(CreateHelperSection(), 4, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

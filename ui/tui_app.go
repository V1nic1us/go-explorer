package ui

import (
	"go-explorer/controller"
	"go-explorer/model"
	"go-explorer/view"
	"github.com/rivo/tview"
)

func RunTUI() {
	app := tview.NewApplication()
	v := view.NewExplorerView()
	m, err := model.NewFileModel("")
	if err != nil {
		panic(err)
	}
	c := controller.NewExplorerController(app, v, m)
	c.Start()
}

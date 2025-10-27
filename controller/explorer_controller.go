package controller

import (
	"fmt"
	"io/ioutil"
	"github.com/rivo/tview"
	"go-explorer/model"
	"go-explorer/view"
)

type ExplorerController struct {
	App  *tview.Application
	View *view.ExplorerView
	Mod  *model.FileModel
}

func NewExplorerController(app *tview.Application, v *view.ExplorerView, m *model.FileModel) *ExplorerController {
	c := &ExplorerController{App: app, View: v, Mod: m}
	c.initBindings()
	return c
}

func (c *ExplorerController) initBindings() {
	c.View.SetKeyBindings(c.App, func(r rune) {
		switch r {
		case 'q':
			c.App.Stop()
		case 'r':
			c.reload()
		case 'd':
			c.delete()
		case 'n':
			c.create()
		case 'u':
			c.goUp()
		case 'e':
			c.enter()
		}
	})
}

func (c *ExplorerController) Start() {
	c.reload()
	if err := c.App.SetRoot(c.View.Root(), true).Run(); err != nil {
		panic(err)
	}
}

func (c *ExplorerController) reload() {
	c.View.RenderFiles(c.Mod.CurrentPath, c.Mod.Files)
	c.View.ShowPreview("[yellow]Use ↑↓ para navegar, 'e' para entrar, 'u' para voltar, 'q' para sair, 'd' para deletar, 'n' para criar")
}

func (c *ExplorerController) enter() {
	node := c.View.Tree.GetCurrentNode()
	if node == nil {
		return
	}
	ref := node.GetReference()
	if ref == nil {
		return
	}
	file := ref.(model.FileItem)
	if !file.IsDir {
		content, _ := ioutil.ReadFile(file.Path)
		c.View.ShowPreview(string(content))
		return
	}
	if err := c.Mod.EnterDir(file.Name); err != nil {
		c.View.ShowPreview(fmt.Sprintf("[red]Erro: %v", err))
	} else {
		c.reload()
	}
}

func (c *ExplorerController) goUp() {
	if err := c.Mod.GoUp(); err != nil {
		c.View.ShowPreview(fmt.Sprintf("[red]Erro: %v", err))
	} else {
		c.reload()
	}
}

func (c *ExplorerController) delete() {
	node := c.View.Tree.GetCurrentNode()
	if node == nil {
		return
	}
	ref := node.GetReference()
	if ref == nil {
		return
	}
	file := ref.(model.FileItem)
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Deletar %s?", file.Name)).
		AddButtons([]string{"Sim", "Não"}).
		SetDoneFunc(func(i int, label string) {
			if label == "Sim" {
				_ = c.Mod.Delete(file.Name)
				c.reload()
			}
			c.App.SetRoot(c.View.Root(), true)
		})
	c.App.SetRoot(modal, false)
}

func (c *ExplorerController) create() {
	form := tview.NewForm()
	form.AddInputField("Nome", "", 30, nil, nil).
		AddDropDown("Tipo", []string{"Arquivo", "Diretório"}, 0, nil).
		AddButton("Criar", func() {
			name := form.GetFormItemByLabel("Nome").(*tview.InputField).GetText()
			_, tipo := form.GetFormItemByLabel("Tipo").(*tview.DropDown).GetCurrentOption()
			c.Mod.Create(name, tipo == "Diretório")
			c.reload()
			c.App.SetRoot(c.View.Root(), true)
		}).
		AddButton("Cancelar", func() {
			c.App.SetRoot(c.View.Root(), true)
		})
	form.SetBorder(true).SetTitle(" Novo arquivo/diretório ")
	c.App.SetRoot(form, true)
}

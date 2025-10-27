package view

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"go-explorer/model"
)

type ExplorerView struct {
	Tree     *tview.TreeView
	Preview  *tview.TextView
	Layout   *tview.Flex
}

func NewExplorerView() *ExplorerView {
	tree := tview.NewTreeView()
		// SetBorder(true).
		// SetTitle(" Files ")

	preview := tview.NewTextView()
		// SetDynamicColors(true).
		// SetBorder(true).
		// SetTitle(" Preview ")

	flex := tview.NewFlex().
		AddItem(tree, 0, 1, true).
		AddItem(preview, 0, 2, false)

	return &ExplorerView{Tree: tree, Preview: preview, Layout: flex}
}

func (v *ExplorerView) RenderFiles(path string, files []model.FileItem) {
	root := tview.NewTreeNode(fmt.Sprintf("[yellow]%s", path)).SetExpanded(true)
	for _, f := range files {
		name := f.Name
		if f.IsDir {
			name = fmt.Sprintf("[green]%s/", f.Name)
		}
		child := tview.NewTreeNode(name).SetReference(f)
		root.AddChild(child)
	}
	v.Tree.SetRoot(root).SetCurrentNode(root)
}

func (v *ExplorerView) ShowPreview(content string) {
	v.Preview.SetText(content).
	SetDynamicColors(true)
}

func (v *ExplorerView) Root() tview.Primitive {
	return v.Layout
}

func (v *ExplorerView) SetKeyBindings(app *tview.Application, onKey func(r rune)) {
	v.Tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune {
			onKey(event.Rune())
			return nil
		}
		return event
	})
}

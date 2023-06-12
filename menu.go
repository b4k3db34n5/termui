package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Initializes the program's menu, return error
func RunTerm() error {
	if err := ui.Init(); err != nil {
		return fmt.Errorf("failed to initialize termui")
	}
	defer ui.Close()

	menuList := widgets.NewList()
	menuList.Title = " MENU "
	menuList.Rows = []string{
		"[STOCKS](fg:green)",
		"[CRYPTO](fg:yellow)",
		"[EXIT](fg:red)",
	}
	menuList.WrapText = false
	menuList.TitleStyle.Fg = ui.ColorCyan
	menuList.BorderStyle.Fg = ui.ColorCyan
	menuList.SelectedRowStyle.Modifier = ui.ModifierBold
	menuList.SetRect(0, 0, 25, 5)

	ui.Render(menuList)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q":
			return nil
		case "<Down>":
			menuList.ScrollDown()
		case "<Up>":
			menuList.ScrollUp()
		}
		ui.Render(menuList)
	}
}

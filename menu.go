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
		"[STOCK INFO](fg:blue)",
		"[EXIT](fg:red)",
	}
	menuList.WrapText = false
	menuList.BorderStyle.Fg = ui.ColorBlue
	menuList.SetRect(0, 0, 25, 4)
	menuList.SelectedRowStyle.Fg = ui.ColorGreen

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

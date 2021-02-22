package ui

import (
	"fmt"

	"github.com/gdamore/tcell"
)

func Start() {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("Error making screen: %+v", err)
	}

	err = screen.Init()
	if err != nil {
		return nil, fmt.Errorf("Error initing screen screen: %+v", err)
	}
}

package ui

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/sirupsen/logrus"
)

type screen struct {
	screen tcell.Screen
}

func newScreen() (*screen, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("Error making screen: %+v", err)
	}

	err = s.Init()
	if err != nil {
		return nil, fmt.Errorf("Error initing screen screen: %+v", err)
	}

	return &screen{screen: s}, nil
}

func (s *screen) destroy() {
	s.screen.Fini()
}

func (s *screen) clear() {
	s.screen.Clear()
}

func (s *screen) flush() {
	s.screen.Show()
}

type CommandType string

const (
	OpenBrowser CommandType = "openBrowser"
)

func (s *screen) listen() <-chan CommandType {
	commands := make(chan CommandType, 5)

	go func() {
		for {
			event := s.screen.PollEvent()
			switch e := event.(type) {
			case *tcell.EventKey:
				switch e.Key() {
				case tcell.KeyRune:
					if e.Rune() == ' ' {
						commands <- OpenBrowser
					}
				}
			case nil:
				logrus.Debug("Keyboard listener loop stopped: screen finalised")
				return
			}
		}
	}()

	return commands
}

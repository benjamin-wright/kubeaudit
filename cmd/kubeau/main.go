package main

import (
	"sync"

	"github.com/benjamin-wright/kubeaudit/internal/server"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go server.Serve(&wg)

	wg.Wait()
}

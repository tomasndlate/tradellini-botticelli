package handlers

import (
	"fmt"
	"sync"
	"time"
)

var (
	stopChannel chan struct{}
	mutex       sync.Mutex
	running     bool
)

func StartPrint(name string) {
	mutex.Lock()
	defer mutex.Unlock()

	stopChannel = make(chan struct{})
	running = true

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("Hello, " + name)
			case <-stopChannel:
				return
			}
		}
	}()
}

func StopPrint() {
	mutex.Lock()
	defer mutex.Unlock()

	close(stopChannel)
	running = false
}

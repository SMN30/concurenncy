package scheduler

import (
	"sync"
	"time"
)

// Every запускает f каждые d и возвращает функцию для остановки.
func Every(d time.Duration, f func()) (stop func()) {
	// TODO: периодический вызов функции с возможностью остановки
	done := make(chan struct{})

	var once sync.Once
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				f()
			case <-done:
				return
			}
		}
	}()

	return func() {
		once.Do(func() {
			close(done)
		})
	}
}

package scheduler

import (
	"sync"
	"time"
)

// Every запускает f каждые d и возвращает функцию для остановки.
func Every(d time.Duration, f func()) (stop func()) {
	// TODO: периодический вызов функции с возможностью остановки
	// Канал для сигнала об остановке
	done := make(chan struct{})

	var once sync.Once
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop() // Важно остановить тикер для освобождения ресурсов

		for {
			select {
			case <-ticker.C:
				f()
			case <-done:
				return
			}
		}
	}()

	// Возвращаем замыкание, которое сигнализирует горутине о выходе
	return func() {
		once.Do(func() {
			close(done)
		})
	}
}

package scheduler

import "time"

// Every запускает f каждые d и возвращает функцию для остановки.
func Every(d time.Duration, f func()) (stop func()) {
	// TODO: периодический вызов функции с возможностью остановки
	// Канал для сигнала об остановке
	done := make(chan struct{})

	// Создаем тикер
	ticker := time.NewTicker(d)

	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			case <-done:
				ticker.Stop() // Важно остановить тикер, чтобы освободить ресурсы
				return
			}
		}
	}()

	// Возвращаем функцию, которая при вызове закроет канал done
	return func() {
		close(done)
	}
}

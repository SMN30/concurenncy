package debounce

import (
	"time"
)

// Debounce принимает значения и отдаёт только последнее после паузы d.
func Debounce(d time.Duration, in <-chan int) <-chan int {
	// TODO: реализовать дебаунс значений из канала
	out := make(chan int)

	go func() {
		defer close(out)

		var value int
		timer := time.NewTimer(d)
		// Сразу останавливаем, так как данных еще нет
		if !timer.Stop() {
			<-timer.C
		}

		var active bool // Флаг, есть ли у нас значение, которое ждет отправки

		for {
			select {
			case v, ok := <-in:
				if !ok {
					if active {
						// Если канал закрылся, но есть значение — ждем таймер
						<-timer.C
						out <- value
					}
					return
				}

				value = v
				active = true

				// Перезапускаем таймер. Stop + Reset — стандарт Go.
				timer.Stop()
				select {
				case <-timer.C:
				default:
				}
				timer.Reset(d)

			case <-timer.C:
				out <- value
				active = false // Значение отправлено, больше ничего не ждем
			}
		}
	}()

	return out
}

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
		if !timer.Stop() {
			<-timer.C
		}

		var active bool

		for {
			select {
			case v, ok := <-in:
				if !ok {
					if active {
						<-timer.C
						out <- value
					}
					return
				}

				value = v
				active = true

				timer.Stop()
				select {
				case <-timer.C:
				default:
				}
				timer.Reset(d)

			case <-timer.C:
				out <- value
				active = false
			}
		}
	}()

	return out
}

package limiter

import (
	"sync"
	"time"
)

// Limiter ограничивает количество событий до 5 в секунду.
type Limiter struct {
	tokens chan struct{}
	done   chan struct{}
	once   sync.Once
}

// NewLimiter создаёт новый лимитер с ёмкостью 5 токенов.
func NewLimiter() *Limiter {
	// TODO: инициализировать канал токенов и запуск пополнения
	l := &Limiter{
		// Канал с буфером на 5 токенов
		tokens: make(chan struct{}, 5),
		done:   make(chan struct{}),
	}

	// Сразу заполняем "бак" начальными токенами
	for i := 0; i < 5; i++ {
		l.tokens <- struct{}{}
	}

	// Запускаем фоновое пополнение: 1 токен каждые 200мс (итого 5 в сек)
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
				case l.tokens <- struct{}{}:
					// Токен добавлен
				default:
					// Бак полон, пропускаем тик
				}
			case <-l.done:
				return
			}
		}
	}()

	return l
}

// Allow возвращает true, если событие разрешено в текущий момент.
func (l *Limiter) Allow() bool {
	// TODO: реализовать получение токена из канала
	select {
	case <-l.tokens:
		return true
	default:
		return false
	}
}

// Stop останавливает лимитер.
func (l *Limiter) Stop() {
	// TODO: остановить пополнение токенов
	l.once.Do(func() {
		close(l.done)
	})
}

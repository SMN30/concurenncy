package generator

import "context"

// Generate возвращает канал, из которого можно читать возрастающие числа,
// начиная с нуля. Генерация прекращается при отмене ctx.
func Generate(ctx context.Context) <-chan int {
	// TODO: реализовать генератор чисел с учётом отмены
	ch := make(chan int)

	go func() {
		defer close(ch)

		var n int
		for {
			// Неблокирующая проверка отмены контекста перед попыткой отправки.
			// Это гарантирует мгновенный выход и закрытие канала через defer,
			// если контекст уже был отменен (до старта или в процессе).
			select {
			case <-ctx.Done():
				return
			default:
			}

			// Основной блокирующий select для отправки данных
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}

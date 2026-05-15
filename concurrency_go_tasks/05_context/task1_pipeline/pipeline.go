package pipelinectx

import "context"

// Run строит конвейер из двух стадий: удвоение и суммирование.
// Конвейер должен останавливаться, если ctx отменён.
// Возвращает итоговую сумму и ошибку контекста при отмене.
func Run(ctx context.Context, nums []int) (int, error) {
	// TODO: реализовать конвейер с остановкой по ctx

	doubled := make(chan int)

	go func() {
		defer close(doubled)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case doubled <- n * 2:
			}
		}
	}()

	total := 0
	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case val, ok := <-doubled:
			if !ok {
				return total, nil
			}
			total += val
		}
	}

}

package fibonacci

// Fib возвращает канал, из которого можно читать первые n чисел Фибоначчи.
func Fib(n int) <-chan int {
	ch := make(chan int)
	// TODO: отправить последовательность Фибоначчи в канал
	go func() {
		defer close(ch)

		var n1 int = 0
		var n2 int = 1

		for i := 0; i < n; i++ {
			tmp := n1
			n1 = n2
			n2 = tmp + n2
			ch <- tmp
		}
	}()

	return ch
}

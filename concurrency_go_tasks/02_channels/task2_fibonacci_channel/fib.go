package fibonacci

// Fib возвращает канал, из которого можно читать первые n чисел Фибоначчи.
func Fib(n int) <-chan int {
	ch := make(chan int)
	// TODO: отправить последовательность Фибоначчи в канал
	go func() {
		defer close(ch)

		n1 := 0
		n2 := 1

		for i := 0; i < n; i++ {
			tmp := n1
			n1 = n2
			n2 = tmp + n2
			ch <- tmp
		}
	}()

	return ch
}

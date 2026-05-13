package pipeline

// Run строит конвейер из трёх стадий: квадрат, умножение на 2 и суммирование.
func Run(nums []int) int {
	// TODO: реализовать конвейер обработки чисел
	in := make(chan int)
	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()

	sq := make(chan int)
	go func() {
		for n := range in {
			sq <- n * n
		}
		close(sq)
	}()

	mult := make(chan int)
	go func() {
		for n := range sq {
			mult <- n * 2
		}
		close(mult)
	}()

	sum := 0
	for i := range mult {
		sum += i
	}

	return sum
}

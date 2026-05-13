package pool

import (
	"sync"
	_ "sync"
)

// RunPool обрабатывает задачи параллельно в заданном количестве воркеров
// и возвращает сумму результатов.
func RunPool(jobs []int, workers int) int {
	// TODO: реализовать пул воркеров и сбор результатов
	if len(jobs) == 0 {
		return 0
	}

	if workers <= 0 {
		sum := 0
		for _, v := range jobs {
			sum += v
		}
		return sum
	}

	// Канал для раздачи задач
	jobsChan := make(chan int, len(jobs))
	// Канал для сбора результатов
	resultsChan := make(chan int, len(jobs))

	var wg sync.WaitGroup

	// 1. Запускаем воркеров
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobsChan {
				// В данной задаче "обработка" — это просто проброс числа
				// (или можно добавить логику, если нужно что-то вычислить)
				resultsChan <- j
			}
		}()
	}

	// 2. Раздаем задачи
	for _, j := range jobs {
		jobsChan <- j
	}
	close(jobsChan) // Сигнализируем воркерам, что задач больше нет

	// 3. Ждем завершения всех воркеров в отдельной горутине,
	// чтобы закрыть канал результатов и не поймать deadlock.
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// 4. Суммируем результаты
	totalSum := 0
	for res := range resultsChan {
		totalSum += res
	}

	return totalSum
}

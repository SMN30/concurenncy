package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// PingPong должен запускать две горутины "ping" и "pong",
// которые поочередно выводят строки пять раз каждая.
// Реализуйте синхронизацию через каналы и ожидание завершения.
func PingPong(w io.Writer) {
	// TODO: реализовать обмен сообщениями между горутинами
	var wg sync.WaitGroup
	wg.Add(2)
	ping := make(chan bool)
	pong := make(chan bool)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-ping
			fmt.Fprintln(w, "ping")
			pong <- true

		}
		//	fmt.Println("Горутина Ping завершилась")
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-pong
			fmt.Fprintln(w, "pong")
			if i < 4 {
				ping <- true
			}
		}
		//	fmt.Println("Горутина Pong завершилась")
	}()

	ping <- true
	wg.Wait()
}

func main() {
	PingPong(os.Stdout)
	//wg.Wait()
	//	time.Sleep(time.Second * 3)
}

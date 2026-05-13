package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// PingPong должен запускать две горутины "ping" и "pong",
// которые поочередно выводят строки пять раз каждая.
// Реализуйте синхронизацию через каналы и ожидание завершения.
func PingPong(w io.Writer) {
	// TODO: реализовать обмен сообщениями между горутинами

	ping := make(chan bool)
	pong := make(chan bool)

	for i := 0; i < 5; i++ {
		go func() {
			<-ping
			fmt.Println("Ping")
			pong <- true
		}()
	}
	for j := 0; j < 5; j++ {
		go func() {
			<-pong
			fmt.Println("Pong")
			ping <- true
		}()
	}
	ping <- true
}

func main() {
	PingPong(os.Stdout)
	time.Sleep(time.Second * 3)
}

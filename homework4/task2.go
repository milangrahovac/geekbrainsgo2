package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Task 2")
	// Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее, чем за одну секунду (установить таймаут).

	signals := make(chan os.Signal, 1)
	// done := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	signal := <-signals

	log.Printf("Got signal %v, exiting...", signal)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	shutdown(ctx)

	fmt.Println("exiting")
}

// shutdown эмулирует завершение программы, как будто нужно завершать какие-то асинхронные процессы
func shutdown(ctx context.Context) {
	done := make(chan struct{}, 1)

	go func() {
		log.Println("exiting...")
		// как будто здесь происходят долгие действия по завершению программы
		time.Sleep(2 * time.Second)
		log.Println("shutdown is done")
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("forced shutdown")
	case <-done:
		return
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func incCounter(ctx context.Context, counter *int) {

	t := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-t.C:
			*counter = *counter + 1
		case <-ctx.Done():
			log.Println("stopping incrementor routine")
			t.Stop()
			return
		}
	}
}

func startServer(counter *int) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Hey, ahh hmmmm")
	})

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Current count : %d\n", counter)
	})

	http.ListenAndServe(":8081", nil)
}

func main() {

	var counter = 0

	ctx, cancel := context.WithCancel(context.Background())
	go incCounter(ctx, &counter)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go startServer(&counter)
	log.Println("server started. listerning at :8081")
	<-done
	cancel()
	log.Println("server stopped")
}

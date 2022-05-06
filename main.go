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

var start time.Time

func incCounter(ctx context.Context, counter *int64) {

	t := time.NewTicker(3 * time.Second)

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

func startServer(counter *int64) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Hey, ahh hmmmm")
	})

	http.HandleFunc("/start-time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Started at : %s\n", start)
	})

	http.HandleFunc("/restart-counter", func(w http.ResponseWriter, r *http.Request) {
		//FIXME: handle race
		*counter = 0
		fmt.Fprintln(w, "counter restarted")
	})

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Current count : %d\n", *counter)
	})

	http.ListenAndServe(":8081", nil)
}

func main() {

	var counter = int64(0)
	loc, _ := time.LoadLocation("Asia/Kolkata")
	start = time.Now().In(loc)

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

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"working/handlers"
)

func main() {
	/*	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Hello world")
			d, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(rw, "Oops an error", http.StatusBadRequest)
				return
			}

			fmt.Fprintf(rw, "Hello %s", d)
		})

		http.HandleFunc("/goodbye", func(_w http.ResponseWriter, _r *http.Request) {
			log.Println("Goodbye world")
		})
		http.ListenAndServe(":9999", nil)*/

	l := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9998",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	// Wait 30 seconds to finish all tasks
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

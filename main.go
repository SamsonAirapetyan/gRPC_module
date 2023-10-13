package main

import (
	"context"
	"fmt"
	"gRPC/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "gRPC", log.LstdFlags)
	ph := handler.NewProduct(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	fmt.Println("Server is running...")
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)
	sig := <-sigchan
	l.Println("Server is interrupted", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	s.ListenAndServe()
}

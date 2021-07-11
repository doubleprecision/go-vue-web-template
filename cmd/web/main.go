package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"yournal/pkg/router"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "timeout", time.Second*5, "duration to wait for gracefully exit")
	port := flag.Int("port", 8000, "port to serve API on")
	flag.Parse()
	addr := fmt.Sprintf("0.0.0.0:%d", *port)

	r := router.GetRouters()

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		log.Printf("Server start on http://%s\n", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}

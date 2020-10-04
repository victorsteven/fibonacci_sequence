package main

import (
	"context"
	"fibonacci_sequence/domain"
	"fibonacci_sequence/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	appAddr := ":" + "8080"

	r := gin.Default()

	fibService := domain.NewFibonacciService()

	handlerService := handler.NewHandlerService(fibService)

	r.GET("/current", handlerService.CurrentNumber)
	r.GET("/next", handlerService.NextNumber)
	r.GET("/previous", handlerService.PreviousNumber)

	//Starting and Shutting down Server
	srv := &http.Server{
		Addr:    appAddr,
		Handler: r,
	}

	go func() {
		//service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

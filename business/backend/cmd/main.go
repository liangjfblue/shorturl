package main

import (
	"context"
	"log"
	"os"
	"time"
)

func main() {
	server, cleanup, err := wireServer()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	err = server.Run()
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("shutdown internal ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Stop(ctx); err != nil {
		panic(err)
	}
}

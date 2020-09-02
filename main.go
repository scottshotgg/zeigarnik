package main

import (
	"context"
	"log"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/server/rest"
	"github.com/scottshotgg/zeigarnik/pkg/server/rpc"
	"github.com/scottshotgg/zeigarnik/pkg/storage/mem"
)

const (
	rpcPort  = 5001
	restPort = 8080
)

func main() {
	var err = start()
	if err != nil {
		log.Fatalln("error:", err)
	}
}

func start() error {
	var (
		store       = mem.New(nil)
		errChan     = make(chan error)
		ctx, cancel = context.WithCancel(context.Background())
	)

	go func() {
		log.Println("Starting RPC")
		errChan <- rpc.Start(ctx, rpcPort, store)
	}()

	go func() {
		log.Println("Starting Rest")
		rest.Start(ctx, restPort, rpcPort)
	}()

	time.Sleep(1 * time.Millisecond)

	log.Println("Up!")

	for err := range errChan {
		cancel()
		return err
	}

	return nil
}

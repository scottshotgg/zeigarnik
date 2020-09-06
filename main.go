package main

import (
	"context"
	"log"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/server/gql"
	"github.com/scottshotgg/zeigarnik/pkg/server/rest"
	"github.com/scottshotgg/zeigarnik/pkg/server/rpc"
	"github.com/scottshotgg/zeigarnik/pkg/server/swaggerui"
	"github.com/scottshotgg/zeigarnik/pkg/storage/mem"
)

const (
	rpcPort  = 5001
	restPort = 8080
	gqlPort  = 8888
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
		errChan <- rest.Start(ctx, restPort, rpcPort)
	}()

	go func() {
		log.Println("Starting GQL")
		errChan <- gql.Start(ctx, gqlPort)
	}()

	go func() {
		log.Println("Starting Swagger UI")
		errChan <- swaggerui.Start(ctx, 9090)
	}()

	time.Sleep(1 * time.Millisecond)

	log.Println("Up!")

	for err := range errChan {
		cancel()
		return err
	}

	return nil
}

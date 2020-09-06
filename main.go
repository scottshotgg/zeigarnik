package main

import (
	"context"
	"log"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/server/gql"
	"github.com/scottshotgg/zeigarnik/pkg/server/rest"
	"github.com/scottshotgg/zeigarnik/pkg/server/rpc"
	"github.com/scottshotgg/zeigarnik/pkg/storage"
	"github.com/scottshotgg/zeigarnik/pkg/storage/mem"
)

// TODO: make these env variables later
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

	defer close(errChan)
	defer cancel()

	// TODO: listen for ctrl-c

	go startRPC(ctx, errChan, store)
	go startREST(ctx, errChan)
	go startGQL(ctx, errChan)

	time.Sleep(1 * time.Millisecond)

	log.Println("Up!")

	for err := range errChan {
		cancel()
		return err
	}

	return nil
}

func startRPC(ctx context.Context, errChan chan<- error, store storage.Storage) {
	log.Println("Starting RPC on:", rpcPort)
	errChan <- rpc.Start(ctx, rpcPort, store)
}

func startREST(ctx context.Context, errChan chan<- error) {
	log.Println("Starting Rest on:", restPort)
	errChan <- rest.Start(ctx, restPort, rpcPort)
}

func startGQL(ctx context.Context, errChan chan<- error) {
	log.Println("Starting GQL on:", gqlPort)
	errChan <- gql.Start(ctx, gqlPort)
}

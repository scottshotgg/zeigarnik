package main

import (
	"log"

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
	var store = mem.New(nil)

	var errChan = make(chan error)

	go func() {
		errChan <- rpc.Start(rpcPort, store)
	}()

	go func() {
		rest.Start(restPort, rpcPort)
	}()

	for err := range errChan {
		return err
	}

	return nil
}

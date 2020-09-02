package rpc

import (
	"context"
	"net"
	"strconv"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
	"github.com/scottshotgg/zeigarnik/pkg/storage"
	"google.golang.org/grpc"
)

type (
	ReminderService struct {
		store storage.Storage
	}
)

func Start(ctx context.Context, port int, store storage.Storage) error {
	var lis, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	var (
		opts       = []grpc.ServerOption{}
		grpcServer = grpc.NewServer(opts...)
	)

	buffs.RegisterReminderServiceServer(grpcServer, &ReminderService{
		store: store,
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				grpcServer.Stop()
			}
		}
	}()

	return grpcServer.Serve(lis)
}

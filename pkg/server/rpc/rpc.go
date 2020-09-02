package rpc

import (
	"net"
	"strconv"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
	"github.com/scottshotgg/zeigarnik/pkg/storage"
	"google.golang.org/grpc"
)

type (
	ReminderService struct {
		s storage.Storage
	}
)

func Start(port int, s storage.Storage) error {
	var lis, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	var grpcServer = grpc.NewServer()

	buffs.RegisterReminderServiceServer(grpcServer, &ReminderService{
		s: s,
	})

	return grpcServer.Serve(lis)
}

package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
	"google.golang.org/grpc"
)

func makeAddr(port int) string {
	return ":" + strconv.Itoa(port)
}

func Start(restPort, rpcPort int) error {
	var (
		ctx = context.Background()

		// Register gRPC server endpoint
		// Note: Make sure the gRPC server is running properly and accessible
		mux  = runtime.NewServeMux()
		opts = []grpc.DialOption{
			grpc.WithInsecure(),
		}

		err = buffs.RegisterReminderServiceHandlerFromEndpoint(ctx, mux, makeAddr(rpcPort), opts)
	)

	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(makeAddr(restPort), mux)
}

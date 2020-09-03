package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
	"google.golang.org/grpc"
)

func makeAddr(port int) string {
	return ":" + strconv.Itoa(port)
}

func Start(ctx context.Context, restPort, rpcPort int) error {
	var (
		// Register gRPC server endpoint
		// Note: Make sure the gRPC server is running properly and accessible
		mux  = runtime.NewServeMux()
		opts = []grpc.DialOption{
			grpc.WithInsecure(),
		}

		err = reminder.RegisterReminderServiceHandlerFromEndpoint(ctx, mux, makeAddr(rpcPort), opts)
	)

	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(makeAddr(restPort), mux)
}

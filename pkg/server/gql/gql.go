package gql

import (
	"context"
	"strconv"

	"net/http"

	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func Start(ctx context.Context, port int) error {
	var (
		addr = ":" + strconv.Itoa(port)
		mux  = runtime.NewServeMux()
		err  = reminder.RegisterReminderServiceGraphql(mux)
		srv  = &http.Server{
			Addr: addr,
		}
	)

	if err != nil {
		return err
	}

	http.Handle("/graphql", mux)

	var errChan = make(chan error)

	// Monitor the context and shutdown the server if
	// it is canceled
	go func() {
		for {
			select {
			case <-ctx.Done():
				errChan <- srv.Shutdown(ctx)
			}
		}
	}()

	// Launch the server; return err if not ErrServerClosed
	go func() {
		var err = http.ListenAndServe(addr, nil)
		if err != nil {
			if err != http.ErrServerClosed {
				errChan <- err
			}
		}
	}()

	// If we get any error, return that
	for err := range errChan {
		return err
	}

	return nil
}

package swaggerui

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func makeAddr(port int) string {
	return ":" + strconv.Itoa(port)
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	http.ServeFile(w, r, "./external/swaggerui/dist/swagger.json")
}

func Start(ctx context.Context, port int) error {
	var (
		// Serve the swagger-ui and swagger file
		mux = http.NewServeMux()
		fs  = http.FileServer(http.Dir("./external/swaggerui/dist"))
	)

	mux.HandleFunc("/swagger.json", serveSwagger)
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui", fs))

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(makeAddr(port), mux)
}

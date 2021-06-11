package main

import (
	"context"
	"net/http"
	"os"

	pb "github.com/PutskouDzmitry/golang-training-library-grpc/proto/go_proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	server = os.Getenv("book-server")
	listen = os.Getenv("LISTEN")
)

func init() {
	if server == "" {
		server = "localhost:8080"
	}
	if listen == "" {
		listen = "localhost:8081"
	}
}

func main() {
	log := logrus.New()
	log.Debug("Client starts")
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	grpcMux := runtime.NewServeMux()
	err = pb.RegisterReaderHandler(context.Background(), grpcMux, conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(listen, grpcMux))
}

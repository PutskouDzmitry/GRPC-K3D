package main

import (
	"github.com/PutskouDzmitry/golang-training-library-grpc/book_server/pkg/api"
	"github.com/PutskouDzmitry/golang-training-library-grpc/book_server/pkg/const_db"
	"github.com/PutskouDzmitry/golang-training-library-grpc/book_server/pkg/data"
	"github.com/PutskouDzmitry/golang-training-library-grpc/book_server/pkg/db"
	pb "github.com/PutskouDzmitry/golang-training-library-grpc/proto/go_proto"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	listen   = os.Getenv("LISTEN")
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if listen == "" {
		listen = const_db.ServerPort
	}
	if host == "" {
		host = const_db.Host
	}
	if port == "" {
		port = const_db.Port
	}
	if user == "" {
		user = const_db.User
	}
	if dbname == "" {
		dbname = const_db.DbName
	}
	if password == "" {
		password = const_db.Password
	}
	if sslmode == "" {
		sslmode = const_db.Sslmode
	}
}

func main() {
	log := logrus.New()
	log.Debug("Server starts")
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}
	userData := data.NewBookData(conn)
	server := grpc.NewServer()
	pb.RegisterReaderServer(server, api.NewReadServer(userData))
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Debug("Server and database ready")
}
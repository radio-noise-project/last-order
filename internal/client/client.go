package client

import (
	"log"
	"strconv"

	"google.golang.org/grpc"
)

func Connect(addr string, port int) (*grpc.ClientConn, error) {
	address := addr + ":" + strconv.Itoa(port)
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal("Connection failed.")
		return nil, err
	}
	return conn, nil
}

package client

import (
	"context"

	"github.com/radio-noise-project/last-order/internal/client/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Version(conn *grpc.ClientConn) *runtime.VersionResponse {
	client := runtime.NewRuntimeServiceClient(conn)
	res, err := client.Version(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	} else {
		return res
	}
}

package client

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/radio-noise-project/last-order/internal/client/container"
	"google.golang.org/grpc"
)

func Run(conn *grpc.ClientConn, buf *bytes.Buffer, sisterId string) *container.StatusResponse {
	client := container.NewContainerServiceClient(conn)
	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatalf("could not upload: %v", err)
	}

	chunkSize := 1024
	for {
		chunk := make([]byte, chunkSize)
		n, err := buf.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read buffer: %v", err)
		}

		req := &container.UploadArchiveRequest{
			Archive: chunk[:n],
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send chunk: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}

	log.Printf("Upload response: %s", resp.GetStatus())
	return resp
}

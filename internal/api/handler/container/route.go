package container

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/radio-noise-project/last-order/internal/client"
	"github.com/radio-noise-project/last-order/internal/client/container"
	"github.com/radio-noise-project/last-order/internal/database"
	"github.com/radio-noise-project/last-order/pkg/database/model"
)

func runContainer(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to get file")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to open file")
	}
	defer src.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to convert to bytes")
	}
	sisterId := c.QueryParam("sisterId")
	err = sendTarballSisters(&buf, sisterId)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, "Directory uploaded successfully")
}

func sendTarballSisters(buf *bytes.Buffer, sisterId string) error {
	address, port, _ := getSistersInfo(sisterId)
	cli, err := client.Connect(address, port)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	stream, err := container.NewContainerServiceClient(cli).Upload(context.Background())
	if err != nil {
		panic(err)
	}

	chunkSize := 1024 * 1024
	for {
		chunk := make([]byte, chunkSize)
		n, err := buf.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read buffer: %v", err)
		}

		req := &container.UploadArchiveRequest{
			Archive: chunk[:n],
		}

		if err := stream.Send(req); err != nil {
			panic(err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	if res.Status != "Success" {
		return nil
	}

	return nil

}

func getSistersInfo(id string) (string, int, string) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	uid := uuid.MustParse(id)
	s, err := model.FindSister(ctx, db, uid)
	if err != nil {
		panic(err)
	}
	return s.Address, s.Port, s.Name
}

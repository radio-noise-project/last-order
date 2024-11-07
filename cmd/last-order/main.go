package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormRepo "github.com/radio-noise-project/last-order/internal/repository/gorm"
	"github.com/radio-noise-project/last-order/internal/rest"
	"github.com/radio-noise-project/last-order/internal/usecase/sisters"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		"db",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect database", slog.String("gorm", err.Error()))
		os.Exit(1)
	}

	e := echo.New()
	e.Use(middleware.CORS())

	// Prepare Repository
	sistersRepo := gormRepo.NewSistersRepository(db)
	// Build service Layer
	ss := sisters.NewService(sistersRepo)
	rest.NewSistersHandler(e, ss)

	e.Logger.Fatal(e.Start("8080"))
}

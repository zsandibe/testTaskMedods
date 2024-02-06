package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testTaskMedods/config"
	"testTaskMedods/internal/api/v3/delivery"
	"testTaskMedods/internal/repository"
	"testTaskMedods/internal/server"
	"testTaskMedods/internal/service"
	"testTaskMedods/internal/storage"
	"testTaskMedods/pkg"
)

func Start() {
	config := config.NewConfig()
	pkg.InfoLog.Println("Config loaded")

	storage, err := storage.NewMongoDb(config)
	if err != nil {
		pkg.ErrorLog.Println(err)
		return
	}
	pkg.InfoLog.Println("Storage loaded")

	repository := repository.NewRepository(storage)
	pkg.InfoLog.Println("Repository loaded")

	service := service.NewService(repository)
	pkg.InfoLog.Println("Service loaded")

	delivery := delivery.NewHandler(service)
	pkg.InfoLog.Println("Delivery loaded")

	server := server.NewServer(config, delivery.Routes())

	go func() {
		if err := server.Run(); err != nil {
			pkg.ErrorLog.Printf("failed to start server: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		pkg.ErrorLog.Printf("failed to shutdown server: %v")
		return
	}

}

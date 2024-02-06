package app

import (
	"testTaskMedods/config"
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

}

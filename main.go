package main

import (
	"delos-intern/config"
	"delos-intern/model"
	"delos-intern/route"
	"delos-intern/service"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		db   = config.InitDatabase()
		PORT = os.Getenv("PORT")
		Svc  = service.NewService(db)
	)
	model.DB.AutoMigrate(&model.Farm{}, &model.Pond{}, &model.Telemetry{})
	app := route.New(Svc)

	app.Start(":" + PORT)
}

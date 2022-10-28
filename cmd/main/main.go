package main

import (
	"log"
	"pahan/config"
	_ "pahan/docs"
	"pahan/internal/app"
)

// @title Автомобилестроение в экономике
// @version 1.0
// @description Курсовая работа по предмету "Информационные системы и базы данных" студента группы P34312, Соловьева Павла

// @host localhost:9000
// @BasePath /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}

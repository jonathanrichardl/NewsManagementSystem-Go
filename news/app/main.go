package main

import (
	"log"
	"news/pkg/database"
	"news/pkg/handlers"
	logger "news/pkg/logger"
	"news/pkg/router"
)

func main() {
	Logger := logger.NewLogger("log.txt")
	Logger.InfoLogger.Println("Reading database configuration")
	/*
		databaseConfig, err := config.LoadDatabaseConfiguration()
		if err != nil {
			log.Printf("Error setting database : %s\n", err.Error())
			return
		}
	*/
	//initializing db and router
	Logger.InfoLogger.Println("Initializing Program")
	/*
		Database, err := database.NewDatabase("mysql",
			databaseConfig.Username, databaseConfig.Password, databaseConfig.Address,
			databaseConfig.DatabaseName)
	*/
	Database, err := database.NewDatabase("mysql",
		"root", "123jonathan123100300!!!", "localhost:3306",
		"testers")
	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
		return
	}
	Router := router.NewRouterInstance()
	handlers := handlers.NewHttpHandlers(Database, Router, Logger)
	handlers.RegisterAllHandlers()
	Router.Start()
}

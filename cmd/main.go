package main

import (
	"log"

	"github.com/gin-gonic/gin"

	seeder "github.com/fathimasithara01/ecommerce/Seeder"
	"github.com/fathimasithara01/ecommerce/config"
	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/migration"
	"github.com/fathimasithara01/ecommerce/routes"
)

func main() {
	r := gin.Default()
	cfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the config file")
	}

	db, err := database.ConnectDB(cfig)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	migration.Migration()

	seeder.GroupSeeder()

	routes.RegisterRoutes(r, db)

	log.Printf("Starting server at %s\n", cfig.ServerAddr)
	if err := r.Run(cfig.ServerAddr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}

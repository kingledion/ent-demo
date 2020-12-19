package main

import (
	"context"
	"log"

	"github.com/kingledion/ent-demo/internal/config"
	"github.com/kingledion/ent-demo/internal/repository"
)

func main() {

	// Retrieve config settings
	config, err := config.GetDBConfig()
	if err != nil {
		log.Fatalf("cannot set up config: %v", err)
	}

	log.Printf("config settings %v", config)

	// Open DB
	repoconfig := repository.Config{
		User:   config.Username,
		Pass:   config.Password,
		Port:   config.Port,
		DBName: config.DBName,
	}
	repo, err := repository.New(repoconfig)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	defer repo.Close()

	// Run the auto migration tool.
	if err := repo.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Launch the server
}

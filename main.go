package main

import (
	"context"
	"log"

	"github.com/kingledion/ent-demo/internal/config"
	"github.com/kingledion/ent-demo/internal/httpservice"
	"github.com/kingledion/ent-demo/internal/repository"
	"github.com/kingledion/ent-demo/internal/service"
)

func main() {

	// Retrieve config settings
	dbConfig, err := config.GetDBConfig()
	if err != nil {
		log.Fatalf("cannot set up config: %v", err)
	}

	httpConfig, err := config.GetHttpConfig()
	if err != nil {
		log.Fatalf("cannot set up config: %v", err)
	}

	// Open DB
	repo, err := repository.New(dbConfig)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	defer repo.Close()

	// Run the auto migration tool.
	if err := repo.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Launch the service
	s := service.New(repo)

	// Launch the server
	h := httpservice.New(s)
	httpservice.Run(h, httpConfig)
}

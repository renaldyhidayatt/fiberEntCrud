package config

import (
	"log"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"

	"entgo.io/ent/dialect"
)

func Database() *ent.Client {
	client, err := ent.Open(dialect.Postgres, "postgresql://postgres:postgres@localhost:5432/fiberent?sslmode=disable")

	if err != nil {
		log.Fatalf("Failed opening connection postgresql: %v", err)
	}

	defer client.Close()

	return client

}

package config

import (
	"context"
	"fmt"
	"log"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/pkg"

	_ "github.com/lib/pq"
)

func Database() *ent.Client {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pkg.GodotEnv("DB_HOST"), pkg.GodotEnv("DB_PORT"), pkg.GodotEnv("DB_USER"), pkg.GodotEnv("DB_PASSWORD"), pkg.GodotEnv("DB_NAME"))

	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// ctx := context.Background()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf(err.Error())
	}

	// defer client.Close()

	return client

}

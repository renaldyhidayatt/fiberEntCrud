package config

import (
	"context"
	"fmt"
	"log"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/pkg"

	_ "github.com/lib/pq"
)

func Database(context context.Context) (*ent.Client, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pkg.GodotEnv("DB_HOST"), pkg.GodotEnv("DB_PORT"), pkg.GodotEnv("DB_USER"), pkg.GodotEnv("DB_PASSWORD"), pkg.GodotEnv("DB_NAME"))

	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Connected to database successfully")
	// defer client.Close()
	if err := client.Schema.Create(context); err != nil {
		log.Fatalf(err.Error())
	}

	return client, nil

}

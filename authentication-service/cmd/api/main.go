package main

import (
	"authentication/data"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const webPort = ":8081"

type Config struct {
	DB     *sqlx.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentiation service")

	conn, err := connectToDB()
	if err != nil {
		log.Panic(err)
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := app.routes()

	err = srv.Run(webPort)
	if err != nil {
		log.Panic(err)
	}
}

func connectToDB() (*sqlx.DB, error) {
	// Use it when running in docker
	//dsn := os.Getenv("DSN")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, err
	}
	return db, nil
}

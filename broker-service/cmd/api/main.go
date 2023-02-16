package main

import "log"

//If running in docker can use port 80, since docker allows multiple services to run on same port
const webPort = ":8080"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := app.routes()

	err := srv.Run(webPort)
	if err != nil {
		log.Panic()
	}
}

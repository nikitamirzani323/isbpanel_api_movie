package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/isbtotogroup/isbpanel_api_movie/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_movie/helpers"
	"bitbucket.org/isbtotogroup/isbpanel_api_movie/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	initRedis := helpers.RedisHealth()

	if !initRedis {
		panic("cannot load redis")
	}

	db.Init()

	app := routers.Init()

	if !initRedis {
		panic("cannot load redis")
	}
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "5051"
		}

		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}

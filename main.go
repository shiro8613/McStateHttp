package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/McStateHttp/config"
	"github.com/McStateHttp/pinger"
	"github.com/McStateHttp/server"
)

const configPath = "./config.yml"

func init() {
	if err := config.Load(configPath); err != nil {
		if os.IsNotExist(err) {
			config.CreateConfig(configPath)
			log.Fatalln("config file doesn't exists.")
		} else {
			log.Fatalln(err)
		}
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	go func() {
		pinger.Run(ctx)
	}()

	if err := server.Run(ctx); err != nil {
		log.Fatalln(err)
	}
}
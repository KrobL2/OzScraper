package main

import (
	"context"
	"flag"
	"log"

	tgClient "scrappy/clients/telegram"
	event_consumer "scrappy/consumer/event-consumer"
	"scrappy/events/telegram"
	"scrappy/internal/config"
	"scrappy/storage/sqlite"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("can't load config", err)
	}

	s, err := sqlite.New(cfg.Storage_path)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(cfg.Host, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, cfg.Batch_size)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

/* const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
) */

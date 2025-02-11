package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/nchhillar2004/gobot/handlers"
)

func initConfig() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    token := os.Getenv("TOKEN")
    if token == "" {
        log.Fatal("Error fetching TOKEN from .env")
    }
    return token
}

func main(){
    Token := initConfig()

    session, err := discordgo.New("Bot " + Token)
    if err != nil {
        log.Fatalf("Error creating Discord session: %v", err)
    }

    handlers.InitHandlers(session)

    err = session.Open()
    if err != nil {
        log.Fatalf("Error opening connection: %v", err)
    }
    defer session.Close()

    log.Println("Bot is now running. Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
    <-sc
}


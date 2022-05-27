package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"CuteyGiffer/common"
	"CuteyGiffer/handler"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const DISCORD_TOKEN = "DISCORD_TOKEN"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Replit doesn't need to read .env files.")
	}

	Token := os.Getenv(DISCORD_TOKEN)
	//refresh token/JWT
	if err != nil {
		fmt.Println("Replit doesn't need to read .env files.")
	}

	// 2
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating a discord Session ", err)
	}

	// Handler funcs
	dg.AddHandler(handler.Ready)
	dg.AddHandler(handler.MessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println(common.ErrorOpenningDiscord)
	}
	fmt.Println("The bot is now running . Press CTRL-C to exit")
	//5
	http.HandleFunc("/", handler123)
	http.ListenAndServe(":8000", nil)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Kill)
	<-sc
}

func handler123(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

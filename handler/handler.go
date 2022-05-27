package handler

import (
	"CuteyGiffer/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

//Ready is handler for getting ready
func Ready(s *discordgo.Session, event *discordgo.Event) {
	s.UpdateGameStatus(0, "!search < keyword >")
}

//MessageCreate is handler for creating response message to !search command
func MessageCreate(s *discordgo.Session, message *discordgo.MessageCreate) {
	// 1
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Replit doesn't need to read .env files.")
	}
	giphyToken := os.Getenv("GIPHY_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	// 2
	if message.Author.ID == s.State.User.ID {
		return
	}
	// 3
	command := strings.Split(message.Content, " ")

	if command[0] == "!help" {
		s.ChannelMessageSend(message.ChannelID, "Type !(exclamation mark) search and add your favorite word to call a cute giff♥")

	}

	if command[0] == "!search" && len(command) > 1 {
		url := "https://api.giphy.com/v1/gifs/random"
		var result models.GifSearch
		// 5
		gifKeyword := strings.Join(command[1:], " ")

		// 6
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error in making a new Request", err)
		}
		query := req.URL.Query()
		query.Add("api_key", giphyToken)
		query.Add("tag", gifKeyword)
		req.URL.RawQuery = query.Encode()
		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Error in getting a response, ", err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshall JSON", err)
		}
		// 7
		s.ChannelMessageSend(message.ChannelID, result.Data.EmbedURL)
		res.Body.Close()
	}
}

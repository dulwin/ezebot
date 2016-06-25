package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"strings"
)

func main() {
	key := os.Getenv("SLACK_KEY")
	if key == nil {
		fmt.Println("Could not find slack api key, make sure to set environment variables")
		return
	}
	api := slack.New(key)
	logger := log.New(os.Stdout, "Jarvis: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)
	rtm := api.NewRTM()

	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				go messageHandler(ev, rtm)
			}

		}
	}

}

func messageHandler(event *slack.MessageEvent, rtm *slack.RTM) {
	if strings.Contains(event.Msg.Text, "<@"+rtm.GetInfo().User.ID+">") {
		rtm.SendMessage(rtm.NewOutgoingMessage(event.Text, event.Channel))
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nlopes/slack"

	"github.com/dulwin/ezebot/db"
	"github.com/dulwin/ezebot/models"
	"github.com/dulwin/ezebot/nlp"
	"github.com/dulwin/ezebot/utils"
)


var logger *log.Logger
var rtm *slack.RTM
var api *slack.Client

func init() {
	logger = log.New(os.Stdout, "Jarvis: ", log.Lshortfile|log.LstdFlags)
	api = initializeApi()
	rtm = api.NewRTM()
}

func initiateGetRequest() {
	var (
		resp *http.Response
		err  error
	)
	for {
		resp, err = http.Get(os.Getenv("API_ENDPOINT"))
		utils.CheckError(err)
		resp.Body.Close()
		time.Sleep(20 * time.Minute)
	}
}

func spawnServer() {
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/", webHandler)
	http.ListenAndServe(port, nil)
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func initializeApi() *slack.Client {
	var key string
	key = os.Getenv("SLACK_KEY")
	logger.Print("KEY: " + key)
	api := slack.New(key)
	slack.SetLogger(logger)
	logger.Print("Initiated slack client")
	return api
}

func messageHandler(event *slack.MessageEvent, rtm *slack.RTM) {
	if strings.Contains(event.Msg.Text, "<@"+rtm.GetInfo().User.ID+">") {
		witResponse := nlp.ProcessMessage(event.Msg.Text)
		s := fmt.Sprintf("%+v \n", witResponse)
		rtm.SendMessage(rtm.NewOutgoingMessage(s, event.Channel))
	}
}

func main() {
	entityManager := db.GetInstance()
	defer entityManager.Close()
	entityManager.Migrate()
	
	// q := models.Query{Category: "test_category", Query: "doorcode", Response: "HAHAHAA"}
	// entityManager.Insert(&q)

	go spawnServer()
	go initiateGetRequest()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			logger.Printf("%+v\n", msg)
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				go messageHandler(ev, rtm)
			}
		}
	}
}

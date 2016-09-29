package main

import(
  "github.com/nlopes/slack"
  "log"
  "os"
  "strings"
  "net/http"
  "fmt"
  "time"
)

func initiate_get_request() {
  var resp *http.Response
  var err error
  for {
    resp, err = http.Get(os.Getenv("API_ENDPOINT"))
    if err != nil {
      panic(err)
    }
    defer resp.Body.Close()
    time.Sleep(20 * time.Minute)
  }
}

func spawnServer(){
  port := ":" + os.Getenv("PORT")
  http.HandleFunc("/", web_handler)
  http.ListenAndServe(port, nil)
}

func web_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main(){
  logger := log.New(os.Stdout, "Jarvis: ", log.Lshortfile|log.LstdFlags)
  api := initialize(logger)
  rtm := api.NewRTM()
  go spawnServer()
  go initiate_get_request()
  go rtm.ManageConnection()
  for {
    select {
    case msg:= <-rtm.IncomingEvents:
      switch ev := msg.Data.(type) {
        case *slack.MessageEvent:
          go messageHandler(ev, rtm)
      }
    }
  }
}

func initialize(logger *log.Logger) *slack.Client  {
  var key string
  key = os.Getenv("SLACK_KEY")
  logger.Print("KEY: " + key)
  api := slack.New(key)
  slack.SetLogger(logger)
  logger.Print("Initiated slack client")
  return api
}

func check(e error) {
  if e != nil{
    panic(e)
  }
}

func messageHandler(event *slack.MessageEvent, rtm *slack.RTM){
  if strings.Contains(event.Msg.Text, "<@" + rtm.GetInfo().User.ID + ">") {
    rtm.SendMessage(rtm.NewOutgoingMessage(event.Text, event.Channel))
  }
}

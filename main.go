package main

import(
  "github.com/nlopes/slack"
  "log"
  "os"
  "strings"
  "net/http"
  "fmt"
  "time"
  "io/ioutil"
)

var production bool

func initiate_get_request() {
  var resp *http.Response
  var err error
  if (production){
    resp, err = http.Get(os.Getenv("API_ENDPOINT"))
  } else {
    resp, err = http.Get("http://uw-ezebot.herokuapp.com")
  }
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  time.Sleep(20 * time.Second)
}

func spawn_server(){
  port := ":" + os.Getenv("PORT")
  http.HandleFunc("/", web_handler)
  http.ListenAndServe(port, nil)
}

func web_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main(){
  production = os.Getenv("ENVIRONMENT") == "production"

  logger := log.New(os.Stdout, "Jarvis: ", log.Lshortfile|log.LstdFlags)
  api := initialize(logger)
  rtm := api.NewRTM()
  go spawn_server()
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
  if (production){
    key = os.Getenv("SLACK_KEY")
  } else {
    key = getKey(logger)
  }
  api := slack.New(key)
  slack.SetLogger(logger)
  logger.Print("Initiated slack client")
  return api
}

func getKey(logger *log.Logger) string {
  data, err := ioutil.ReadFile("keys/api_key.txt")
  check(err)
  logger.Print("Got slack key")
  key := string(data)
  length := len(key)
  return key[:length - 1]
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

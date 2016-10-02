package nlp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/dulwin/ezebot/utils"
)

type WitResponse struct {
	MessageId string              `json:"msg_id"`
	Text      string              `json:"_text"`
	Entities  map[string][]Entity `json:"entities"`
}

type Entity struct {
	Confidence float32 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type,omitempty"`
	Grain      string  `json:"grain,omitempty"`
}

const URL string = "https://api.wit.ai/message?"

func ProcessMessage(message string) *WitResponse {

	client := &http.Client{}
	queryString := url.Values{}
	queryString.Set("q", message)

	requestUrl := URL + queryString.Encode()
	req, err := http.NewRequest("GET", requestUrl, nil)
	utils.CheckError(err)

	req.Header.Add("Authorization", "Bearer "+os.Getenv("WIT_API_KEY"))
	resp, err := client.Do(req)
	utils.CheckError(err)

	buff, err := ioutil.ReadAll(resp.Body)
	utils.CheckError(err)
	witResponse := WitResponse{}

	json.Unmarshal(buff, &witResponse)
	log.Printf("%+v", witResponse)
	return &witResponse

}

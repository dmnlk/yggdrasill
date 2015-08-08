package gomadare

import (
	"bufio"
	"encoding/json"

	"log"
)

// userstream url
const (
	STREAM_URL = "https://userstream.twitter.com/1.1/user.json"
)

// get User Stream
func (client *Client) GetUserStream(params map[string]string, f func(Status, Event)) {
	response, err := client.consumer.Get(STREAM_URL, params, client.accessToken)
	if err != nil {
		return
	}
	defer response.Body.Close()
	scanner := bufio.NewScanner(response.Body)
	// ignore friend list id
	scanner.Scan()
	for {
		if ok := scanner.Scan(); !ok {
			log.Fatal("scan error")
			continue
		}
		var result interface{}
		b := scanner.Bytes()
		//json化
		if err := json.Unmarshal(b, &result); err != nil {
			log.Println(err)
			continue
		}
		var event Event
		var status Status
		msg := result.(map[string]interface{})
		if _, ok := msg["event"]; ok {
			if err := json.Unmarshal(b, &event); err != nil {
				continue
			}
		}
		if _, ok := msg["user"]; ok {
			if err := json.Unmarshal(b, &status); err != nil {
				continue
			}
		}
		f(status, event)
	}
}

package main

import (
	"fmt"
	"os"

	"log"

	"github.com/dmnlk/gomadare"
	"github.com/dmnlk/stringUtils"
	"github.com/rem7/goprowl"
	"strings"
)

var (
	CONSUMER_KEY        string
	CONSUMER_KEY_SECRET string
	ACCESS_TOKEN        string
	ACCESS_TOKEN_SECRET string
	PROWL_API_KEY       string
	PROWL               goprowl.Goprowl
)

func main() {
	err := configureToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = PROWL.RegisterKey(PROWL_API_KEY)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := gomadare.NewClient(CONSUMER_KEY, CONSUMER_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
	fmt.Println("aa")
	client.GetUserStream(nil, func(s gomadare.Status, e gomadare.Event) {
		if &s != nil {
			go sendReplyAndRetweetToProwl(s)
		}
		if &e != nil {
			go sendEventToProwl(e)
		}
	})
}

func configureToken() error {
	CONSUMER_KEY = os.Getenv("CONSUMER_KEY")
	CONSUMER_KEY_SECRET = os.Getenv("CONSUMER_KEY_SECRET")
	ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	PROWL_API_KEY = os.Getenv("PROWL_API_KEY")
	if ng := stringUtils.IsAnyEmpty(CONSUMER_KEY, CONSUMER_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET, PROWL_API_KEY); ng {
		return fmt.Errorf("some key invalid")
	}

	return nil
}

func sendEventToProwl(e gomadare.Event) {
	if stringUtils.IsEmpty(e.Event) {
		return
	}
	emoji := getEventEmoji(e)
	n := &goprowl.Notification{
		Application: "yggdrasill",
		Description: emoji + " " + e.TargetObject.Text,
		Event:       e.Event + " by " + e.Source.ScreenName,
		Priority:    "1",
	}

	PROWL.Push(n)
}

func getEventEmoji(event gomadare.Event) string {
	if event.Event == "favorite" {
		return "\u2b50"
	}
	if event.Event == "unfavorite" {
		return "\U0001f44e"
	}
	if event.Event == "list_member_removed" {
		return "\u274c"
	}
	if event.Event == "list_member_added" {
		return "\u2755"
	}
	if event.Event == "follow" {
		return "\u2661"
	}
	log.Println("unknown event:" + event.Event)
	return event.Event
}


func getProwlNotification(event gomadare.Event, string emoji) goprowl.Notification {
	n := &goprowl.Notification{
		Application: "Yggdrsill",
	}
	return *n
}

func sendReplyAndRetweetToProwl(s gomadare.Status) {
	// reply Event
	if len(s.Entities.UserMentions) > 0 {
		for _, mention := range s.Entities.UserMentions {
			if mention.ScreenName == "dmnlk" {
				var n *goprowl.Notification
				if strings.Contains(s.Text, "RT") {
					n = &goprowl.Notification{
						Application: "Golang",
						Description: "\U0001f4a1" + " " + s.Text,
						Event:       "RT by " + s.User.ScreenName,
						Priority:    "1",
					}
				} else {
					n = &goprowl.Notification{
						Application: "Golang",
						Description: "\U0001f4a1" + " " + s.Text,
						Event:       "Mentioned by " + s.User.ScreenName,
						Priority:    "1",
					}
				}

				PROWL.Push(n)
			}
		}
	}
}

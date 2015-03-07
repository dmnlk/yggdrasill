package main

import (
	"fmt"
	"os"

	"github.com/dmnlk/gomadare"
	"github.com/dmnlk/stringUtils"
	"github.com/rem7/goprowl"
	"github.com/k0kubun/pp"
)

var (
	CONSUMER_KEY        string
	CONSUMER_KEY_SECRET string
	ACCESS_TOKEN        string
	ACCESS_TOKEN_SECRET string
	PROWL_API_KEY       string
	PROWL goprowl.Goprowl
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
			sendReplyAndRetweetToProwl(s)
		}
		if &e != nil {
			sendEventToProwl(e)
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
		return;
	}
	emoji := getEmoji(e)
	n := &goprowl.Notification{
		Application: "Twitter",
		Description:  emoji + " " + e.TargetObject.Text,
		Event:       e.Event + " by " + e.Source.Name,
		Priority:    "1",
	}

	PROWL.Push(n)
}

func getEmoji(event gomadare.Event) string {
	if event.Event == "favorite" {
		return "\u2b50"
	}
	if event.Event == "unfavorite" {
		return "\U0001f44e"
	}
	return ""
}

func sendReplyAndRetweetToProwl(s gomadare.Status) {
	pp.Print(s.Entities.UserMentions)
}

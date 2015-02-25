package main

import (
	"fmt"
	"os"

	"github.com/dmnlk/gomadare"
	"github.com/dmnlk/stringUtils"
	"github.com/rem7/goprowl"
)

var (
	CONSUMER_KEY        string
	CONSUMER_KEY_SECRET string
	ACCESS_TOKEN        string
	ACCESS_TOKEN_SECRET string
	PROWL_API_KEY       string
)

func main() {
	error := configureToken()
	if error != nil {
		fmt.Println(error)
		return
	}

	client := gomadare.NewClient(CONSUMER_KEY, CONSUMER_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	client.GetUserStream(nil, func(s gomadare.Status, e gomadare.Event) {
		if &s != nil {

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
	if ng := stringUtils.IsAnyEmpty(CONSUMER_KEY, CONSUMER_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET); ng {
		return fmt.Errorf("some key invalid")
	}

	return nil
}

func sendEventToProwl(e gomadare.Event) {

	var p goprowl.Goprowl
	err := p.RegisterKey(PROWL_API_KEY)
	if err != nil {
		fmt.Println(err)
		return
	}

	n := &goprowl.Notification{
		Application: "Twitter",
		Description:  e.Event + e.TargetObject.Text,
		Event:       e.Event + e.Source.Name,
		Priority:    "1",
		Providerkey: "",
		Url:         "www.foobar.com",
	}

	p.Push(n)

}

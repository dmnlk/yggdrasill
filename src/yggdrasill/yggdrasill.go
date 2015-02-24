package main

import (
	"fmt"
	"os"
	"github.com/dmnlk/gomadare"
	"github.com/k0kubun/pp"
	"github.com/dmnlk/stringUtils"
)



var (
	CONSUMER_KEY string
	CONSUMER_KEY_SECRET string
	ACCESS_TOKEN string
	ACCESS_TOKEN_SECRET string
)



func main() {
	error := configureToken()
	if error != nil {
		fmt.Println(error)
		return
	}

	client := gomadare.NewClient(CONSUMER_KEY, CONSUMER_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	client.GetUserStream(nil,  func(s gomadare.Status, e gomadare.Event) {
		if &s != nil {
			fmt.Println("return status")
			pp.Print(s)
		}
		if &e != nil {
			fmt.Println("return event")
			pp.Print(e)
		}
	})
}

func  configureToken()(error) {
	CONSUMER_KEY = os.Getenv("CONSUMER_KEY")
	CONSUMER_KEY_SECRET = os.Getenv("CONSUMER_KEY_SECRET")
	ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	if ng := stringUtils.IsAnyEmpty(CONSUMER_KEY, CONSUMER_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET); ng {
		return  fmt.Errorf("some key invalid")
	}


	return nil
}


package main

import (
	"fmt"
	"os"
	"github.com/mrjones/oauth"
)

type  api struct {
	consumer *oauth.Consumer
	token *oauth.AccessToken
}

var (
	CONSUMER_KEY string
	CONSUMER_KEY_SECRET string
	ACCESS_TOKEN string
	ACCESS_TOKEN_SECRET string
)

var provider  = oauth.ServiceProvider{
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}

func main() {

	error := configureToken()

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(CONSUMER_KEY)
	fmt.Println(CONSUMER_KEY_SECRET)
	fmt.Println(ACCESS_TOKEN)
	fmt.Println(ACCESS_TOKEN_SECRET)
}

func  configureToken()(error) {
	CONSUMER_KEY = os.Getenv("CONSUMER_KEY")
	CONSUMER_KEY_SECRET = os.Getenv("CONSUMER_KEY_SECRET")
	ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")

	//if key is not complete, throw error
	if len(CONSUMER_KEY) == 0 {
		return  fmt.Errorf("CONSUMER_KEY is blank")
	}
	if len(CONSUMER_KEY_SECRET) == 0 {
		return  fmt.Errorf("CONSUMER_KEY_SECRET is blank")
	}
	if len(ACCESS_TOKEN) == 0 {
		return  fmt.Errorf("ACCESS_TOKEN is  blank")
	}
	if len(ACCESS_TOKEN_SECRET) == 0 {
		return  fmt.Errorf("ACCESS_TOKEN_SECRET is blank")
	}

	return nil
}

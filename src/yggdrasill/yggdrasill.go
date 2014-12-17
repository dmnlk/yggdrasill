package main

import (
	"fmt"
	"os"
)

var (
	CONSUMER_KEY string
	CONSUMER_KEY_SECRET string
	ACCESS_TOKEN string
	ACCESS_TOKEN_SECRET string
)
func main() {

	_ = configureToken()

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

	if CONSUMER_KEY != "" || CONSUMER_KEY_SECRET == "" || ACCESS_TOKEN == "" || ACCESS_TOKEN_SECRET == "" {
		return  fmt.Errorf("error %d", 1)
	}
	return nil
}

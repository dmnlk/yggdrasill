package main

import (
	"os"

	"fmt"

	"github.com/dmnlk/gomadare"
	"github.com/k0kubun/pp"
)

func main() {
	ck := os.Getenv("ck")
	cs := os.Getenv("cs")
	at := os.Getenv("at")
	as := os.Getenv("as")
	client := gomadare.NewClient(ck, cs, at, as)
	client.GetUserStream(nil, func(s gomadare.Status, e gomadare.Event) {
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

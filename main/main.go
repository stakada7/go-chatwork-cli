package main

import (
	"flag"
	"fmt"
	cw "github.com/griffin-stewie/go-chatwork"
	"strconv"
)

var apiToken string

func init() {
	flag.StringVar(&apiToken, "token", "", "Chatwork API key")
	flag.Parse()
}

func main() {

	chatwork := cw.NewClient(apiToken)
	fmt.Printf("Your rate limit: %+v\n", chatwork.RateLimit())

	s, err := chatwork.MyStatus()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("Your status: %+v\n", s)

	r, err := chatwork.Rooms()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	for i := 0; i < len(r); i++ {
		if r[i].UnreadNum > 0 {
			fmt.Printf("room name: %+v\n", r[i].Name)
			fmt.Printf("unread unmber: %+v\n", r[i].UnreadNum)

			m, err := chatwork.RoomMessages(strconv.Itoa(r[i].RoomID), map[string]string{"force": "1"})
			if err != nil {
				fmt.Printf("%+v\n", err)
			}

			for j := len(m) - 1; j >= len(m)-r[i].UnreadNum; j-- {
				fmt.Printf("messages: %+v\n", m[j])
			}

		}
	}

}

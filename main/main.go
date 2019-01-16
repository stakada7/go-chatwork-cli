package main

import (
	"flag"
	"fmt"
	cw "github.com/griffin-stewie/go-chatwork"
	"strconv"
	"strings"
	"time"
)

var apiToken string
var chatwork *cw.Client
var accountId int

func init() {
	flag.StringVar(&apiToken, "token", "", "Chatwork API key")
	flag.Parse()

	chatwork = cw.NewClient(apiToken)
	fmt.Printf("Your rate limit: %d\n", chatwork.RateLimit().Remaining)

	me, err := chatwork.Me()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	accountId = me.AccountID

}

func main() {

	s, err := chatwork.MyStatus()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("Unread: %d\n", s.UnreadNum)
	fmt.Printf("Unread room: %d\n", s.UnreadRoomNum)
	fmt.Printf("Mention: %d\n", s.MentionNum)
	fmt.Printf("Mention room: %d\n", s.MentionRoomNum)
	//fmt.Printf("%+v\n", s)

	r, err := chatwork.Rooms()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	for i := 0; i < len(r); i++ {
		if r[i].UnreadNum > 0 {

			fmt.Println()
			fmt.Println("===")
			fmt.Printf("room name: %+v\n", r[i].Name)
			fmt.Printf("unread unmber: %+v\n", r[i].UnreadNum)

			m, err := chatwork.RoomMessages(strconv.Itoa(r[i].RoomID), map[string]string{"force": "1"})
			if err != nil {
				fmt.Printf("%+v\n", err)
			}

			for j := len(m) - r[i].UnreadNum; j < len(m); j++ {

				fmt.Println("---")

				fmt.Printf("send time: %s\n", time.Unix(m[j].SendTime, 0))
				fmt.Printf("update time: %s\n", time.Unix(m[j].UpdateTime, 0))

				if strings.LastIndex(m[j].Body, strconv.Itoa(accountId)) > 0 {
					fmt.Printf("\x1b[31mfrom: %s\x1b[0m\n", m[j].Account.Name)
					fmt.Printf("\x1b[31mmessage: %s\x1b[0m\n", m[j].Body)
				} else {
					fmt.Printf("from: %s\n", m[j].Account.Name)
					fmt.Printf("message: %s\n", m[j].Body)
				}

			}

		}
	}

}

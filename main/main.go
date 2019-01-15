package main

import (
	"fmt"
	cw "github.com/griffin-stewie/go-chatwork"
	"strconv"
)

func main() {

	cw := cw.NewClient(`4d74d8bbb2c42c166156ccce2e6ab44e`)
	fmt.Printf("Your rate limit: %+v\n", cw.RateLimit())

	s, err := cw.MyStatus()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("Your status: %+v\n", s)

	r, err := cw.Rooms()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	for i := 0; i < len(r); i++ {
		if r[i].UnreadNum > 0 {
			fmt.Printf("room name: %+v\n", r[i].Name)
			fmt.Printf("unread unmber: %+v\n", r[i].UnreadNum)

			m, err := cw.RoomMessages(strconv.Itoa(r[i].RoomID), map[string]string{"force": "1"})
			if err != nil {
				fmt.Printf("%+v\n", err)
			}

			for j := len(m) - 1; j >= len(m)-r[i].UnreadNum; j-- {
				fmt.Printf("messages: %+v\n", m[j])
			}

		}
	}

}

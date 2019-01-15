package main

import (
	cw "github.com/griffin-stewie/go-chatwork"
	"log"
)

func main() {

	cw := cw.NewClient(`4d74d8bbb2c42c166156ccce2e6ab44e`)
	log.Printf("%+v", cw.RateLimit())

	m, err := cw.Me()
	if err != nil {
		log.Printf("%+v", err)
	}
	log.Printf("%+v", m)

	s, err := cw.MyStatus()
	if err != nil {
		log.Printf("%+v", err)
	}
	log.Printf("%+v", s)

	t, err := cw.MyTasks(map[string]string{})
	if err != nil {
		log.Printf("%+v", err)
	}
	log.Printf("%+v", t)

	//c, err := cw.Contacts()
	//if err != nil {
	//	log.Printf("%+v", err)
	//}
	//log.Printf("%+v", c)
	//
	//r, err := cw.Rooms()
	//if err != nil {
	//	log.Printf("%+v", err)
	//}
	//log.Printf("%+v", r)
}

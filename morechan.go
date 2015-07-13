package main

import (
	"encoding/json"
	"github.com/grimborg-trash/morechan/fakehttp"
	"log"
)

type ResponseMeta struct {
	HasMore       bool `json:"has_more"`
	LastTimestamp int  `json:"last_timestamp"`
}

type Cat struct {
	Name string `json:"name"`
}

type IceCream struct {
	Flavor string `json:"flavor"`
}

type Cats struct {
	Items []Cat `json:"items"`
	ResponseMeta
}

type IceCreams struct {
	Items []IceCream `json:"items"`
	ResponseMeta
}

func GetAll(resource string, c *chan []byte) {
	lastTimestamp := 0
	for {
		b := fakehttp.Get(resource, lastTimestamp)
		*c <- b
		resp := &ResponseMeta{}
		err := json.Unmarshal(b, &resp)
		if err != nil {
			log.Fatalln(err)
		}
		if !resp.HasMore {
			close(*c)
			return
		}
		lastTimestamp = resp.LastTimestamp
	}
}

func GetAllCats(resp *chan Cat) {
	c := make(chan []byte)
	go GetAll("cat", &c)
	for data := range c {
		cats := &Cats{}
		err := json.Unmarshal(data, &cats)
		if err != nil {
			log.Fatalln(err)
		}
		for _, item := range cats.Items {
			*resp <- item
		}
	}
	close(*resp)
}

func GetAllIceCreams(resp *chan IceCream) {
	c := make(chan []byte)
	go GetAll("ice_cream", &c)
	for data := range c {
		iceCreams := &IceCreams{}
		err := json.Unmarshal(data, &iceCreams)
		if err != nil {
			log.Fatalln(err)
		}
		for _, item := range iceCreams.Items {
			*resp <- item
		}
	}
	close(*resp)
}

func main() {
	c := make(chan Cat)
	go GetAllCats(&c)
	for x := range c {
		log.Printf("%+v\n", x)
	}

	c2 := make(chan IceCream)
	go GetAllIceCreams(&c2)
	for x := range c2 {
		log.Printf("%+v\n", x)
	}
}

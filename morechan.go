package main

import (
	"encoding/json"
	"github.com/grimborg-trash/morechan/fakehttp"
	"log"
)

type Item map[string]interface{}

type Response struct {
	HasMore       bool   `json:"has_more"`
	LastTimestamp int    `json:"last_timestamp"`
	Items         []Item `json:"items,inline"`
}

type Cat struct {
	Name string `json:"name"`
}

type IceCream struct {
	Flavor string `json:"flavor"`
}

func GetAll(resource string, c chan Item) {
	lastTimestamp := 0
	for {
		b := fakehttp.Get(resource, lastTimestamp)
		resp := Response{}
		err := json.Unmarshal(b, &resp)
		if err != nil {
			log.Fatalln(err)
		}
		if !resp.HasMore {
			close(c)
			return
		}
		for _, i := range resp.Items {
			c <- i
		}
		lastTimestamp = resp.LastTimestamp
	}
}

func (item Item) Cat() Cat {
	return Cat{Name: item["name"].(string)}
}

func (item Item) IceCream() IceCream {
	return IceCream{Flavor: item["flavor"].(string)}
}

func main() {
	c := make(chan Item)
	go GetAll("cat", c)
	for x := range c {
		cat := x.Cat()
		log.Printf("%+v", cat)
	}

	c2 := make(chan Item)
	go GetAll("ice_cream", c2)
	for x := range c2 {
		cat := x.IceCream()
		log.Printf("%+v", cat)
	}
}

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

func GetCats(c chan Cat) {
	items := make(chan Item)
	go GetAll("cat", items)
	for item := range items {
		c <- item.Cat()
	}
	close(c)
}

func GetIceCreams(c chan IceCream) {
	items := make(chan Item)
	go GetAll("ice_cream", items)
	for item := range items {
		c <- item.IceCream()
	}
	close(c)
}

func main() {
	c := make(chan Cat)
	go GetCats(c)
	for cat := range c {
		log.Printf("%+v", cat)
	}

	c2 := make(chan IceCream)
	go GetIceCreams(c2)
	for iceCream := range c2 {
		log.Printf("%+v", iceCream)
	}

}

package main

import (
	"encoding/json"
	"github.com/grimborg/morechan/fakehttp"
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

// func GetAllCats() 59 Cat {
//     c := make(chan []Cat)
//     for {

//     }
// }

func main() {
	c := make(chan []byte)
	go GetAll("cat", &c)
	for x := range c {
		log.Println(string(x))
	}
}

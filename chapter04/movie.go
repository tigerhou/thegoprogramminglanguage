package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color,omitempty"`
	Actor []string
}

func main() {
	movies := []Movie{
		{
			Title: "战狼2",
			Year:  2017,
			Color: false,
			Actor: []string{"吴京", "张瀚"},
		},
		{
			Title: "妖猫传",
			Year:  2018,
			Color: true,
			Actor: []string{"黄轩", "张雨绮"},
		},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshal failed:%s", err)
	}
	fmt.Printf("%s\n", data)
	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshal failed:%s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshal failed:%v", err)
	}
	fmt.Printf("%v\n", titles)

}

package main

import "time"

func main() {
	mirroredQuery()

}

func mirroredQuery() string {
	responses := make(chan string)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}

func request(text string) string {
	var length = len(text)
	time.Sleep(time.Duration(length) * time.Second)
	print(text + "------done")
	return text
}

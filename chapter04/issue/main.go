package main

import (
	"fmt"
	"github.com/tigerhou/thegoprogramminglanguage/chapter04/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55\n", item.Number, item.User.Login, item.Title)
	}
}

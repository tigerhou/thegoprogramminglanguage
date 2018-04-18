package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"sock": 20, "skirt": 100}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:%s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "sock": 20}
	servemux := http.NewServeMux()
	//servemux.Handle("/list", http.HandlerFunc(db.list))
	//servemux.Handle("/price", http.HandlerFunc(db.price))
	servemux.HandleFunc("/list", db.list)
	servemux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8080", servemux))

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

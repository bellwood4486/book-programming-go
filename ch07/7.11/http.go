package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var mu sync.Mutex

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
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

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	q := req.URL.Query()
	item := q.Get("item")

	currentPrice, ok := (*db)[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	p := q.Get("price")
	newPrice, err := strconv.ParseFloat(p, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", p)
		return
	}

	(*db)[item] = dollars(newPrice)

	fmt.Fprintf(w, "updated price: %s -> %s\n", currentPrice, dollars(newPrice))
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// 参考：https://github.com/torbiak/gopl/blob/master/ex7.12/main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Dollars float32

func (d Dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func ParseDollars(s string) (Dollars, error) {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	return Dollars(f), nil
}

type PriceDB struct {
	sync.Mutex
	db map[string]Dollars
}

func NewPriceDB() *PriceDB {
	return &PriceDB{db: make(map[string]Dollars)}
}

type Stock struct {
	Item         string
	Price        Dollars
	DisplayPrice string
}

func (pdb *PriceDB) Create(w http.ResponseWriter, req *http.Request) {
	var stockReq Stock
	if err := json.NewDecoder(req.Body).Decode(&stockReq); err != nil {
		http.Error(w, fmt.Sprintf("json parse failed: %v\n", err), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	pdb.Lock()
	defer pdb.Unlock()
	if _, ok := pdb.db[stockReq.Item]; ok {
		http.Error(w, fmt.Sprintf("already created: item = %q\n", stockReq.Item), http.StatusConflict)
		return
	}
	pdb.db[stockReq.Item] = stockReq.Price

	w.WriteHeader(http.StatusCreated)
}

func (pdb *PriceDB) Read(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}
	price, ok := pdb.db[item]
	if !ok {
		http.Error(w, fmt.Sprintf("item not found: item = %q", item), http.StatusNotFound)
		return
	}

	jsonStock, err := json.Marshal(Stock{item, price, price.String()})
	if err != nil {
		http.Error(w, fmt.Sprintf("json marshal failure: %v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonStock))
}

func (pdb *PriceDB) Update(w http.ResponseWriter, req *http.Request) {
	var stockReq Stock
	if err := json.NewDecoder(req.Body).Decode(&stockReq); err != nil {
		http.Error(w, fmt.Sprintf("json parse failed: %v\n", err), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	pdb.Lock()
	defer pdb.Unlock()
	if _, ok := pdb.db[stockReq.Item]; !ok {
		http.Error(w, fmt.Sprintf("item not found: item = %q\n", stockReq.Item), http.StatusNotFound)
		return
	}
	pdb.db[stockReq.Item] = stockReq.Price

	w.WriteHeader(http.StatusNoContent)
}

func (pdb *PriceDB) Delete(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	pdb.Lock()
	defer pdb.Unlock()
	delete(pdb.db, item)

	w.WriteHeader(http.StatusNoContent)
}

func (pdb *PriceDB) List(w http.ResponseWriter, req *http.Request) {
	var stocks []Stock
	for item, price := range pdb.db {
		stocks = append(stocks, Stock{item, price, price.String()})
	}
	jsonStocks, err := json.Marshal(stocks)
	if err != nil {
		http.Error(w, fmt.Sprintf("json marshal failure: %v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonStocks))
}

func main() {
	pdb := &PriceDB{
		db: map[string]Dollars{
			"shoes": 50.1,
			"socks": 5,
		},
	}
	http.HandleFunc("/create", pdb.Create)
	http.HandleFunc("/read", pdb.Read)
	http.HandleFunc("/update", pdb.Update)
	http.HandleFunc("/delete", pdb.Delete)
	http.HandleFunc("/list", pdb.List)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

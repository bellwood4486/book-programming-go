// 参考：https://github.com/torbiak/gopl/blob/master/ex7.12/main.go
package main

import (
	"fmt"
	"html/template"
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

func (pdb *PriceDB) Create(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}
	p := req.FormValue("price")
	if p == "" {
		http.Error(w, "no price given", http.StatusBadRequest)
		return
	}
	price, err := ParseDollars(p)
	if err != nil {
		http.Error(w, fmt.Sprintf("parse failed: price = %q\n", p), http.StatusBadRequest)
		return
	}
	pdb.Lock()
	defer pdb.Unlock()
	if _, ok := pdb.db[item]; ok {
		http.Error(w, fmt.Sprintf("already created: item = %q\n", item), http.StatusConflict)
		return
	}
	pdb.db[item] = price

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
	fmt.Fprintf(w, "%s\n", price)
}

func (pdb *PriceDB) Update(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}
	p := req.FormValue("price")
	if p == "" {
		http.Error(w, "no price given", http.StatusBadRequest)
		return
	}
	price, err := ParseDollars(p)
	if err != nil {
		http.Error(w, fmt.Sprintf("parse failed: price = %q\n", p), http.StatusBadRequest)
		return
	}

	pdb.Lock()
	defer pdb.Unlock()
	if _, ok := pdb.db[item]; !ok {
		http.Error(w, fmt.Sprintf("item not found: item = %q\n", item), http.StatusNotFound)
		return
	}
	pdb.db[item] = price

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

var listHTML = template.Must(template.New("listHTML").Parse(`
<!DOCTYPE html>
<html>
<body>
    <table>
        <tr>
            <th>item</th>
            <th>price</th>
        </tr>
        {{range $k, $v := .}}
        <tr>
            <td>{{$k}}</td>
            <td>{{$v}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>
`))

func (pdb *PriceDB) List(w http.ResponseWriter, req *http.Request) {
	if err := listHTML.Execute(w, pdb.db); err != nil {
		http.Error(w, fmt.Sprintf("template failure: %v\n", err), http.StatusInternalServerError)
		return
	}
}

func main() {
	pdb := &PriceDB{
		db: map[string]Dollars{
			"shoes": 50,
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

package main

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
)

/*
 * Request Examples
 * localhost:8000/list
 * localhost:8000/price?item=socks
 * localhost:8000/create?item=hoge&price=5
 * localhost:8000/read?item=socks
 * localhost:8000/update?item=socks&price=10
 * localhost:8000/delete?item=socks
 */
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	// check: Is Item exist?
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s is exsit.\n", item)
		return
	}

	priceStr := req.URL.Query().Get("price")
	priceInt, ok := convertPositiveInt(w, req, priceStr)
	if !ok {
		return
	}
	db[item] = dollars(priceInt)
	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %s", item, db[item])
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s is not found.\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	// check: Is item exist?
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	intPrice, ok := convertPositiveInt(w, req, priceStr)
	if !ok {
		return
	}
	db[item] = dollars(intPrice)
	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "%s was deleted.\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
}

func convertPositiveInt(w http.ResponseWriter, req *http.Request, str string) (int, bool) {
	intValue, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "illegal value: %s", str)
		return 0, false
	}
	if intValue <= 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "illegal value: %s, set more than zero.", str)
		return 0, false
	}
	return int(intValue), true
}

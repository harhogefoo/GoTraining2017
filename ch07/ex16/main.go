package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"github.com/harhogefoo/go_training2017/ch07/ex16/eval"
)

func main() {
	fmt.Println("browse localhost:8000")
	http.HandleFunc("/", index)
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", 9); err != nil {
		log.Fatal(err)
	}
}

func calc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values, ok := r.Form["expr"]
	if !ok {
		http.Error(w, "no expr", http.StatusBadRequest)
		return
	}

	for _, v := range values {
		expr, err := eval.Parse(v)
		if err != nil {
			http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
			return
		}

		result := expr.Eval(eval.Env{})
		fmt.Fprintf(w, "%s = %g\n", v, result)
	}
}

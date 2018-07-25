package main

import (
	"database/sql"
	"fmt"
	"github.com/yiv/go-lab/database/injection/models"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models.AllBooks(env.db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type PageData struct {
	Title   string
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("Page/main.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:0808")
	connected()
	http.ListenAndServe(":0808", nil)
}

func connected() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 測試連接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MySQL database!")
}

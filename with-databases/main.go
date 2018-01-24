package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "firsttest:somepassword@tcp(mydbinstance.cdxjey534wpm.us-east-1.rds.amazonaws.com:3306)/new_schema?charset=utf8")

	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/friends", friends)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("now serving at localhost:8080")
	err = http.ListenAndServe(":8080", nil)

	check(err)

}

func friends(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM friends;`)
	check(err)
	defer rows.Close()

	var s, first, last string
	var id int
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&id, &last, &first)
		check(err)
		s += first + last + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customers (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customers", n)

}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE	customers;`)
	check(err)
	_, err = stmt.Exec()

	fmt.Fprintln(w, "DROPPED it like it's hot")
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customers VALUES ("James");`)
	check(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORDS", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM customers;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customers SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customers WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "Successfully connected!")
	check(err)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

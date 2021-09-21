package main

import (
	"encoding/json"
	"fmt"

	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "db"
	port = 3306
)

var (
	db       *sql.DB
	id       int
	date     string
	thisTodo string
	err      error
	user     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	dbname   = os.Getenv("MYSQL_DATABASE")
)

type todo struct {
	Todo string `json:"todo"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func testPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test page")
	fmt.Println("test page hit")
}

func crudPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		// get all todos in order they were submitted
		todos, err := db.Query("select * from crud.todos order by id desc;")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}
		res := []string{}
		for todos.Next() {
			err := todos.Scan(&id, &date, &thisTodo)
			checkError(err)
			res = append(res, thisTodo)
		}
		jsonRes, err := json.Marshal(res)
		checkError(err)
		w.Write(jsonRes)
		return

	case "POST":
		err = r.ParseForm()
		checkError(err)
		decoder := json.NewDecoder(r.Body)
		var t todo
		err = decoder.Decode(&t)
		checkError(err)
		q := fmt.Sprintf("insert into crud.todos (todo) values ('%s');", t.Todo)
		_, err := db.Query(q)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Println("post successful")
		return

	case "DELETE":
		q := "truncate table crud.todos"
		_, err := db.Query(q)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}
		w.WriteHeader(http.StatusAccepted)
		fmt.Println("clear successful")
		return
	}
}

func dbConnect() {
	sqlconn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	db, err = sql.Open("mysql", sqlconn)
	checkError(err)
	err = db.Ping()
	checkError(err)
	fmt.Println("Connected!")
}

func serve() {
	dbConnect()
	http.HandleFunc("/test", testPage)
	http.HandleFunc("/crud", crudPage)
	log.Fatal(http.ListenAndServe("0.0.0.0:5001", nil))
}

func main() {
	serve()
}

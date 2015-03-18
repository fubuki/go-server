package main

import (
	"database/sql"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/googollee/go-gcm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mikespook/gearman-go/client"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/test/:name", func(res http.ResponseWriter, req *http.Request) { // res 和 req 是由 Martini 注入
		res.WriteHeader(200) // HTTP 200
	})

	m.Get("/data", func(res http.ResponseWriter, req *http.Request) {
		table()
	}
	
	m.Get("/queue", func(res http.ResponseWriter, req *http.Request) {
		queue()
	})
	
	m.NotFound(func() {
		// handle 404
	})
	m.Run()
}

func gcm_sender() {
	client := gcm.New("key")
	load := gcm.NewMessage("device_id")
	load.AddRecipient("abc")
	load.SetPayload("data", "1")
	load.CollapseKey = "demo"
	load.DelayWhileIdle = true
	load.TimeToLive = 10
	resp, err := client.Send(load)
	fmt.Printf("id: %+v\n", resp)
	fmt.Println("err:", err)
	fmt.Println("err index:", resp.ErrorIndexes())
	fmt.Println("reg index:", resp.RefreshIndexes())
}

func table() {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
		delete from token;
		create table token (id integer not null primary key, name, token);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}

func queue() {
	c, err := client.New("tcp4", "127.0.0.1:4730")
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()
	c.ErrorHandler = func(e error) {
		log.Println(e)
	}

	jobHandler := func(resp *client.Response) {
		log.Printf("%s", resp.Data)
	}
}

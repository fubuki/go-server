package main

import (
	"database/sql"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/googollee/go-gcm"
	"github.com/kr/beanstalk"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)


func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		gcm_sender()
		return "ok"
	})

	m.Get("/test/:name", func(res http.ResponseWriter, req *http.Request) { // res 和 req 是由 Martini 注入
		res.WriteHeader(200) // HTTP 200
	})

	m.Get("/data", func(res http.ResponseWriter, req *http.Request) {
		table()
	}
	
	m.NotFound(func() {
		// handle 404
	})
	
	// 在請求前後加 logs
	m.Use(func(c martini.Context, log *log.Logger) {
		log.Println("before a request")
		c.Next()
		log.Println("after a request")
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

func reserve() {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		panic(err)
	}
	for true {
		name := "DailyCount"
		tube := beanstalk.NewTubeSet(c, name)
		id, body, err := tube.Reserve(10 * time.Second)
		if err != nil {
			//panic(err)
			fmt.Println(err)
			continue
		}
		fmt.Println("job", id)
		fmt.Println(string(body))
		c.Delete(id)
	}
}

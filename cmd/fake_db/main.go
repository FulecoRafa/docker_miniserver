package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"encoding/json"
)

// Global logger for http requests
var httplog *log.Logger = log.New(
	os.Stdout,
	"Endpoint hit: ",
	log.Ldate|log.LstdFlags|log.Lshortfile,
)

type User struct {
	Name   string
	Review float64
}

var users []User = []User{
	{Name: "John", Review: 4.5},
	{Name: "Jane", Review: 4.0},
}

type Travel struct {
	Owner User
	Rider User
	Time time.Time
}

var travels []Travel = []Travel{
	{Owner: users[0], Rider: users[1], Time: time.Now()},
}

func getUser(w http.ResponseWriter, r *http.Request) {
	httplog.Println("getUser")
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(users)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Fatal(err)
	}
	w.Write(res)
}

func getTravel(w http.ResponseWriter, r *http.Request) {
	httplog.Println("getTravel")
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(travels)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Fatal(err)
	}
	w.Write(res)
}

func main() {
	log.SetFlags(log.Ldate | log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/user", getUser)
	http.HandleFunc("/travel", getTravel)

	log.Println("Server up and listening 🚀")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}

}

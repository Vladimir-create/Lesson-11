package main

import (
	"net/http"
	"io"
	//"fmt"
	"time"
)

type (
	Session struct{
		ID string
		User_ID string
		ID_Adress string
		Opened_Data time.Time
		Expiration time.Time
	}
	User struct{
		ID string
		Username string
		Passwordhash string
	}
)

var Users []User
var Sessions []Session

func Handlerlogin(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	password := req.FormValue("password")
	User_id := "fff"
	Users = append(Users, User{ID:"User_id", Username:name, Passwordhash:password})	
	Sessions = append(Sessions, Session{ID:"http://localhost:8080", User_ID:"User_id", ID_Adress:"fsdf", Opened_Data:10, Expiration:20})	
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	for i:=0; i<len(Users); i++{
		if Users[i].Username == req.FormValue("name") && Users[i].Passwordhash == req.FormValue("password"){
			//get 
			return
		}
	}
	io.WriteString(w, "LOGIN pls!!!")
	http.HandleFunc("/login", Handlerlogin)
}


func main() {
	http.HandleFunc("/login", Handlerlogin)
	http.HandleFunc("/data", Handler2)
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}

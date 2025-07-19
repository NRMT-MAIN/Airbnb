package controllers

import "net/http"

func PingHandeler(w http.ResponseWriter , r *http.Request){
	w.Write([]byte ("Pong"))
}
package main

import (
	//"github.com/gorilla/context"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":3000")
}


func HomeHandler(w http.ResponseWriter, req *http.Request){
	
	fmt.Fprintf(w, "Hello World")
}

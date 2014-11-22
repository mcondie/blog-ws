package main

import (
	"github.com/gorilla/context"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"encoding/json"
)

func main(){

	db, err := sqlx.Connect("postgres", "user=blogws dbname=blogws_dev password=blogws sslmode=disable")
	Log(err)
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/blocks", Block_List)

	n := negroni.Classic()

	n.Use(negroni.HandlerFunc(
		func (rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
			context.Set(r, "db", db)
			next(rw, r)
		}))
	n.Use(negroni.HandlerFunc(JSONEncoderHandler))
	n.UseHandler(r)
	n.Run(":3000")
}


func HomeHandler(w http.ResponseWriter, req *http.Request){
	
	fmt.Fprintf(w, "Hello World")
}

func JSONEncoderHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	enc := json.NewEncoder(rw)
	context.Set(r, "enc", enc)
	next(rw, r)
}

func Log(err error){
	if err != nil {
        log.Fatalln(err)
    }
}

func GetDB(r *http.Request) *sqlx.DB {
    if rv := context.Get(r, "db"); rv != nil {
        return rv.(*sqlx.DB)
    }
    return nil
}

func GetENC(r *http.Request) *json.Encoder {
    if rv := context.Get(r, "enc"); rv != nil {
        return rv.(*json.Encoder)
    }
    return nil
}
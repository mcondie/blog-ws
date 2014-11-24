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
	Logerr(err)
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/blocks", Block_List).Methods("GET")
	r.HandleFunc("/blocks/", Block_List).Methods("GET")
	r.HandleFunc("/blocks", Block_Create).Methods("POST")
	r.HandleFunc("/blocks/", Block_Create).Methods("POST")
	r.HandleFunc("/blocks/{key}", Block_View).Methods("GET")
	r.HandleFunc("/blocks/{key}/", Block_View).Methods("GET")
	r.HandleFunc("/blocks/{key}", Block_Update).Methods("POST")
	r.HandleFunc("/blocks/{key}/", Block_Update).Methods("POST")
	r.HandleFunc("/blocks/{key}", Block_Delete).Methods("DELETE")
	r.HandleFunc("/blocks/{key}/", Block_Delete).Methods("DELETE")


	n := negroni.Classic()

	n.Use(negroni.HandlerFunc(
		func (rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
			context.Set(r, "db", db)
			next(rw, r)
		}))
	n.Use(negroni.HandlerFunc(JSONEncoderHandler))
	n.Use(negroni.HandlerFunc(ContentTypeHandler))


	n.UseHandler(r)
	n.Run(":3000")
}

func ContentTypeHandler(w http.ResponseWriter, req *http.Request, next http.HandlerFunc){
	w.Header().Set("Content-Type", "text/json")
	next(w, req)
}

func HomeHandler(w http.ResponseWriter, req *http.Request){
	
	fmt.Fprintf(w, "Hello World")
}

func JSONEncoderHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	enc := json.NewEncoder(rw)
	context.Set(r, "enc", enc)
	next(rw, r)
}

func Logerr(err error){
	if err != nil {
        log.Panicln(err)
    }
}

func Log(s string){
    log.Println(s)
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
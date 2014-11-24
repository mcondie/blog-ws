package main

import (
	"net/http"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
)

type Block struct{
	Content_id	*int 		`json:"content_id"`
	Content 	string		`json:"content"`
	Key		string		`json:"key"`
}

func Block_List(w http.ResponseWriter, req *http.Request){
	db := GetDB(req)
	blocks := []Block{}
    db.Select(&blocks, "SELECT * FROM blocks")

    enc := GetENC(req)
    enc.Encode(blocks)
}

func Block_View(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	key := vars["key"]

	Log("key: " + key)

	db := GetDB(req)

	var block Block
	err := db.Get(&block, "SELECT * FROM blocks where key = $1", key)
	Logerr(err)

	enc := GetENC(req)
	enc.Encode(block)
}

func Block_Create(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
	var block Block

	err := decoder.Decode(&block)
	Logerr(err)

	db := GetDB(req)
	var id int
    err = db.QueryRow("INSERT INTO blocks (content, key) VALUES ($1, $2) RETURNING content_id", block.Content, block.Key).Scan(&id)
    Logerr(err)

    block.Content_id = &id

    enc := GetENC(req)
    enc.Encode(block)
}

func Block_Update(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
	var block Block

	err := decoder.Decode(&block)
	Logerr(err)

	vars := mux.Vars(req)
	key := vars["key"]

	if(key != block.Key){
		Logerr(errors.New("keys do not match"))
	}

	db := GetDB(req)
    _, err = db.Exec("UPDATE blocks SET content = $1 where key = $2", block.Content, block.Key)
    Logerr(err)

    enc := GetENC(req)
    enc.Encode(block)
}

func Block_Delete(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	key := vars["key"]

	db := GetDB(req)
    _, err := db.Exec("DELETE from blocks where key = $1", key)
    Logerr(err)


}
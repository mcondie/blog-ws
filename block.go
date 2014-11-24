package main

import (
	"net/http"
)

type Block struct{
	Content_id	int
	Content 	string
	Name		string
}

func Block_List(w http.ResponseWriter, req *http.Request){
	db := GetDB(req)
	blocks := []Block{}
    db.Select(&blocks, "SELECT * FROM blocks")

    enc := GetENC(req)
    enc.Encode(blocks)
}

func Block_View(w http.ResponseWriter, req *http.Request){
	
}

func Block_Create(w http.ResponseWriter, req *http.Request){
	
}

func Block_Update(w http.ResponseWriter, req *http.Request){
	
}

func Block_Delete(w http.ResponseWriter, req *http.Request){
	
}
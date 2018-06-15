package main

import (
	"net/http"
	"book_borrow/config"
	"book_borrow/handlers"
	_ "book_borrow/models"
)

func main() {
	for path, handleFunc := range handlers.Router{
		http.HandleFunc(path, handleFunc)
	}
	
	server := http.Server {
		Addr: config.SERVER_ADDR,
	}
	
	server.ListenAndServe()
}


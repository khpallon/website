package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// socketReader struct
type socketReader struct {
	con  *websocket.Conn
	mode int
	name string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var savedsocketreader []*socketReader

func socketReaderCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("socket request")
	if savedsocketreader == nil {
		savedsocketreader = make([]*socketReader, 0)
	}

	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		r.Body.Close()

	}()
	con, _ := upgrader.Upgrade(w, r, nil)

	ptrSocketReader := &socketReader{
		con: con,
	}

	savedsocketreader = append(savedsocketreader, ptrSocketReader)
}

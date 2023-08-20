// Copyright (c) Christian Surlykke
//
// This file is part of the RefudeServices project.
// It is distributed under the GPL v2 license.
// Please refer to the GPL2 file for a copy of the license.
package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)


var pingHandler = websocket.Handler(func(conn *websocket.Conn) {
	var ticker = time.NewTicker(time.Second*10)	
	defer ticker.Stop()
	for range ticker.C {
		if err := websocket.JSON.Send(conn, "PING"); err != nil {
			fmt.Println("Got error sending:", err)
			return
		}
	}
})

func main() {
	http.Handle("/ping", pingHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("http.ListenAndServe failed:", err)
	}
}


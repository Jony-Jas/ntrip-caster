package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

func FrontService(){
	http.HandleFunc("/admin", frontHandler)
	fmt.Println("WebSocket server listening on :8080/admin")
	http.ListenAndServe(":8080", nil)
}

func frontHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New connection from front end")

	for {
		data := map[string]interface{}{
			"BaseStations":BaseStations,
			"Users":Users,
		}
	
		jsonData, err := json.Marshal(data)
		fmt.Println(string(jsonData))
		if err != nil {
			fmt.Println(err)
			return
		}
	
		err = conn.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(1 * time.Second)
	}

}
package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
	"strconv"
)



func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()

	params := r.URL.Query()
	name := params.Get("name")
	x,_:= strconv.Atoi(params.Get("x"))
	y,_:= strconv.Atoi(params.Get("y"))

	fmt.Printf("New connection from BaseStation: %s\n", name)
	curBaseStation := BaseStation{name, x, y, 0, 0}
	BaseStations = append(BaseStations, curBaseStation)

	go handleMessages(conn, name, &BaseStations[len(BaseStations)-1])
}

func handleMessages(conn *websocket.Conn, name string, curBaseStation *BaseStation) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message from %s: %s\n", name, err)
			removeBaseStation(*curBaseStation)
			return
		}
		
		// Convert the received JSON message to a Go struct
		var message BaseStationMessage
		if err := json.Unmarshal(p, &message); err != nil {
			fmt.Printf("Error decoding JSON message: %s\n", err)
			return
		}
		(*curBaseStation).CrtX = message.CrtX
		(*curBaseStation).CrtY = message.CrtY

		fmt.Printf("Received message from %s: %+v\n", curBaseStation.Name, message)
	}
}

func removeBaseStation(target BaseStation) {
	// Remove a base station from the list
	for i, bs := range BaseStations {
		if bs.Name == target.Name {
			BaseStations = append(BaseStations[:i], BaseStations[i+1:]...)
			break
		}
	}
}

func ReadBaseStation() {
	http.HandleFunc("/base-station", handler)
	fmt.Println("WebSocket server listening on :8080/base-station")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	// "strconv"
)

func ReadAndSendUser() {
	http.HandleFunc("/user", userHandler)
	fmt.Println("WebSocket server listening on :8080/user")
	http.ListenAndServe(":8080", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()

	params := r.URL.Query()
	name := params.Get("name")
	fmt.Printf("New user : %s\n", name)
	newUser := User{name,0,0,""}
	Users = append(Users, newUser)

	go handleUserMessage(conn,name, &BaseStations, &Users[len(Users)-1])
}

func handleUserMessage(conn *websocket.Conn, name string, bs *[]BaseStation, user *User){
	for{
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message from %s: %s\n", name, err)
			removeUser(user);
			return
		}

		var message UserMessage
		if err := json.Unmarshal(p, &message); err != nil {
			fmt.Printf("Error decoding JSON message: %s\n", err)
			return
		}
		x := message.PosX
		y := message.PosY
		(*user).PosX = x
		(*user).PosY = y
		fmt.Printf("Received position from %s: %+v\n", name, message)

		pos := -1
		min := 1000000
		for j:=0; j<len(*bs); j++ {
			dist := (x-(*bs)[j].PosX)*(x-(*bs)[j].PosX) + (y-(*bs)[j].PosY)*(y-(*bs)[j].PosY)
			if dist < min {
				min = dist
				pos = j
			}
		}

		if(pos!=-1){
			fmt.Printf("Sending to %s from: %v\n", name, (*bs)[pos].Name)
			(*user).Bs = (*bs)[pos].Name
		}else{
			fmt.Printf("Sending to %s: No baseStation found\n", name)
			(*user).Bs = ""
		}
	}
}

func removeUser(target *User) {
	// Remove a user from the list
	for i, bs := range Users {
		if bs.Name == target.Name {
			Users = append(Users[:i], Users[i+1:]...)
			break
		}
	}
}
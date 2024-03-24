package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

type BaseStationMessage struct {
	CrtX int `json:"crtX"`
	CrtY int `json:"crtY"`
}

type UserMessage struct {
	PosX int `json:"posX"`
	PosY int `json:"posY"`
}

type BaseStation struct {
	Name string
	PosX int
	PosY int
	CrtX int
	CrtY int
}
type User struct {
	Name string
	PosX int
	PosY int
	Bs string
}
var BaseStations []BaseStation;
var Users []User;

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main(){
	fmt.Println("Starting NTRIP Caster")

	go ReadBaseStation()
	go ReadAndSendUser()
	go FrontService()
	go func(){
		for{
			fmt.Println(BaseStations)
			fmt.Println(Users)
			time.Sleep(1 * time.Second)
		}
	}()

	for{}
}
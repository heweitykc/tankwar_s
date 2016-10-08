package main

import (
	"log"
)

type Room struct {
	roomid   	uint64
	players  	map[uint64]*Player
}

func (this *Room) Log(){
	log.Print("room id ...", this.roomid)
}

func (this *Room) Create(roomid uint64){
	this.roomid = roomid
	this.players = make(map[uint64]*Player)
}

func (this *Room) AddUser(uid uint64){
	newplayer := new(Player)
	this.players[uid] = newplayer
	newplayer.Create(uid)
}

func (this *Room) Loop() {
	for _, player := range this.players {
		player.Loop()
	}
}

func (this *Room) IsFull() bool {
	return false
}

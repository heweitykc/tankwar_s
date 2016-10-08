package main

import (
	"log"
)

const (
	MAX_PLAYER = 10
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

func (this *Room) AddUser(newplayer *Player){	
	this.players[newplayer.uid] = newplayer	
}

func (this *Room) Loop() {
	for _, player := range this.players {
		player.Loop()
	}	
}

func (this *Room) IsFull() bool {	
	return len(this.players) < MAX_PLAYER
}

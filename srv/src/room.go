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
	newplayer.room = this
}

func (this *Room) RemoveUser(uid uint64){	
	delete(this.players, uid)
}

func (this *Room) Update(dt float64) {
	for _, player := range this.players {
		player.Update(dt)
	}
}

func (this *Room) FixedUpdate(dt float64) {
	for _, player := range this.players {
		player.FixedUpdate(dt)
	}
}

func (this *Room) IsFull() bool {	
	return len(this.players) < MAX_PLAYER
}

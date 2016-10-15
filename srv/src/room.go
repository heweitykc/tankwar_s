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
	this.syncUserEnter(newplayer.uid)
	this.syncRoom(newplayer.uid)
}

func (this *Room) RemoveUser(uid uint64){	
	delete(this.players, uid)
	this.syncUserLeave(uid)
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
	length := len(this.players)
	log.Print("players.length = ", length)
	return length >= MAX_PLAYER
}

//同步新玩家的房间数据
func (this *Room) syncRoom(uid_enter uint64){
	log.Print("syncRoom : uid_enter=", uid_enter)
	msg := make(map[string]interface{})
	msg["id"] = MSG_ENTER_
	msg["room"] = make(map[uint64]interface{})
	msg["init"] = 1
	for _, player := range this.players {	
		var roominfo = msg["room"].(map[uint64]interface{})		
		roominfo[player.uid] = make(map[string]interface{})
		var playerinfo = roominfo[player.uid].(map[string]interface{})		
		fillJsonData(player, playerinfo)
	}
	SendPlayerMsg(uid_enter, &msg)
}

//同步新玩家到其他玩家
func (this *Room) syncUserEnter(uid_enter uint64){
	msg := make(map[string]interface{})
	msg["id"] = MSG_ENTER_
	msg["uid"] = uid_enter
	msg["enter"] = 1
	fillJsonData(this.players[uid_enter], msg)
	for _, player := range this.players {
		log.Print("player.uid = ", player.uid , "  roomid = ", this.roomid)
		 if player.uid == uid_enter {
		 	continue
		 }
		SendPlayerMsg(player.uid, &msg)
	}
}

//同步离开玩家到其他玩家
func (this *Room) syncUserLeave(uid_leave uint64){
	msg := make(map[string]interface{})
	msg["id"] = MSG_ENTER_
	msg["uid"] = uid_leave
	msg["leave"] = 1
	for _, player := range this.players {
		log.Print("player.uid = ", player.uid, "  roomid = ", this.roomid)
		 if player.uid == uid_leave {
		 	continue
		 }
		SendPlayerMsg(player.uid, &msg)
	}
}

func fillJsonData(p *Player, msg map[string]interface{}){
	msg["uid"] 	  = p.uid
	msg["x"] 	  = p.x
	msg["y"] 	  = p.y
	msg["rot"]    = p.rot
	msg["cannon"] = p.cannon
	msg["health"] = p.health
	msg["speedx"] = p.speedx
	msg["speedy"] = p.speedy
	msg["energy"] = p.energy	
}

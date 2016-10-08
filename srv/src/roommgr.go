package main

import (
	"log"
	"github.com/gorilla/websocket"
)

var __roomMgr *RoomMgr

func RoomManagerPtr() *RoomMgr {
	if __roomMgr == nil {
		__roomMgr = new(RoomMgr).Create()		
	}
	return __roomMgr
}

type RoomMgr struct {
	roomcount uint64
	playercount uint64
	rooms  map[uint64]*Room
	players  	map[uint64]*Player
}

func (this *RoomMgr) Create() *RoomMgr {
	this.rooms = make(map[uint64]*Room)
	this.roomcount = 100
	//this.roomIdGen = new(util.UniqueIDGenerator).Create()
	log.Print("RoomMgr Init")
	return this
}

func (this *RoomMgr) AddUser(uid uint64, conn *websocket.Conn){
	this.playercount++
	
	newplayer := new(Player)
	newplayer.Create(uid, conn)
	
	this.players[uid] = newplayer	
	log.Print("new player added ", uid)
	
	this.SearchRoomAndInsertUser(uid)
}

func (this *RoomMgr) SearchRoomAndInsertUser(uid uint64) {
	var availableRoom *Room
	for _, room := range this.rooms {
		if room.IsFull() == false {
			availableRoom = room
			break
		}
	}
	
	this.roomcount++
	roomid := this.roomcount
	availableRoom = new(Room)
	availableRoom.Create(roomid)
	this.rooms[roomid] = availableRoom
	
	availableRoom.AddUser(this.players[uid])
}

func (this *RoomMgr) Loop() {
	for _, room := range this.rooms {
		room.Loop()
	}
}

func (this *RoomMgr) HandleNetMsg(msg map[string]interface{}){		
	var msgid = msg["id"]
	log.Print("id=",  msgid)
}
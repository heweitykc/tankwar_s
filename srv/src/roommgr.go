package main

import (
	"log"
)

var __roomMgr *RoomMgr

func RoomManagerPtr() *RoomMgr {
	if __roomMgr == nil {
		__roomMgr = new(RoomMgr).Create()		
	}
	return __roomMgr
}

type RoomMgr struct {
	count uint64
	rooms  map[uint64]*Room
}

func (this *RoomMgr) Create() *RoomMgr {
	this.rooms = make(map[uint64]*Room)
	this.count = 100
	//this.roomIdGen = new(util.UniqueIDGenerator).Create()
	log.Print("RoomMgr Init")
	return this
}

func (this *RoomMgr) SearchRoomAndInsertUser(uid uint64) {
	var availableRoom *Room
	for _, room := range this.rooms {
		if room.IsFull() == false {
			availableRoom = room
			break
		}
	}
		
	this.count++
	roomid := this.count
	availableRoom = new(Room)
	availableRoom.Create(roomid)
	this.rooms[roomid] = availableRoom
	
	availableRoom.AddUser(uid)
	
	log.Print("new player added ", uid)
}

func (this *RoomMgr) Loop() {
	for _, room := range this.rooms {
		room.Loop()
	}
}

func (this *RoomMgr) HandleNetMsg(msg map[string]interface{}){	
	log.Print("id=",  msg["id"])
	
}
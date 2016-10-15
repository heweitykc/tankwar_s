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
	roomcount uint64
	playercount uint64
	rooms  		map[uint64]*Room
	players  	map[uint64]*Player
}

func (this *RoomMgr) Create() *RoomMgr {
	this.rooms = make(map[uint64]*Room)
	this.players = make(map[uint64]*Player)
	
	this.roomcount = 100
	//this.roomIdGen = new(util.UniqueIDGenerator).Create()
	log.Print("RoomMgr Init")
	return this
}

func (this *RoomMgr) AddUser(uid uint64){
	this.playercount++
	
	newplayer := new(Player)
	newplayer.Create(uid)
	
	this.players[uid] = newplayer	
	log.Print("new player added ", uid)
	
	this.SearchRoomAndInsertUser(uid)	
}

func (this *RoomMgr) RemoveUser(uid uint64){
	p := this.players[uid] 
	if p == nil {
		return
	}
	p.Dispose()
	delete(this.players, uid)
	
	log.Print("remove user")
}

func (this *RoomMgr) SearchRoomAndInsertUser(uid uint64) {
	var availableRoom *Room
	length := len(this.rooms)
	log.Print("rooms.length = ", length)
	
	for _, room := range this.rooms {		
		if room.IsFull() == false {
			availableRoom = room
			break
		}
	}
	
	this.roomcount++
	roomid := this.roomcount
	if availableRoom == nil {
		availableRoom = new(Room)
		availableRoom.Create(roomid)
		this.rooms[roomid] = availableRoom
		log.Print("new room = ", roomid)	
	}	
	availableRoom.AddUser(this.players[uid])
}

func (this *RoomMgr) Update(dt float64) {
	for _, room := range this.rooms {
		room.Update(dt)
	}
}

func (this *RoomMgr) FixedUpdate(dt float64) {
	for _, room := range this.rooms {
		room.FixedUpdate(dt)
	}
}

func (this *RoomMgr) HandleNetMsg(msg map[string]interface{}, current_uid uint64){			
	var msgid = uint32(msg["id"].(float64))
	log.Print("id=",  msgid, "  uid=", current_uid)
	if msgid == MSG_ENTER {
		log.Print("AddUser=",  current_uid)
		this.AddUser(current_uid)		
	} else {
		this.players[current_uid].HandleNetMsg(msg)
	}			
}

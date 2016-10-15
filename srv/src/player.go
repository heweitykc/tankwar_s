package main

import (
	"log"
)

type Player struct {
	// 
	uid   		uint64
	
	// spatial attribute
	x			int32
	y			int32
	rot			int32
	cannon		int32
	
	// property
	health		int32
	speedx		int32
	speedy		int32
	energy		int32
	
	status		int32
	
	room		*Room
}

func (this *Player) Log(){
	log.Print("player id ...", this.uid)
}

func (this *Player) Create(uid uint64){
	this.uid = uid	

	this.x = 0	
	this.y =  0	
	this.rot = 0
	this.cannon = 0
	
	this.health = 100	
	this.speedx =  0	
	this.speedy =  0	
	this.energy = 100	
}

func (this *Player) Update(dt float64) {
	this.x += this.speedx
	this.y += this.speedy
}

func (this *Player) FixedUpdate(dt float64) {
	log.Print(this.uid, " FixedUpdate")
}

func (this *Player) Dispose() {
	this.room.RemoveUser(this.uid)
}

func (this *Player) HandleNetMsg(msg map[string]interface{}){			
	var msgid = uint32(msg["id"].(float64))	
	if msgid == MSG_MOVE {
		
	} else if msgid == MSG_FIRE {
		
	}
}





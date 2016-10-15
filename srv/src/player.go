package main

import (
	"log"
)

type Player struct {
	// 
	uid   		uint64	
	
	// spatial attribute
	x			uint32
	y			uint32
	rot			uint32
	cannon		uint32
	
	// property
	health		int32
	speed		int32
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
	this.speed =  1	
	this.energy = 100	
}

func (this *Player) Update(dt float64) {
	
}

func (this *Player) FixedUpdate(dt float64) {
	log.Print(this.uid, " FixedUpdate")
}

func (this *Player) Dispose() {
	this.room.RemoveUser(this.uid)
}





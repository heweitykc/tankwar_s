package main

import (
	"log"
	//"github.com/gorilla/websocket"
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
}

func (this *Player) Log(){
	log.Print("player id ...", this.uid)
}

func (this *Player) Create(uid uint64){
	this.uid = uid	
	this.health = 100	
	this.speed =  1	
	this.energy = 100	
}

func (this *Player) Loop() {
	
}




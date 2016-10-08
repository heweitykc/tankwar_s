package main

import (
	"flag"
	"log"
	"time"	
	"net/http"
	"utilities"	
	"encoding/json"	
	"github.com/gorilla/websocket"
)

const (	
	MaxPackageLen uint32 = 200000
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")
var upgrader = websocket.Upgrader{}
var roomMgr *RoomMgr

func mainloop(){
	roomMgr = RoomManagerPtr()
	timeTicker := time.NewTicker(time.Millisecond * time.Duration(int32(1000 / 15)))
	defer func() {
		timeTicker.Stop()
		log.Print("loop end")		
	}()
	
	for {
		select {
			case <-timeTicker.C:
				roomMgr.Loop()
		}
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()	
	deliveries := utilities.ParseMessage(c, MaxPackageLen, utilities.WS_REQUEST)
	
	for {
		select {
			case recvData, _ := <-deliveries:
				j2 := make(map[string]interface{})
				err = json.Unmarshal(recvData, &j2)
				roomMgr.HandleNetMsg(j2)
		}
	}
}

func main() {	
	go mainloop()	
	http.HandleFunc("/game", socketHandler)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

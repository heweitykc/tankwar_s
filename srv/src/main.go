package main

import (
	"flag"
	"log"
	"os"
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
var uid_count = (uint64)(0)

func mainloop(){
	roomMgr = RoomManagerPtr()
	logicTimeTicker := time.NewTicker(time.Millisecond * time.Duration(int32(1000)))
	timeTicker := time.NewTicker(time.Millisecond * time.Duration(int32(1000/15)))
	defer func() {
		timeTicker.Stop()
		log.Print("loop end")
	}()
	
	for {
		select {
			case <-logicTimeTicker.C:
				roomMgr.Update(1000/15)
				
			case <-timeTicker.C:
				roomMgr.FixedUpdate(1000)				
		}
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	uid_count++
	current_uid := uid_count
	defer func() {
		roomMgr.RemoveUser(current_uid)
		c.Close()	
	}()
	deliveries := utilities.ParseMessage(c, MaxPackageLen, utilities.WS_REQUEST)
		
	log.Print("player connected current_uid=", current_uid)
	for {
		select {
			case recvData, ok := <-deliveries:
				if !ok {
					goto Finish_Label
				}
				log.Print("recvData=", recvData)
				msg := make(map[string]interface{})
				err = json.Unmarshal(recvData, &msg)
				roomMgr.HandleNetMsg(msg, current_uid)
		}
	}
Finish_Label:
	log.Println("conn close() ", current_uid)	
}

func main() {	
	fileName := "xxx_debug.log"
    logFile,_  := os.Create(fileName)
	debugLog := log.New(logFile,"[Debug]",log.Llongfile)
	debugLog.Println("A debug message here")
    debugLog.SetPrefix("[Info]")
    debugLog.Println("A Info Message here ")
    debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
    debugLog.Println("A different prefix")
	
	go mainloop()	
	http.HandleFunc("/game", socketHandler)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

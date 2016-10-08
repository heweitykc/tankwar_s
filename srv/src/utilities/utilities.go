package utilities

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"github.com/gorilla/websocket"
)

const (
	PacageHeadLen uint32 = 2 + 4 + 4
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

type WS_TYPE int

const (
	WS_REQUEST WS_TYPE = 0x01
	WS_RESPOND WS_TYPE = 0x02
)

func SendMessage(c *websocket.Conn, contents *map[string]interface{}, magic WS_TYPE) error {
	/*
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic catched SendMessage(...)")
				log.Println(err)
			}
		}()
	*/

	js1, err := json.Marshal(contents)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, []byte{0x01, byte(magic)})
	binary.Write(buf, binary.BigEndian, uint32(len(js1)))
	binary.Write(buf, binary.BigEndian, []byte{0, 0, 0, 0})
	binary.Write(buf, binary.BigEndian, []byte(js1))

	return c.WriteMessage(websocket.TextMessage, []byte(buf.Bytes()))
}

func ParseMessage(c *websocket.Conn, maxPackageLen uint32, magic WS_TYPE) <-chan []byte {
	deliveries := make(chan []byte)
	go func() {
		defer close(deliveries)
		defer func() {
			log.Println("panic catched ParseMessage(...)")
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		for {
			_, buff, err := c.ReadMessage()

			if err != nil {
				log.Println("read:", err)
				break
			}

			if 0 == bytes.Compare(buff[:2], []byte{0x01, byte(magic)}) { // magic num
				var packageLen uint32
				var reserved uint32
				binary.Read(bytes.NewBuffer(buff[2:6]), binary.BigEndian, &packageLen)
				binary.Read(bytes.NewBuffer(buff[6:10]), binary.BigEndian, &reserved)
				if 0 == maxPackageLen || packageLen < maxPackageLen {
					message := buff[PacageHeadLen:]
					if int(packageLen) == len(message) {
						//						log.Printf("recv: %s", message)
						//						continue
						deliveries <- message
						continue
					}
				} else {
					log.Printf("over MaxPackageLen: %d,drop whole package", packageLen)
					continue
				}
			}
			break
		}
	}()
	return deliveries

}

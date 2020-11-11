package app

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
)


type MsgNo int
const (
	HeartBeat MsgNo = iota
	Go
	ReadBufferSize = 1024
)

type Message struct {
	MsgNo MsgNo
	Data  []byte
}

var (
	HeartBeatMsg = &Message{MsgNo: HeartBeat}
	GoMsg = &Message{MsgNo: Go}
)

type SlaveProcessMaster struct {
	conn        net.Conn
	sendMsgLock sync.Mutex
}

func NewSlaveProcessMaster(conn net.Conn) *SlaveProcessMaster {
	ret := &SlaveProcessMaster{
		conn: conn,
	}
	return ret
}


func (s *SlaveProcessMaster) ProcessMasterMsg() {
	if s.conn != nil {
		for {
			rBuf := make([]byte, ReadBufferSize)
			read, err := s.conn.Read(rBuf)
			if err != nil {
				os.Exit(-1)
			}

			content := &Message{}

			errDes := json.Unmarshal(rBuf[0:read], content)
			if errDes != nil {
				fmt.Println(string(rBuf[0:read]))
				panic(errDes)
			}

			switch content.MsgNo {
			case Go:
				fmt.Println("slave go")
			case HeartBeat:
				fmt.Println("heart beat")
			}
		}
	}
}

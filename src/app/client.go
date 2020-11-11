package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
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
	GoMsg        = &Message{MsgNo: Go}
)

type SlaveProcessMaster struct {
	conn net.Conn
}

func NewSlaveProcessMaster(conn net.Conn) *SlaveProcessMaster {
	ret := &SlaveProcessMaster{
		conn: conn,
	}
	return ret
}

func (s *SlaveProcessMaster) ProcessMasterMsg() {
	br := bufio.NewReader(s.conn)
	dec := json.NewDecoder(br)
	for {
		var content Message
		err := dec.Decode(&content)
		if err != nil {
			panic(err)
		}

		switch content.MsgNo {
		case Go:
			fmt.Println("slave go")
		case HeartBeat:
			fmt.Println("heart beat")
		}
	}
}

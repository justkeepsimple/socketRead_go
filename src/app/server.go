package app

import (
	"encoding/json"
	"net"
	"strconv"
	"sync"
)


func CreateListen() (net.Listener, string) {
	temps := 0
	port := 7475
	for {
		listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			if temps > 5 {
				panic(err)
			}
			temps++
			port++
			continue
		}
		return listen, strconv.Itoa(port)
	}

}






func NewMasterProcessSlave(conn net.Conn) *masterProcessSlave {
	ret := &masterProcessSlave{conn: conn}
	return ret
}


type masterProcessSlave struct {
	conn            net.Conn
	sendMsgLock     sync.Mutex
}

func (sc *masterProcessSlave) HeartBeat() {
	hb, err := json.Marshal(HeartBeatMsg)
	if err != nil {
		panic(err)
	}
	for {
		//time.Sleep(time.Second * 2)
		sc.sendMsgLock.Lock()
		_, errTcp := sc.conn.Write(hb)
		sc.sendMsgLock.Unlock()
		if errTcp != nil {
			panic(errTcp)
		}
	}
}


func (sc *masterProcessSlave) Go() {
	hb, err := json.Marshal(GoMsg)
	if err != nil {
		panic(err)
	}
	for {
		//time.Sleep(time.Second * 1)
		sc.sendMsgLock.Lock()
		_, errTcp := sc.conn.Write(hb)
		sc.sendMsgLock.Unlock()
		if errTcp != nil {
			panic(errTcp)
		}
	}
}


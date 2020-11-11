package main

import (
	"connRead/src/app"
	"fmt"
	"net"
	"time"
)

func main() {
	//create a server
	listen, port := app.CreateListen()
	fmt.Println(port)

	//waiting for client connect
	go func() {
		for {
			client, err := listen.Accept()
			if err != nil {
				panic(err)
			}
			mps := app.NewMasterProcessSlave(client)
			//this to send msgBeatHeartMsg
			go mps.HeartBeat()
			go mps.Go()
		}
	}()

	go doClient(port)
	time.Sleep(time.Second * 1000)
}

func doClient(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		panic(err)
	}
	client := app.NewSlaveProcessMaster(conn)
	client.ProcessMasterMsg()
}

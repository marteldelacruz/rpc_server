package main

import (
	"fmt"
	"net"
	"net/rpc"

	Admin "./admin"
	Util "./util"
)

func server() {
	rpc.Register(new(Admin.Server))
	ln, err := net.Listen(Util.PROTOCOL, Util.PORT)
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("New conection...")
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	Util.ScanString()
}

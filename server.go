package main

import (
	"fmt"
	"net"
	"net/rpc"

	Admin "./admin"
	Util "./util"
)

func server() {
	rpc.Register(new(Admin.Admin))
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
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	Util.ScanString()
}

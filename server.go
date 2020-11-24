package main

import (
	"fmt"
	"net"
	"net/rpc"

	Util "./util"
)

type Server struct{}

func (this *Server) Hello(name string, reply *string) error {
	*reply = "Hello " + name
	return nil
}

func server() {
	rpc.Register(new(Server))
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

	var input string
	fmt.Scanln(&input)
}

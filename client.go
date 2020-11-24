package main

import (
	"fmt"
	"net/rpc"

	Util "./util"
)

func client() {
	var opt string
	c, err := rpc.Dial(Util.PROTOCOL, Util.PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Println("1) Hello")
		fmt.Println("0) Exit")
		fmt.Scanln(&opt)

		switch opt {
		case "1":
			var name string
			fmt.Print("Name: ")
			fmt.Scanln(&name)

			var result string
			err = c.Call("Server.Hello", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Hello =", result)
			}
		case "0":
			return
		}
	}
}

func main() {
	client()
}

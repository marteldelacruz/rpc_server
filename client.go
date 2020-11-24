package main

import (
	"fmt"
	"net/rpc"
	"strconv"

	Util "./util"
)

func client() {
	var opt, result string
	var args Util.Args

	// inits server conection
	c, err := rpc.Dial(Util.PROTOCOL, Util.PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.Call("Admin.Init", nil, &result)
	if err != nil {
		fmt.Println(err)
	}

	for {
		fmt.Println("1) Assign subject grade")
		fmt.Println("2) Student average grade")
		fmt.Println("3) General average")
		fmt.Println("4) Subject general average")
		fmt.Println("0) Exit")
		fmt.Scanln(&opt)

		switch opt {
		// add new student, grade and subject
		case "1":
			fmt.Print("Subject: ")
			args.Subject = Util.ScanString()
			fmt.Print("Name: ")
			args.Name = Util.ScanString()
			fmt.Print("Grade: ")
			args.Grade, _ = strconv.ParseFloat(Util.ScanString(), 64)

			err = c.Call("Admin.Add", args, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Student data added!")
			}
			break
		// get student average grade
		case "2":
			fmt.Print("Name: ")
			name := Util.ScanString()

			err = c.Call("Admin.StudentAverage", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.StudentAverage =", result)
			}
			break
		case "0":
			return
		}
	}
}

func main() {
	client()
}

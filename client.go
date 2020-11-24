package main

import (
	"fmt"
	"net/rpc"
	"strconv"

	Admin "./admin"
	Util "./util"
)

func client() {
	var opt, result string
	var avrg float64

	// init data
	var args Admin.Args

	// inits server conection
	c, err := rpc.Dial(Util.PROTOCOL, Util.PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.Call("Server.Init", args, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SERVER => " + result)
	}

	for {
		fmt.Println("\n------------------------------------")
		fmt.Println("1) Assign subject grade")
		fmt.Println("2) Student average grade")
		fmt.Println("3) General average")
		fmt.Println("4) Subject general average")
		fmt.Println("0) Exit")
		fmt.Print("\nSelect an option: ")
		opt = Util.ScanString()

		switch opt {
		// add new student, grade and subject
		case "1":
			fmt.Print("Subject: ")
			args.Subject = Util.ScanString()
			fmt.Print("Name: ")
			args.Name = Util.ScanString()
			fmt.Print("Grade: ")
			args.Grade, _ = strconv.ParseFloat(Util.ScanString(), 64)

			err = c.Call("Server.Add", args, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("SERVER => " + result)
			}
			break
		// get student average grade
		case "2":
			fmt.Print("Name: ")
			args.Name = Util.ScanString()

			err = c.Call("Server.StudentAverage", args, &avrg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("SERVER => ", avrg)
			}
			break
		// get subject average grade
		case "3":
			fmt.Print("Subject: ")
			args.Subject = Util.ScanString()

			err = c.Call("Server.SubjectAverage", args, &avrg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("SERVER => ", avrg)
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

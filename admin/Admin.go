package admin

import "fmt"

type Data struct {
	Subject map[string]map[string]float64
	Student map[string]map[string]float64
}

type Args struct {
	Name    string
	Subject string
	Grade   float64
}

type Server struct {
	Maps Data
}

// Inits the maps
func (server *Server) Init(args *Args, reply *string) error {
	(*server).Maps.Student = make(map[string]map[string]float64)
	(*server).Maps.Subject = make(map[string]map[string]float64)

	*reply = "Successful creation of maps!"
	return nil
}

// Adds a new stutent to the map thing
func (server *Server) Add(args Args, reply *string) error {
	// create new subject
	(*server).Maps.Student[args.Name] = make(map[string]float64)
	(*server).Maps.Student[args.Name][args.Subject] = args.Grade

	// create new subject
	(*server).Maps.Subject[args.Subject] = make(map[string]float64)
	(*server).Maps.Subject[args.Subject][args.Name] = args.Grade

	*reply = "New student added!"
	fmt.Println(len((*server).Maps.Subject))
	return nil
}

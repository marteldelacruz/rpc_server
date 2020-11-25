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
	fmt.Println(*reply)
	return nil
}

// Adds a new stutent to the map thing
func (server *Server) Add(args Args, reply *string) error {
	// create new student if not exist
	if (*server).Maps.Student[args.Name] == nil {
		(*server).Maps.Student[args.Name] = make(map[string]float64)
	}
	// check if grade is already saved
	(*server).Maps.Student[args.Name][args.Subject] = args.Grade

	// verificar que la materia no exista antes de reasignar
	// create new subject
	if (*server).Maps.Subject[args.Subject] == nil {
		(*server).Maps.Subject[args.Subject] = make(map[string]float64)
	}
	(*server).Maps.Subject[args.Subject][args.Name] = args.Grade

	*reply = "New student added!"
	fmt.Println(*reply)
	return nil
}

// Returns the student average grade of all its subjects
func (server *Server) StudentAverage(args Args, reply *float64) error {
	*reply = 0
	fmt.Println("Client ask for student average")

	for _, g := range (*server).Maps.Student[args.Name] {
		*reply += g
	}

	*reply = *reply / float64(len((*server).Maps.Student[args.Name]))
	return nil
}

// Returns the subject average grade of all its students
func (server *Server) SubjectAverage(args Args, reply *float64) error {
	*reply = 0
	fmt.Println("Client ask for subject average")

	for _, g := range (*server).Maps.Subject[args.Subject] {
		*reply += g
	}

	*reply = *reply / float64(len((*server).Maps.Subject[args.Subject]))

	return nil
}

// Returns the average of all students and subjects
func (server *Server) GeneralAverage(string, reply *float64) error {
	*reply = 0

	/**for _, student := range admin.Subject {
		for _, grade := range admin.Student[student] {
			avrg += grade
		}
	}**/

	return nil
}

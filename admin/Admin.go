package admin

import (
	Util "../util"
)

type Admin struct {
	Subject map[string]map[string]float64
	Student map[string]map[string]float64
}

// Inits the subjects data maps
func (admin *Admin) Init(int, reply *string) error {
	admin.Subject = make(map[string]map[string]float64)
	admin.Student = make(map[string]map[string]float64)

	*reply = "Init successful!"
	return nil
}

// Adds a new stutent to the map thing
func (admin *Admin) Add(args Util.Args, reply *string) error {
	admin.Subject[args.Subject][args.Name] = args.Grade
	admin.Student[args.Name][args.Subject] = args.Grade

	*reply = "New student added!"
	return nil
}

// Returns the student average grade of all its subjects
func (admin *Admin) StudentAverage(name string, reply *float64) error {
	*reply = 0

	for _, grade := range admin.Student[name] {
		*reply += grade
	}

	return nil
}

// Returns the subject average grade of all its students
func (admin *Admin) SubjectAverage(sub string, reply *float64) error {
	*reply = 0

	for _, grade := range admin.Subject[sub] {
		*reply += grade
	}

	return nil
}

// Returns the average of all students and subjects
func (admin *Admin) GeneralAverage(string, reply *float64) error {
	*reply = 0

	/**for _, student := range admin.Subject {
		for _, grade := range admin.Student[student] {
			avrg += grade
		}
	}**/

	return nil
}

package admin

type Admin struct {
	Subject map[string]map[string]float64
	Student map[string]map[string]float64
}

// Inits the subjects data maps
func (admin *Admin) Init() {
	admin.Subject = make(map[string]map[string]float64)
	admin.Student = make(map[string]map[string]float64)
}

// Adds a new stutent to the map thing
func (admin *Admin) Add(name string, sub string, grade float64) error {
	admin.Subject[sub][name] = grade
	admin.Student[name][sub] = grade

	return nil
}

// Returns the student average grade of all its subjects
func (admin *Admin) StudentAverage(name string) float64 {
	var avrg float64 = 0

	for _, grade := range admin.Student[name] {
		avrg += grade
	}

	return avrg
}

// Returns the subject average grade of all its students
func (admin *Admin) SubjectAverage(sub string) float64 {
	var avrg float64 = 0

	for _, grade := range admin.Subject[sub] {
		avrg += grade
	}

	return avrg
}

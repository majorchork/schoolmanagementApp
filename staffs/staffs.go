package staff

// HonoursRole is a variable that stores a list of names of honours role students

var HonoursRole []string

// Staff is a custom datatype of underlying type of struct

type Staff struct {
	Name         string
	StudentFiles []string
	StudentGpa   float64
	studentName  string
}

// AwardRecommendation is a method for a staff to recommend students for an honours role award, it takes the
// students Gpa then checks it against the threshold Gpa then returns a recommendation

func (i *Staff) AwardRecommendation(threshold float64) string {
	if i.StudentGpa >= threshold {
		return "Student is recommended for HonourRole Awards"
		HonoursRole = append(HonoursRole, i.studentName)
	}
	return "student needs to work harder to be eligible for honourRole"
}

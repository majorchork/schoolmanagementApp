package teachers

// Teachers is a custom datatype of underlying type of struct

type Teachers struct {
	Name          string
	CoursesTaught []string
	NoiseMakers   []string
}

// GradeStudents is a method for the teachers custom datatype
// the method enables a teacher to grade students if the course submitted by the student is a course taught by the teacher
// it iterates through the array and checks if the course submitted by the student is in the array of the courses taught
// by the teacher then returns a value

func (g *Teachers) GradeStudent(courseSubmitted string) string {
	//var gradeStatus = "Teacher does not teach this course and cannot grade it. "
	for _, course := range g.CoursesTaught {
		if course == courseSubmitted {
			return "student graded"
		}
	}
	return "Teacher does not teach this course and cannot grade it."
}

// DisciplineStudents is a method for the teachers custom datatype
// This method enables a teacher to discpline a student by sending the student to detention
// if the name of the student is on the list of noise makers
// This method iterates through the array of list of noise makers and compares names in the list to the students name

func (g *Teachers) DisciplineStudent(nameOfStudent string) string {
	for _, name := range g.NoiseMakers {
		if name == nameOfStudent {
			return name + " will be in detention for 3hrs after school today"
		}
	}
	return "You're Safe"
}

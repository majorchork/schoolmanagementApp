package teachers

type Teachers struct {
	Name          string
	CoursesTaught []string
	NoiseMakers   []string
}

func (g *Teachers) GradeStudent(courseSubmitted string) string {
	//var gradeStatus = "Teacher does not teach this course and cannot grade it. "
	for _, course := range g.CoursesTaught {
		if course == courseSubmitted {
			return "student graded"
		}
	}
	return "Teacher does not teach this course and cannot grade it."
}
func (g *Teachers) DisciplineStudent(nameOfStudent string) string {
	for _, name := range g.NoiseMakers {
		if name == nameOfStudent {
			return name + " will be in detention for 3hrs after school today"
		}
	}
	return "You're Safe"
}

package student

type Student struct {
	Name                string
	Age                 float64
	CreditStatus        bool
	FeesPaid            float64
	DepartmentalCourses []string
	Cgpa                float64
}

func (s *Student) TakeCourse(desiredCourse string) string {
	for _, courses := range s.DepartmentalCourses {
		if courses == desiredCourse {
			return "You can take this Course"
		}
	}
	return "Course is not a Departmental Course"
}

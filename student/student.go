package student

//type Course struct {
//	Title string
//	Grade float64
//}
type Student struct {
	Name                string
	Age                 float64
	CreditStatus        bool
	FeesPaid            float64
	DepartmentalCourses []string
	Cgpa                float64
}

//func (s *Student) CalcCGPA(desiredCourse string) string {
//	var total float64
//	for _, course := range s.DepartmentalCourses {
//		total += course.Grade
//	}
//	s.Cgpa = total
//}

func (s *Student) TakeCourse(desiredCourse string) string {
	for _, courses := range s.DepartmentalCourses {
		if courses == desiredCourse {
			return "You can take this Course"
		}
	}
	return "Course is not a Departmental Course"
}

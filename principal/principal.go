package principal

var SliceOfAdmittedStudents []string

type Principal struct {
	Name              string
	YearsOfExperience float64
	Qualifications    []string
}

func (P *Principal) Admits(age float64, applicantName string) string {
	if age > 10 {
		SliceOfAdmittedStudents = append(SliceOfAdmittedStudents, applicantName)
		return "Congratulations You have Gained Admission"
	}
	return "Not Admitted, Applicant is too Young for JSS1"
}
func (P *Principal) Expels(failedCourses float64) string {
	if failedCourses > 0 && failedCourses <= 2 {
		return "student is advised to put in more effort"
	} else if failedCourses > 2 && failedCourses <= 4 {
		return "student is advised to withdraw"
	} else if failedCourses > 5 {
		return "student expelled for poor grades"
	}
	return "good student keep it up"
}
func (p *Principal) Suspends(creditStatus bool) string {
	if creditStatus == true {
		return "student is suspended"
	}
	return "student does not owe and can continue taking classes"
}

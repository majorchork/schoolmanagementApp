package school

import (
	"week3/applicants"
	"week3/principal"
	Staff "week3/staffs"
	"week3/student"
	"week3/teachers"
)

type NonAcademicStaff struct{}
type Courses struct{}
type Classes struct{}

type School struct {
	Principal        principal.Principal
	Student          student.Student
	Staffs           Staff.Staff
	Teachers         teachers.Teachers
	NonAcademicStaff NonAcademicStaff
	Courses          Courses
	Classes          Classes
	Applicants       applicants.Applicants
}

var PrincipalInstance = principal.Principal{
	Name:              "aboy",
	YearsOfExperience: 32,
	Qualifications:    []string{"Phd", "Msc", "Bsc"},
}
var TeacherInstance = teachers.Teachers{
	Name:          "kunu",
	CoursesTaught: []string{"maths", "english", "civiceducation", "sexEd", "geography"},
	NoiseMakers:   []string{"carl", "pal", "val"},
}

var StudentInstance = student.Student{
	Name:                "Oreva Omoale",
	Age:                 70,
	CreditStatus:        false,
	FeesPaid:            2500000.00,
	DepartmentalCourses: []string{"science", "sport", "Algo101", "mathsLogic"},
	Cgpa:                0.15,
}
var Applicantsinstance = applicants.Applicants{
	Name:          "Portable Joseph",
	Age:           70,
	SkillStrength: 60,
}
var StaffInstance = Staff.Staff{
	Name: "azeez",
	StudentFiles: []string{
		"medicalhistory", "Addresss:12Avaway"},
	StudentGpa: 0,
}

func (z *School) HirePrincipal(qualification string) string {
	for _, value := range PrincipalInstance.Qualifications {
		if qualification == value {
			return "hired"
		}
	}
	return "Sorry, you do not met the qualification requirements"
}

//func (z *School) CheckStudent() string {
//	result := ""
//	//if PrincipalInstance.Admits() ==
//	return result
//}

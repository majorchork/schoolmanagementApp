package school

import (
	"strings"
	"week3/applicants"
	"week3/nonacademicstaffs"
	"week3/principal"
	staff "week3/staffs"
	"week3/student"
	"week3/teachers"
)

type School struct {
	Principal        principal.Principal
	Student          student.Student
	Staffs           staff.Staff
	Teachers         teachers.Teachers
	NonAcademicStaff nonacademicstaffs.NoneAcademicStaff
	Applicants       applicants.Applicants
}

var ChuksSchool = School{
	PrincipalInstance, StudentInstance, StaffInstance, TeacherInstance, NoneAcaademicStaffInstance, Applicantsinstance,
}

var SchoolInstance = School{
	Principal:        PrincipalInstance,
	Student:          StudentInstance,
	Staffs:           StaffInstance,
	Teachers:         TeacherInstance,
	NonAcademicStaff: NoneAcaademicStaffInstance,
	Applicants:       Applicantsinstance,
}

var PrincipalInstance = principal.Principal{
	Name:                    "aboy",
	YearsOfExperience:       32,
	SliceOfAdmittedStudents: []string{"Phd", "Msc", "Bsc"},
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
var StaffInstance = staff.Staff{
	Name: "azeez",
	StudentFiles: []string{
		"medicalhistory", "Addresss:12Avaway"},
	StudentGpa: 0,
}
var NoneAcaademicStaffInstance = nonacademicstaffs.NoneAcademicStaff{
	StaffName:       "Morris",
	Purview:         []string{"Admin Block", "Block 1", "Block 2"},
	StaffDepartment: "Maintenance",
}

//func (z *School) HirePrincipal(qualification string) string {
//	for _, value := range PrincipalInstance.Qualifications {
//		if qualification == value {
//			return "hired"
//		}
//	}
//	return "Sorry, you do not met the qualification requirements"
//}

//func (z *School) CheckStudent() string {
//	result := ""
//	//if PrincipalInstance.Admits() ==

//	return result
//}

var email string = "sochokwus@gmail.com"
var specialChar string = "#()&*!~"

func testEmail(email string) bool {
	if strings.Contains(email, "@") {
		return true
	}
	if strings.Contains(email, specialChar) {
		return false
	}
	return false
}

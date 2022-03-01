package main

import (
	"fmt"
	"week3/principal"
	"week3/teachers"
)

var prinstruct = principal.Principal{
	Name:              "John",
	YearsOfExperience: 32,
	Qualifications:    []string{"Phd", "Msc", "Bsc"},
}
var teacherStruct = teachers.Teachers{
	Name:          "kunu",
	CoursesTaught: []string{"maths", "english", "civiceducation", "sexEd", "geography"},
	NoiseMakers:   []string{"carl", "pal", "val"},
}

func main() {
	fmt.Println(prinstruct.Admits(9, "oboy"))
	fmt.Println(teacherStruct.GradeStudent("sexEd"))
}

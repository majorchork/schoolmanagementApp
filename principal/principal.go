package principal

// Principal is a custom datatype of underlying type of struct

type Principal struct {
	Name                    string
	YearsOfExperience       float64
	SliceOfAdmittedStudents []string
}

// Admit is a method for the Principal custom datatype
// the method enables a principal to admit a students if the age of the student
//greater than 10years old as is the school policy for admissions
// if the age of the students meets university policy then we can add the students name
//to the slice of  admitted students by appending

func (P *Principal) Admits(age float64, applicantName string) string {
	if age > 10 {
		P.SliceOfAdmittedStudents = append(P.SliceOfAdmittedStudents, applicantName)
		return "Congratulations You have Gained Admission"
	}
	return "Not Admitted, Applicant is too Young for JSS1"
}

// Expel is a method for the Principal custom datatype
// the method enables a principal to sanction or expel a student a students based on
//the number of failed courses the student has
//if the student has 5 failed courses the student is expelled
// if the student 3 or 4 failed courses the student is advised to withdraw
// if the student has 1 or 2 failed courses the student is advised to put in ore effort

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

// Suspend is a method for the Principal custom datatype
// this method enables a principal suspend a student based on their credit status,
// in essence to check if a students owe's or if a student is debt free
// if the students credit status is true then the student is suspended
// else the student is allowed to continue taking clases

func (P *Principal) Suspends(creditStatus bool) string {
	if creditStatus == true {
		return "student is suspended"
	}
	return "student does not owe and can continue taking classes"
}

// Suspend is a method for the Principal custom datatype
// This is the method a principal uses to validate a student's admission status
func (P Principal) ValidateStudent(studentName string) bool {
	for _, name := range P.SliceOfAdmittedStudents {
		if name == studentName {
			return true
		}
	}
	return false
}

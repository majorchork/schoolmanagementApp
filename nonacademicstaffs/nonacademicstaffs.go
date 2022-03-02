package nonacademicstaffs

// NoneAcademicStaff is a custom datatype of underlying type of struct
type NoneAcademicStaff struct {
	StaffName       string
	Purview         []string
	StaffDepartment string
}

// a staff can fix issues within a school if the issue falls within his purview
// this method receives issues from different areas in the school, then checks
// if the error falls within the purview of the staff in question, then the staff can  fix the issue

func (s *NoneAcademicStaff) FixIssues(issue, areaOfFault string) string {
	for _, location := range s.Purview {
		if areaOfFault == location {
			return issue + " can be fixed"
		}
	}
	return issue + " is out of my purview, please ask the staff in charge of your purview"
}

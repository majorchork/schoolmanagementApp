package nonacademicstaffs

import "testing"

func TestNoneAcademicStaff_FixIssues(t *testing.T) {
	var TestFixIssues = []struct {
		input  NoneAcademicStaff
		input1 string
		input2 string

		output string
	}{
		{NoneAcademicStaff{"Mark", []string{"Block 1", "Block 2", "Admin Block"}, "Maintenance"}, "faulty Socket", "Block 1", "faulty Socket can be fixed"},
		{NoneAcademicStaff{"Mark", []string{"Block 1", "Block 2", "Admin Block"}, "Maintenance"}, "faulty Socket", "Block 3", "faulty Socket is out of my purview, please ask the staff in charge of your purview"},
	}
	for _, value := range TestFixIssues {
		functionOutput := value.input.FixIssues(value.input1, value.input2)
		if functionOutput != value.output {
			t.Errorf("funtion expected %v but got %v", value.output, functionOutput)
		}
	}
}

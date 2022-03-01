package applicants

import "testing"

func TestApplicants_EntranceExam(t *testing.T) {
	var TestEntranceExam = []struct {
		input  Applicants
		input1 float64
		input2 float64

		output string
	}{
		{Applicants{"Sam", 14, 320}, 300, 10, "Applicant Successful"},
		{Applicants{"Rex", 16, 280}, 300, 10, "Applicant Rejected"},
	}
	for _, value := range TestEntranceExam {
		got := value.input.EntranceExam(value.input1, value.input2)
		if got != value.output {
			t.Errorf("expected output to be %v but got output %v", value.output, got)
		}
	}
}

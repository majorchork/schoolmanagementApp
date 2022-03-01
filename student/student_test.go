package student

import "testing"

func TestStudent_TakeCourse(t *testing.T) {
	var TestTakeCourse = []struct {
		input  Student
		input1 string

		output string
	}{
		{Student{"Sam", 14, false, 250000.02, []string{"Physics", "Maths", "Phe"}, 3.5}, "Phe", "You can take this Course"},
		{Student{"Rex", 16, false, 250000.03, []string{"Chemistry", "Biology", "Science"}, 2.5}, "Economics", "Course is not a Departmental Course"},
	}
	for _, value := range TestTakeCourse {
		got := value.input.TakeCourse(value.input1)
		if got != value.output {
			t.Errorf("expected output to be %v but got output %v", value.output, got)
		}
	}
}

package teachers

import "testing"

func TestTeachers_GradeStudent(t *testing.T) {
	testGrade := []struct {
		input  Teachers
		input1 string
		Output string
	}{
		{Teachers{"Kayode", []string{"maths", "english", "civiceducation"}, []string{"carl", "pal", "val"}}, "english", "student graded"},
		{Teachers{"Kayode", []string{"maths", "english", "civiceducation"}, []string{"carl", "pal", "val"}}, "biology", "Teacher does not teach this course and cannot grade it."},
	}
	for _, value := range testGrade {
		got := value.input.GradeStudent(value.input1)
		if got != value.Output {
			t.Errorf("expected %v, got %v", value.Output, got)
		}

	}
}
func TestTeachers_DisciplineStudent(t *testing.T) {
	testGrade := []struct {
		input  Teachers
		input1 string
		Output string
	}{
		{Teachers{"Kayode", []string{"maths", "english", "civiceducation"}, []string{"carl", "pal", "val"}}, "carl", "carl will be in detention for 3hrs after school today"},
		{Teachers{"Kayode", []string{"maths", "english", "civiceducation"}, []string{"carl", "daniel", "rob"}}, "david", "You're Safe"},
	}
	for _, value := range testGrade {
		got := value.input.DisciplineStudent(value.input1)
		if got != value.Output {
			t.Errorf("expected %v, got %v", value.Output, got)
		}

	}
}

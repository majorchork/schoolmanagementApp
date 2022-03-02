package principal

import "testing"

func TestPrincipal_Admits(t *testing.T) {
	var Testadmits = []struct {
		input  Principal
		input1 float64
		input2 string

		output string
	}{
		{Principal{"Bob", 34, []string{"Bob", "Tom", "Ross"}}, 15, "Oreva", "Congratulations You have Gained Admission"},
		{Principal{"Bob", 34, []string{"ROB", "Jim", "rose"}}, 9, "Debs", "Not Admitted, Applicant is too Young for JSS1"},
	}
	for _, value := range Testadmits {
		got := value.input.Admits(value.input1, value.input2)
		if got != value.output {
			t.Errorf("expected output to be %v but got output %v", value.output, got)
		}
	}
}
func TestPrincipal_Expels(t *testing.T) {
	var TestExpel = []struct {
		input  Principal
		input1 float64

		output string
	}{
		{Principal{"Bob", 34, []string{"Bob", "Tom", "Ross"}}, 15, "student expelled for poor grades"},
		{Principal{"Bob", 34, []string{"Bob", "Tom", "Ross"}}, 4, "student is advised to withdraw"},
		{Principal{"Bob", 34, []string{"Bob", "Tom", "Ross"}}, 1, "student is advised to put in more effort"},
		{Principal{"Bob", 34, []string{"Bob", "Tom", "Ross"}}, 0, "good student keep it up"},
	}
	for _, value := range TestExpel {
		got := value.input.Expels(value.input1)
		if got != value.output {
			t.Errorf("expected output to be %v but got output %v", value.output, got)
		}
	}
}
func TestPrincipal_Suspends(t *testing.T) {
	var TestSuspend = []struct {
		input  Principal
		input1 bool

		output string
	}{
		{Principal{"Bob", 34, []string{"Phd", "Bsc", "Msc"}}, true, "student is suspended"},
		{Principal{"Bob", 34, []string{"Phd", "Bsc", "Msc"}}, false, "student does not owe and can continue taking classes"},
		{Principal{"Bob", 34, []string{"Phd", "Bsc", "Msc"}}, false, "student does not owe and can continue taking classes"},
		{Principal{"Bob", 34, []string{"Phd", "Bsc", "Msc"}}, true, "student is suspended"},
	}
	for _, value := range TestSuspend {
		got := value.input.Suspends(value.input1)
		if got != value.output {
			t.Errorf("expected output to be %v but got output %v", value.output, got)
		}
	}
}
func TestPrincipal_ValidateStudent(t *testing.T) {
	var TestValidateStudent = []struct {
		input  Principal
		input1 string

		output bool
	}{
		{Principal{"Bob", 34, []string{"Bob", "Tom", "Ross"}}, "Ross", true},
		{Principal{"Bob", 34, []string{"ROB", "Jim", "rose"}}, "Debs", false},
	}
	for _, value := range TestValidateStudent {
		got := value.input.ValidateStudent(value.input1)
		if got != value.output {
			t.Errorf("expected output to be %v but got output %v", value.output, got)
		}
	}
}

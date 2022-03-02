package staff

import "testing"

func TestStaff_AwardRecommendation(t *testing.T) {
	var TestAwardRecognition = []struct {
		input  Staff
		input1 float64

		output string
	}{
		{Staff{"Sam", []string{"medicalhistory", "Address: 420 weedLane Sativile"}, 4.83, "Mike"}, 4.7, "Student is recommended for HonourRole Awards"},
		{Staff{"Alex", []string{"medicalhistory", "Address: 21 woodLane Oakland"}, 1.1, "Oreva"}, 4.7, "student needs to work harder to be eligible for honourRole"},
	}
	for _, value := range TestAwardRecognition {
		got := value.input.AwardRecommendation(value.input1)
		if got != value.output {
			t.Errorf("expected output to be %v but got output %v", value.output, got)
		}
	}
}

package Staff

var HonoursRole []string

type Staff struct {
	Name         string
	StudentFiles []string
	StudentGpa   float64
	studentName  string
}

func (i *Staff) AwardRecommendation(threshold float64) string {
	if i.StudentGpa >= threshold {
		return "Student is recommended for HonourRole Awards"
		HonoursRole = append(HonoursRole, i.studentName)
	}
	return "student needs to work harder to be eligible for honourRole"
}

package applicants

var ArrayOfSuccessfulApplicantts []string

type Applicants struct {
	Name          string
	Age           float64
	SkillStrength float64
}

func (a *Applicants) EntranceExam(cutOffMark, minimumAge float64) string {
	if a.Age >= minimumAge && a.SkillStrength >= cutOffMark {
		ArrayOfSuccessfulApplicantts = append(ArrayOfSuccessfulApplicantts, a.Name)
		return "Applicant Successful"
	}
	return "Applicant Rejected"
}

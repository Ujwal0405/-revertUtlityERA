package models

type LearnerReport struct {
	ProgramID    string    `json:"programId"`
	ProgramName  string    `json:"programName"`
	CourseID     string    `json:"courseId"`
	CourseName   string    `json:"courseName"`
	BatchID      string    `json:"batchId"`
	BatchName    string    `json:"batchName"`
	DistrictID   string    `json:"districtId"`
	DistrictName string    `json:"districtName"`
	Center       []Centers `json:"center"`
}
type Centers struct {
	CenterID   string     `json:"centerId"`
	CenterName string     `json:"centerName"`
	CenterCode string     `json:"centerCode"`
	Learner    []Learners `json:"learner"`
}
type Learners struct {
	LearnerID                    string  `json:"learnerId"`
	LearnerName                  string  `json:"learnerName"`
	LearnerCode                  string  `json:"learnerCode"`
	Marks                        float64 `json:"marks"`
	LoginID                      string  `json:"loginId"`
	Password                     string  `json:"password"`
	ExpectedAttendance           string  `json:"expectedAttendance"`
	ActualAttendance             string  `json:"actualAttendance"`
	PercentageOfCourseCompletion float64 `json:"percentageOfCourseCompletion"`
	TimeSpent                    int64   `json:"timeSpent"`
	SessionCompleted             float64 `json:"sessionCompleted"`
	InternalMarks                string  `json:"internalMarks"`
	TotalAssignments             string  `json:"totalAssignments"`
	UploadedAssignments          string  `json:"uploadedAssignments"`
}

package models

type MongoAssignmentSubmission struct {
	LearnerId                string              `bson:"learnerId"`
	ProgramId                string              `bson:"programId"`
	ECourseId                string              `bson:"eCourseId"`
	PatternID                string              `bson:"patternID"`
	SessionAggregationId     string              `bson:"sessionAggregationId"`
	ECoursePatternId         string              `bson:"eCoursePatternId"`
	ConfirmationMonth        int                 `bson:"confirmationMonth"`
	ConfirmationYear         int                 `bson:"confirmationYear"`
	AggregationId            string              `bson:"aggregationId"`
	AggregationName          string              `bson:"aggregationName"`
	AssignmentType           string              `bson:"assignmentType"`
	AssignmentId             string              `bson:"assignmentId"`
	AttemptNo                int                 `bson:"attemptNo"`
	AttemptStartDate         int64               `bson:"attemptStartDate"` // Assuming NumberLong is mapped to int64
	AttemptEndDate           int64               `bson:"attemptEndDate"`
	UploadedRelativeFilePath string              `bson:"uploadedRelativeFilePath"`
	MarksObtained            float64             `bson:"marksObtained"`
	LanguageId               string              `bson:"languageId"`
	AssignmentStatus         string              `bson:"assignmentStatus"`
	IsFileUploaded           bool                `bson:"isFileUploaded"`
	EvaluationDetails        []EvaluationDetails `bson:"evaluationDetails"`
}

// type EvaluationDetail struct {
// 	AssignmentCriteriaId   string  `bson:"assignmentCriteriaId"`
// 	AssignmentCriteriaName string  `bson:"assignmentCriteriaName"`
// 	EvaluatedMarks         float64 `bson:"evaluatedMarks"`
// 	CriteriaMarks          float64 `bson:"criteriaMarks"`
// }

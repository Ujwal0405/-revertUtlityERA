package models

type MongoLearnerDetails struct {
	LearnerId      string        `bson:"learnerId"`
	LearnerCode    string        `bson:"learnerCode"`
	Gender         string        `bson:"gender"`
	MobileNo       string        `bson:"mobileNo"`
	IDProof        string        `bson:"idProof"`
	CurrentAddress string        `bson:"currentAddress"`
	Disability     []interface{} `bson:"disability"` // Assuming disability is an array that can contain various types
	EmailId        string        `bson:"emailId"`
	LegalFullName  string        `bson:"legalFullName"`
	MobileNo2      string        `bson:"mobileNo2"`
	PhotoPath      string        `bson:"photoPath"`
}

type MongoLearnerCourceAllocation struct {
	LearnerId                     string               `bson:"learnerId"`
	LearnerCode                   string               `bson:"learnerCode"`
	LearnerCourseId               string               `bson:"learnerCourseId"`
	LearningCenterId              string               `bson:"learningCenterId"`
	LearningCenterCode            string               `bson:"learningCenterCode"`
	LearningCenterName            string               `bson:"learningCenterName"`
	ConfirmationMonth             int                  `bson:"confirmationMonth"`
	ConfirmationYear              int                  `bson:"confirmationYear"`
	BatchName                     string               `bson:"batchName"`
	ProgramId                     string               `bson:"programId"`
	ProgramName                   string               `bson:"programName"`
	ECourseId                     string               `bson:"eCourseId"`
	ECourseName                   string               `bson:"eCourseName"`
	ECourseVersion                string               `bson:"eCourseVersion"`
	ECoursePatternId              string               `bson:"eCoursePatternId"`
	DateOfTransfer                int64                `bson:"dateOfTransfer"`
	ECourseImage                  string               `bson:"eCourseImage"`
	LearningEndDate               int64                `bson:"learningEndDate"`
	LearningStartDate             int64                `bson:"learningStartDate"`
	LearningLimit                 int                  `bson:"learningLimit"`
	PerDaySessionLimit            int                  `bson:"perDaySessionLimit"`
	LearnerLifeCycleEndDate       int64                `bson:"learnerLifeCycleEndDate"`
	LearnerCoursePrefLanguageCode string               `bson:"learnerCoursePrefLanguageCode"`
	LearningBlock                 MongoLearningBlock   `bson:"learningBlock"`
	LearningDetails               MongoLearningDetails `bson:"learningDetails"`
}

type MongoLearningScoreDetails struct {
	PassingHeadCode        string            `bson:"passingHeadCode"`
	PassingHeadName        string            `bson:"passingHeadName"`
	PassingHeadDisplayName string            `bson:"passingHeadDisplayName"`
	Score                  float64           `bson:"score,omitempty"`
	ScoreDetails           MongoScoreDetails `bson:"scoreDetails,omitempty"`
	MaxMarks               float64           `bson:"maxMarks"`
}

// ScoreDetails
type MongoScoreDetails struct {
	TotalScore      float64 `bson:"totalScore,omitempty"`
	AdditionalMarks float64 `bson:"additionalScore,omitempty"`
}

type MongoLearningDetails struct {
	CurrentSession                 string                      `bson:"currentSession"`
	CurrentSessionId               string                      `bson:"currentSessionId"`
	NextSequentialAggregationId    string                      `bson:"nextSequentialAggregationId"`
	NextSequentialSessionId        string                      `bson:"nextSequentialSessionId"`
	LearningScoreDetails           []MongoLearningScoreDetails `bson:"learningScoreDetails"`
	LearningStartedOn              int64                       `bson:"learningStartedOn"`
	CurrentNodeAggregationId       string                      `bson:"currentNodeAggregationId"`
	CompletedAggregationCount      int                         `bson:"completedAggregationCount"`
	CompletedSessionCount          int                         `bson:"completedSessionCount"`
	CompletionPercentage           float64                     `bson:"completionPercentage"`
	CurrentSequentialAggregationId string                      `bson:"currentSequentialAggregationId"`
	FirstLogin                     int64                       `bson:"firstLogin"`
	LastUpdateDateTime             int64                       `bson:"lastUpdateDateTime"`
	LatestLogin                    int64                       `bson:"latestLogin"`
	SolvedAssignmentCount          int                         `bson:"solvedAssignmentCount"`
	SolvedChallengeCount           int                         `bson:"solvedChallengeCount"`
	SolvedDIYCount                 int                         `bson:"solvedDIYCount"`
	TimeSpent                      int                         `bson:"timeSpent"`
	VisitedTime                    int64                       `bson:"visitedTime"`
}

type MongoLearningBlock struct {
	HasBlocked  bool   `bson:"hasBlocked"`
	BlockedDate string `bson:"blockedDate,omitempty"`
	BlockDates  int64  `bson:"blockDates"`
}

type MongoSessionCompletionDetails struct {
	AggregationId            string                     `bson:"aggregationId"`
	ConfirmationMonth        int                        `bson:"confirmationMonth"`
	ConfirmationYear         int                        `bson:"confirmationYear"`
	ECourseId                string                     `bson:"eCourseId"`
	LearnerId                string                     `bson:"learnerId"`
	ProgramId                string                     `bson:"programId"`
	AggregationDuration      int                        `bson:"aggregationDuration"`
	CompletedOn              int64                      `bson:"completedOn"`
	ECoursePatternId         string                     `bson:"eCoursePatternId"`
	IsClassRoomAttempted     bool                       `bson:"isClassRoomAttempted"`
	IsCompleted              bool                       `bson:"isCompleted"`
	IsSysEntry               bool                       `bson:"isSysEntry"`
	LanguageId               string                     `bson:"languageId"`
	LastUpdateDateTime       int64                      `bson:"lastUpdateDateTime"`
	PassingHeadWiseMarks     []MongoPassingHeadWiseMark `bson:"passingHeadWiseMarks"`
	SessionStatus            string                     `bson:"sessionStatus"`
	TotalAggregationDuration int                        `bson:"totalAggregationDuration"`
	TotalClassroomDuration   int                        `bson:"totalClassroomDuration"`
	SessionMarks             float64                    `bson:"sessionMarks"`
}

type MongoPassingHeadWiseMark struct {
	PassingHeadCode string            `bson:"passingHeadCode"`
	Score           float64           `bson:"score"`
	ScoreDetails    MongoScoreDetails `bson:"scoreDetails"`
}

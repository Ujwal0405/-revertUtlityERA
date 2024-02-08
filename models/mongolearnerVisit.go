package models

type MongoLearnerVisit struct {
	LearnerID         string               `json:"learnerId,omitempty" bson:"learnerId,omitempty"`
	LearningCenterID  string               `json:"learningCenterId,omitempty" bson:"learningCenterId,omitempty"`
	ProgramID         string               `json:"programId,omitempty" bson:"programId,omitempty"`
	ECourseID         string               `json:"eCourseId,omitempty" bson:"eCourseId,omitempty"`
	ECoursePatternID  string               `json:"eCoursePatternId,omitempty" bson:"eCoursePatternId,omitempty"`
	LanguageID        string               `json:"languageId,omitempty" bson:"languageId,omitempty"`
	ConfirmationMonth int                  `json:"confirmationMonth,omitempty" bson:"confirmationMonth,omitempty"`
	ConfirmationYear  int                  `json:"confirmationYear,omitempty" bson:"confirmationYear,omitempty"`
	SessionID         string               `json:"sessionId,omitempty" bson:"sessionId,omitempty"`
	AggregationID     string               `json:"aggrId,omitempty" bson:"aggrId,omitempty"`
	AggregationTime   int                  `json:"aggrTm,omitempty" bson:"aggrTm,omitempty"`
	VisitedCount      int                  `json:"vCnt,omitempty" bson:"vCnt,omitempty"`
	IsVisited         bool                 `json:"isV,omitempty" bson:"isV,omitempty"`
	IsFirstAttempt    bool                 `json:"isFA,omitempty" bson:"isFA,omitempty"`
	TotalSteps        int                  `json:"tSteps,omitempty" bson:"tSteps,omitempty"`
	CorrectSteps      int                  `json:"cSteps,omitempty" bson:"cSteps,omitempty"`
	ContentResponse   MongoContentResponse `json:"cr,omitempty" bson:"cr,omitempty"`
	Marks             float64              `json:"mks,omitempty" bson:"mks,omitempty"`
	PassingHeadCode   string               `json:"phc,omitempty" bson:"phc,omitempty"`
	CreatedDate       int64                `json:"cDt,omitempty" bson:"cDt,omitempty"`
	ModifiedDate      int64                `json:"mDt,omitempty" bson:"mDt,omitempty"`
	UploadedDate      int64                `json:"ulDt,omitempty" bson:"ulDt,omitempty"`
	IsSystemEntry     bool                 `json:"isSysEnt,omitempty" bson:"isSysEnt,omitempty"`
}

type MongoContentResponse struct {
	Response MongoResponse `bson:"response,omitempty"`
	Marks    string        `bson:"marks,omitempty"`
	MaxMarks string        `bson:"maxMarks,omitempty"`
}

type MongoResponse struct {
	LearnerID               string `bson:"learnerId,omitempty"`
	ProgramID               string `bson:"programId,omitempty"`
	ECourseID               string `bson:"eCourseId,omitempty"`
	ECoursePatternID        string `bson:"eCoursePatternId,omitempty"`
	LanguageID              string `bson:"languageId,omitempty"`
	ExamID                  string `bson:"examId,omitempty"`
	ExamType                string `bson:"examType,omitempty"`
	SessionAggregationId    string `bson:"sessionAggregationId,omitempty"`
	AggregationID           string `bson:"aggregationId,omitempty"`
	ExamCompleted           bool   `bson:"examCompleted,omitempty"`
	QuestionsCount          int    `bson:"questionsCount,omitempty"`
	AttemptedQuestionsCount int    `bson:"attemptedQuestionsCount,omitempty"`
	CorrectAnswersCount     int    `bson:"correctAnswersCount,omitempty"`
	WrongAnswersCount       int    `bson:"wrongAnswersCount,omitempty"`
	TotalMarks              int    `bson:"totalMarks,omitempty"`
	ObtainedMarks           int    `bson:"obtainedMarks,omitempty"`
	TotalPercentage         int    `bson:"totalPercentage,omitempty"`
	ExamAttemptNumber       int    `bson:"examAttemptNumber,omitempty"`
}

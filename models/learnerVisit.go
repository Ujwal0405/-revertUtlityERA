package models

import "time"

// LearnerVisit represents learner visit details
type LearnerVisit struct {
	LearnerID        string         `json:"learnerId,omitempty" bson:"learnerId,omitempty"`
	LearningCenterID string         `json:"learningCenterId,omitempty" bson:"learningCenterId,omitempty"`
	ProgramID        string         `json:"programId,omitempty" bson:"programId,omitempty"`
	ECourseID        string         `json:"eCourseId,omitempty" bson:"eCourseId,omitempty"`
	ECoursePatternID string         `json:"eCoursePatternId,omitempty" bson:"eCoursePatternId,omitempty"`
	LanguageID       string         `json:"languageId,omitempty" bson:"languageId,omitempty"`
	BatchID          string         `json:"batchId,omitempty" bson:"batchId,omitempty"`
	SessionID        string         `json:"sessionId,omitempty" bson:"sessionId,omitempty"`
	VisitsDetail     []VisitsDetail `json:"vd,omitempty" bson:"vd,omitempty"`
}

// VisitsDetail represents visit details
type VisitsDetail struct {
	ScoLanguageID         string                `json:"slId" bson:"slId"`
	AggregationID         string                `json:"aggrId" bson:"aggrId"`
	AggregationTime       int                   `json:"aggrTm" bson:"aggrTm"`
	VisitedCount          int                   `json:"vCnt" bson:"vCnt"`
	IsVisited             bool                  `json:"isV" bson:"isV"`
	IsFirstAttempt        bool                  `json:"isFA" bson:"isFA"`
	TotalSteps            int                   `json:"ts" bson:"ts"`
	CorrectSteps          int                   `json:"cs" bson:"cs"`
	ContentResponse       ContentResponse       `json:"cr" bson:"cr"`
	Marks                 float64               `json:"mks" bson:"mks"`
	PassingHeadCode       string                `json:"phc" bson:"phc"`
	IsUploaded            bool                  `json:"isUL" bson:"isUL"`
	CreatedDate           *time.Time            `json:"cDt" bson:"cDt"`
	ModifiedDate          *time.Time            `json:"mDt" bson:"mDt"`
	UploadedDate          *time.Time            `json:"ulDt" bson:"ulDt"`
	IsExpired             bool                  `json:"isExp" bson:"isExp"`
	IsSystemEntry         bool                  `json:"ise" bson:"ise"`
	FeedbackObj           FeedbackObj           `json:"feedbackObj" bson:"feedbackObj"`
	ElectiveCourseDetails ElectiveCourseDetails `json:"ecd,omitempty" bson:"ecd,omitempty"`
	ClassroomDetails      ClassroomDetails      `json:"cd,omitempty" bson:"cd,omitempty"`
	ClassroomAttemptCount ClassroomAttemptCount `json:"clsdtl" bson:"clsdtl"`
}

// ClassroomAttemptCount represents classroom attempt count
type ClassroomAttemptCount struct {
	AttemptedCnt int        `json:"attmcnt,omitempty" bson:"attmcnt,omitempty"`
	CorrectCnt   int        `json:"corrtcnt,omitempty" bson:"corrtcnt,omitempty"`
	InCorrectCnt int        `json:"incrrtcnt,omitempty" bson:"incrrtcnt,omitempty"`
	CreatedDate  *time.Time `json:"cDt,omitempty" bson:"cDt,omitempty"`
	ModifiedDate *time.Time `json:"mDt,omitempty" bson:"mDt,omitempty"`
}

// ClassroomDetails represents classroom details
type ClassroomDetails struct {
	TotalAttempts      int        `json:"tA,omitempty" bson:"tA,omitempty"`
	CorrectAnswerCount int        `json:"cAC,omitempty" bson:"cAC,omitempty"`
	WrongAnswerCount   int        `json:"wAC,omitempty" bson:"wAC,omitempty"`
	CreatedDate        *time.Time `json:"cDt,omitempty" bson:"cDt,omitempty"`
	ModifiedDate       *time.Time `json:"mDt,omitempty" bson:"mDt,omitempty"`
}

// ContentResponse represents content response
type ContentResponse struct {
	Response         Response         `json:"response,omitempty" bson:"response,omitempty"`
	Marks            string           `json:"marks,omitempty" bson:"marks,omitempty"`
	MaxMarks         string           `json:"maxMarks,omitempty" bson:"maxMarks,omitempty"`
	MaxFacesCount    int              `json:"mfct,omitempty" bson:"mfct,omitempty"`
	FacesData        []FacesData      `json:"fcdt,omitempty" bson:"fcdt,omitempty"`
	TrackerData      []TrackerData    `json:"trdt,omitempty" bson:"trdt,omitempty"`
	ClassroomDetails ClassroomDetails `json:"cd" bson:"cd"`
}

// FacesData represents faces data
type FacesData struct {
	ImageNumber int `json:"imageNumber,omitempty" bson:"imageNumber,omitempty"`
	FacesCount  int `json:"facesCount,omitempty" bson:"facesCount,omitempty"`
}

// TrackerData represents tracker data
type TrackerData struct {
	ActionNumber int    `json:"ActionNumber,omitempty" bson:"ActionNumber,omitempty"`
	Description  string `json:"Description,omitempty" bson:"Description,omitempty"`
	Action       string `json:"Action,omitempty" bson:"Action,omitempty"`
	ProgramName  string `json:"ProgramName,omitempty" bson:"ProgramName,omitempty"`
	Title        string `json:"Title,omitempty" bson:"Title,omitempty"`
	KBInput      string `json:"KBInput,omitempty" bson:"KBInput,omitempty"`
	Date         string `json:"Date,omitempty" bson:"Date,omitempty"`
}

// ElectiveCourseDetails represents elective course details
type ElectiveCourseDetails struct {
	ProgramID         string     `json:"pId,omitempty" bson:"pId,omitempty"`
	ECourseID         string     `json:"ecId,omitempty" bson:"ecId,omitempty"`
	ECoursePatternID  string     `json:"ecpId,omitempty" bson:"ecpId,omitempty"`
	BatchID           string     `json:"bId,omitempty" bson:"bId,omitempty"`
	IsCompleted       bool       `json:"iC,omitempty" bson:"iC,omitempty"`
	CourseStatus      string     `json:"cs,omitempty" bson:"cs,omitempty"`
	CourseCompletedOn *time.Time `json:"cCmpOn,omitempty" bson:"cCmpOn,omitempty"`
	Marks             float64    `json:"mks,omitempty" bson:"mks,omitempty"`
}

// FeedbackObj represents feedback object
type FeedbackObj struct {
	Rating         int        `json:"rating,omitempty" bson:"rating,omitempty"`
	FeedbackOption []int      `json:"feedbackOption,omitempty" bson:"feedbackOption,omitempty"`
	CommentText    string     `json:"commentText,omitempty" bson:"commentText,omitempty"`
	ContentOption  []int      `json:"contentOption,omitempty" bson:"contentOption,omitempty"`
	CreatedDate    *time.Time `json:"createdDate,omitempty" bson:"createdDate,omitempty"`
	ModifiedDate   *time.Time `json:"modifiedDate,omitempty" bson:"modifiedDate,omitempty"`
	FeedbackCount  int        `json:"feedbackCount,omitempty" bson:"feedbackCount,omitempty"`
}

type Response struct {
	LearnerID               string `json:"learnerId,omitempty" bson:"learnerId,omitempty"`
	ProgramID               string `json:"programId,omitempty" bson:"programId,omitempty"`
	ECourseID               string `json:"eCourseId,omitempty" bson:"eCourseId,omitempty"`
	ECoursePatternID        string `json:"eCoursePatternId,omitempty" bson:"eCoursePatternId,omitempty"`
	LanguageID              string `json:"languageId,omitempty" bson:"languageId,omitempty"`
	ExamID                  string `json:"examId,omitempty" bson:"examId,omitempty"`
	ExamType                string `json:"examType,omitempty" bson:"examType,omitempty"`
	SessionAggregationId    string `json:"sessionAggregationId,omitempty" bson:"sessionAggregationId,omitempty"`
	AggregationID           string `json:"aggregationId,omitempty" bson:"aggregationId,omitempty"`
	ExamCompleted           bool   `json:"examCompleted,omitempty" bson:"examCompleted,omitempty"`
	QuestionsCount          int    `json:"questionsCount,omitempty" bson:"questionsCount,omitempty"`
	AttemptedQuestionsCount int    `json:"attemptedQuestionsCount,omitempty" bson:"attemptedQuestionsCount,omitempty"`
	CorrectAnswersCount     int    `json:"correctAnswersCount,omitempty" bson:"correctAnswersCount,omitempty"`
	WrongAnswersCount       int    `json:"wrongAnswersCount,omitempty" bson:"wrongAnswersCount,omitempty"`
	TotalMarks              int    `json:"totalMarks,omitempty" bson:"totalMarks,omitempty"`
	ObtainedMarks           int    `json:"obtainedMarks,omitempty" bson:"obtainedMarks,omitempty"`
	TotalPercentage         int    `json:"totalPercentage,omitempty" bson:"totalPercentage,omitempty"`
	ExamAttemptNumber       int    `json:"examAttemptNumber,omitempty" bson:"examAttemptNumber,omitempty"`
}

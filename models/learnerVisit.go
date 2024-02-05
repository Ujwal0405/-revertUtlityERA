package models

//go:generate ffjson $GOFILE

import "time"

//LearnerVisit LearnerVisit
type LearnerVisit struct {
	LearnerID        string         `json:"learnerId,omitempty"`
	LearningCenterID string         `json:"learningCenterId,omitempty"`
	ProgramID        string         `json:"programId,omitempty"`
	ECourseID        string         `json:"eCourseId,omitempty"`
	ECoursePatternID string         `json:"eCoursePatternId,omitempty"`
	LanguageID       string         `json:"languageId,omitempty"`
	BatchID          string         `json:"batchId,omitempty"`
	SessionID        string         `json:"sessionId,omitempty"`
	VisitsDetail     []VisitsDetail `json:"vd,omitempty"`
}

//VisitsDetail vd
type VisitsDetail struct {
	ScoLanguageID   string          `json:"slId"`
	AggregationID   string          `json:"aggrId"`
	AggregationTime int             `json:"aggrTm"`
	VisitedCount    int             `json:"vCnt"`
	IsVisited       bool            `json:"isV"`
	IsFirstAttempt  bool            `json:"isFA"`
	TotalSteps      int             `json:"ts"`
	CorrectSteps    int             `json:"cs"`
	ContentResponse ContentResponse `json:"cr"`
	Marks           float64         `json:"mks"`
	PassingHeadCode string          `json:"phc"`
	IsUploaded      bool            `json:"isUL"`
	CreatedDate     *time.Time      `json:"cDt"`
	ModifiedDate    *time.Time      `json:"mDt"`
	UploadedDate    *time.Time      `json:"ulDt"`
	IsExpired       bool            `json:"isExp"`
	IsSystemEntry   bool            `json:"ise"`
	FeedbackObj     FeedbackObj     `json:"feedbackObj"`
	// Elective Course related attributes
	ElectiveCourseDetails ElectiveCourseDetails `json:"ecd,omitempty"`
	ClassroomDetails      ClassroomDetails      `json:"cd,omitempty"`
	ClassroomAttemptCount ClassroomAttemptCount `json:"clsdtl"`
}

//ClassroomAttemptCount ClassroomAttemptCount
type ClassroomAttemptCount struct {
	AttemptedCnt int        `json:"attmcnt,omitempty"`
	CorrectCnt   int        `json:"corrtcnt,omitempty"`
	InCorrectCnt int        `json:"incrrtcnt,omitempty"`
	CreatedDate  *time.Time `json:"cDt,omitempty"`
	ModifiedDate *time.Time `json:"mDt,omitempty"`
}

//ClassroomDetails cd
type ClassroomDetails struct {
	TotalAttempts      int        `json:"tA,omitempty"`
	CorrectAnswerCount int        `json:"cAC,omitempty"`
	WrongAnswerCount   int        `json:"wAC,omitempty"`
	CreatedDate        *time.Time `json:"cDt,omitempty"`
	ModifiedDate       *time.Time `json:"mDt,omitempty"`
}

//ContentResponse ContentResponse
type ContentResponse struct {
	Response         interface{}      `json:"response,omitempty"`
	Marks            string           `json:"marks,omitempty"`
	MaxMarks         string           `json:"maxMarks,omitempty"`
	MaxFacesCount    int              `json:"mfct,omitempty"`
	FacesData        []FacesData      `json:"fcdt,omitempty"`
	TrackerData      []TrackerData    `json:"trdt,omitempty"`
	ClassroomDetails ClassroomDetails `json:"cd"`
}
type FacesData struct {
	ImageNumber int `json:"imageNumber,omitempty"`
	FacesCount  int `json:"facesCount,omitempty"`
}

type TrackerData struct {
	ActionNumber int    `json:"ActionNumber,omitempty"`
	Description  string `json:"Description,omitempty"`
	Action       string `json:"Action,omitempty"`
	ProgramName  string `json:"ProgramName,omitempty"`
	Title        string `json:"Title,omitempty"`
	KBInput      string `json:"KBInput,omitempty"`
	Date         string `json:"Date,omitempty"`
}
type ElectiveCourseDetails struct {
	ProgramID        string `json:"pId,omitempty"`
	ECourseID        string `json:"ecId,omitempty"`
	ECoursePatternID string `json:"ecpId,omitempty"`
	BatchID          string `json:"bId,omitempty"`

	//transactional fields not for saving into file
	IsCompleted       bool       `json:"iC,omitempty"`
	CourseStatus      string     `json:"cs,omitempty"`
	CourseCompletedOn *time.Time `json:"cCmpOn,omitempty"`
	Marks             float64    `json:"mks,omitempty"`
}

type FeedbackObj struct {
	Rating         int        `json:"rating,omitempty"`
	FeedbackOption []int      `json:"feedbackOption"`
	CommentText    string     `json:"commentText,omitempty"`
	ContentOption  []int      `json:"contentOption"`
	CreatedDate    *time.Time `json:"createdDate,omitempty"`
	ModifiedDate   *time.Time `json:"modifiedDate,omitempty"`
	FeedbackCount  int        `json:"feedbackCount,omitempty"`
}

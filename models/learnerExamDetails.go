package models

//go:generate ffjson $GOFILE

import "time"

//LearnerExamDetails LearnerExamDetails
type LearnerExamDetails struct {
	ProgramID            string `json:"programId"`
	ECourseID            string `json:"eCourseId"`
	ECoursePatternID     string `json:"eCoursePatternId"`
	LearnerID            string `json:"learnerId"`
	SessionAggregationID string `json:"sessionAggregationId"`
	// NOTE :  Save Practise Attempt
	MockExamAttemptDetails []ExamAttemptDetails `json:"mockExamAttemptDetails"`
	ExamAttemptDetails     []ExamAttemptDetails `json:"examAttemptDetails"`

	ClassRoomAttemptDetails []ClassRoomAttemptDetail `json:"classRoomAttemptDetails"`

	// Final Exam Related changes
	ExamID                       string `json:"examId"`
	ExamPatternID                string `json:"examPatternId"`
	ExamType                     string `json:"examType"`
	AggregationID                string `json:"aggregationId"`
	LanguageID                   string `json:"languageID"`
	CurrentExamAttemptNumber     int    `json:"currentExamAttemptNumber"`
	MockCurrentExamAttemptNumber int    `json:"mockcurrentExamAttemptNumber"`
	// NOTE :  Save Practise Attempt
	MockSectionsAttemptDetails []SectionsAttemptDetail `json:"mockSectionsAttemptDetails"`
	SectionsAttemptDetails     []SectionsAttemptDetail `json:"sectionsAttemptDetails"`
}

type SectionsAttemptDetail struct {
	AttemptNumber  int             `json:"attemptNumber"`
	SectionDetails []SectionDetail `json:"sectionDetails"`
}

//ClassRoomAttemptDetail ClassRoomAttemptDetail
type ClassRoomAttemptDetail struct {
	ClassRoomPin                  string                          `json:"classRoomPin"`
	ClassRoomName                 string                          `json:"classRoomName"`
	ClassRoomStartedOn            time.Time                       `json:"classRoomStartedOn"`
	ClassRoomClosedOn             time.Time                       `json:"classRoomClosedOn"`
	ExamID                        string                          `json:"examId"`
	AggregationID                 string                          `json:"aggregationId"`
	AttemptNumber                 int                             `json:"attemptNumber"`
	ClickerQuestionAttemptDetails []ClickerQuestionAttemptDetails `json:"clickerQuestionAttemptDetails"`
	IsClassRoomCoompleted         bool                            `json:"isClassRoomCompleted"`
	TotalMarksObtained            float64                         `json:"totalMarksObtained"`
	TotalMarks                    float64                         `json:"totalMarks"`
	TotalItems                    int                             `json:"totalItems"`
	LanguageID                    string                          `json:"languageID"`
	Percentage                    float64                         `json:"percentage"`
	AttemptedItemCount            int                             `json:"attemptedItemCount"`
	CorrectItemCount              int                             `json:"correctItemCount"`
	WrongItemCount                int                             `json:"wrongItemCount"`
}

//ExamAttemptDetails ExamAttemptDetails
type ExamAttemptDetails struct {
	ExamID                       string                   `json:"examId"`
	AggregationID                string                   `json:"aggregationId"`
	AttemptNumber                int                      `json:"attemptNumber"`
	QuestionAttemptDetails       []QuestionAttemptDetails `json:"questionAttemptDetails"`
	IsExamCompleted              bool                     `json:"isExamCompleted"`
	TotalCorrectChars            int                      `json:"totalCorrectChars"`
	Accuracy                     string                   `json:"accuracy"`
	NetWordsPerMinute            int                      `json:"netWordsPerMinute"`
	IsSectionTimed               bool                     `json:"isSectionTimed"`
	DurationInSec                int                      `json:"durationInSec"`
	ElapsedTime                  int                      `json:"elapsedTime"`
	TotalMarksObtained           float64                  `json:"totalMarksObtained"`
	TotalTalentZoneMarksObtained float64                  `json:"totalTalentZoneMarksObtained"`
	ExamStartedOn                time.Time                `json:"examStartedOn"`
	ExamEndedOn                  time.Time                `json:"examEndedOn"`
	TotalMarks                   float64                  `json:"totalMarks"`
	TotalTalentZoneMarks         float64                  `json:"totalTalentZoneMarks"`
	TotalItems                   int                      `json:"totalItems"`
	TotalTalentZoneItems         int                      `json:"totalTalentZoneItems"`
	LanguageID                   string                   `json:"languageID"`
	Percentage                   float64                  `json:"percentage"`
	AttemptedItemCount           int                      `json:"attemptedItemCount"`
	AttemptedTalentZoneItemCount int                      `json:"attemptedTalentZoneItemCount"`
	CorrectItemCount             int                      `json:"correctItemCount"`
	CorrectTalentZoneItemCount   int                      `json:"correctTalentZoneItemCount"`
	WrongItemCount               int                      `json:"wrongItemCount"`
	WrongTalentZoneItemCount     int                      `json:"wrongTalentZoneItemCount"`
	ExamCurrentStatus            string                   `json:"examCurrentStatus"`
	ExamSubmittedOn              *time.Time               `json:"examSubmittedOn,omitempty"`
	ExamEvaluatedOn              *time.Time               `json:"examEvaluatedOn,omitempty"`
	IsQuestionPaperGenerated     bool                     `json:"isQuestionPaperGenerated"`
	EvaluationID                 string                   `json:"evaluationId"`
	EvaluationServerStatus       string                   `json:"evaluationServerStatus"`
	EvaluationServerError        string                   `json:"evaluationServerError"`
	// NOTE : CET
	CloudEvidenceUploadStatus       string  `json:"cloudEvidenceUploadStatus"`
	LocalEvidenceUploadStatus       string  `json:"localEvidenceUploadStatus"`
	EvidenceUploadedPath            string  `json:"evidenceUploadedPath"`
	MinimumPassingMarks             float64 `json:"minimumPassingMarks"`
	EvidenceUploadAcknowledgementID string  `json:"evidenceUploadAcknowledgementID"`
	ExamEndCause                    string  `json:"examEndCause"`
	//LocalTOC                        LocalTOC `json:"localTOC"`
	// Logs for CET Exam
	// ExamLogs				[]LogDetails				`json:"examLogs"`

}

//QuestionAttemptDetails QuestionAttemptDetails
type QuestionAttemptDetails struct {
	QuestionID              string          `json:"questionID"`
	Group                   int             `json:"group,omitempty"`
	Ordinality              int             `json:"ordinality"`
	SectionID               string          `json:"sectionId"`
	SectionName             string          `json:"sectionName"`
	QuestionType            string          `json:"questionType"`
	LanguageID              string          `json:"languageID"`
	AttemptDate             time.Time       `json:"attemptDate"`
	QuestionSequenceNo      string          `json:"questionSequenceNo"`
	WrongCharsCount         int             `json:"wrongCharsCount"`
	TypedText               string          `json:"typedText"`
	CorrectWordsCount       string          `json:"correctWordsCount"`
	WordsPerMinute          int             `json:"wordsPerMinute"`
	TimeElapsed             int             `json:"timeElapsed"`
	TypingAccuracy          float64         `json:"typingAccuracy"`
	MarksObtained           float64         `json:"marksObtained"`
	SubmittedAnswer         SubmittedAnswer `json:"submittedAnswer"`
	FeedbackReceived        string          `json:"feedbackReceived"`
	QuestionStatus          string          `json:"questionStatus"`
	QuestionViewedCount     int             `json:"questionViewedCount"`
	IsQuestionPublished     bool            `json:"isQuestionPublished"`
	Tag                     string          `json:"tag"`
	QuestionDifficultyLevel string          `json:"questionDifficultyLevel"`
	TotalMarks              float64         `json:"totalMarks"`
	DurationInSec           int             `json:"durationInSec"`

	// Added for TYP
	GrossWPM         int    `json:"grossWPM"`
	TotalDeleteCount int    `json:"totalDeleteCount"`
	WrongWordCount   int    `json:"wrongWordCount"`
	TotalKeyStrokes  int    `json:"totalKeyStrokes"`
	NetWPM           int    `json:"netWPM"`
	FeedbackReason   string `json:"feedbackReason"`
	FeedbackComment  string `json:"feedbackComment"`

	// Added for Question Simulator
	SimulatorCurrentStepID    string `json:"simulatorCurrentStepId"`
	SimulatorCurrentControlID string `json:"simulatorCurrentControlId"`

	// Added for Question feedback functionality
	SubmitCount int `json:"submitCount"`
}

//ClikerQuestionAttemptDetails ClikerQuestionAttemptDetails
type ClickerQuestionAttemptDetails struct {
	AggregationID        string          `json:"aggregationId"`
	SessionAggregationID string          `json:"sessionAggregationId"`
	QuestionID           string          `json:"questionID"`
	Ordinality           int             `json:"ordinality"`
	SectionName          string          `json:"sectionName"`
	QuestionType         string          `json:"questionType"`
	LanguageID           string          `json:"languageID"`
	AttemptDate          time.Time       `json:"attemptDate"`
	TimeElapsed          int             `json:"timeElapsed"`
	MarksObtained        float64         `json:"marksObtained"`
	SubmittedAnswer      SubmittedAnswer `json:"submittedAnswer"`
	FeedbackReceived     string          `json:"feedbackReceived"`
	QuestionStatus       string          `json:"questionStatus"`
}

//SubmittedAnswer SubmittedAnswer
type SubmittedAnswer struct {
	AnswerDetails string `json:"answerDetails"`
	IsCorrect     bool   `json:"isCorrect"`
	TotalSeconds  int    `json:"totalSeconds"`
	// Save submitted answer languageId
	SubmittedAnswerLanguageID string `json:"submittedAnswerLanguageId"`
}

// Final Exam related changes

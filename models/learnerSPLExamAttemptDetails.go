package models

import (
	"time"
)

type LearnerSPLExamDetails struct {
	ProgramID                string                  `json:"programId"`
	ECourseID                string                  `json:"eCourseId"`
	ECoursePatternID         string                  `json:"eCoursePatternId"`
	LearnerID                string                  `json:"learnerId"`
	SessionAggregationID     string                  `json:"sessionAggregationId"`
	SPLExamAttemptDetails    []SPLExamAttemptDetails `json:"SPLExamAttemptDetails"`
	ExamID                   string                  `json:"examId"`
	ExamPatternID            string                  `json:"examPatternId"`
	ExamType                 string                  `json:"examType"`
	AggregationID            string                  `json:"aggregationId"`
	LanguageID               string                  `json:"languageID"`
	CurrentExamAttemptNumber int                     `json:"currentExamAttemptNumber"`
	QuestionID               string                  `json:"questionId"`
}

type SPLExamAttemptDetails struct {
	ExamID             string               `json:"examId"`
	AggregationID      string               `json:"aggregationId"`
	AttemptNumber      int                  `json:"attemptNumber"`
	StepAttemptDetails []StepAttemptDetails `json:"stepAttemptDetails"`
	IsExamCompleted    bool                 `json:"isExamCompleted"`
	DurationInSec      int                  `json:"durationInSec"`
	ElapsedTime        int                  `json:"elapsedTime"`
	TotalMarksObtained float64              `json:"totalMarksObtained"`
	ExamStartedOn      time.Time            `json:"examStartedOn"`
	ExamEndedOn        time.Time            `json:"examEndedOn"`
	ExamEndCause       string               `json:"examEndCause"`
	TotalMarks         float64              `json:"totalMarks"`
	TotalSteps         int                  `json:"totalSteps"`
	LanguageID         string               `json:"languageID"`
	Percentage         float64              `json:"percentage"`
	AttemptedStepCount int                  `json:"attemptedStepCount"`
	CorrectItemCount   int                  `json:"correctItemCount"`
	WrongItemCount     int                  `json:"wrongItemCount"`
	ExamCurrentStatus  string               `json:"examCurrentStatus"`
	ExamSubmittedOn    *time.Time           `json:"examSubmittedOn,omitempty"`
}

type StepAttemptDetails struct {
	StepID                 string                 `json:"stepId"`
	StepTitle              string                 `json:"stepName"`
	StepDescription        string                 `json:"stepDescription"`
	StepType               string                 `json:"stepType"` // Random / Specific
	StepStartTime          time.Time              `json:"stepStartTime"`
	StepEndTime            time.Time              `json:"stepEndTime"`
	QuestionCategory       string                 `json:"questionCategory"` // Video, flash, flash + TAC, MCQ, practical
	TotalMarks             float64                `json:"totalMarks"`
	ElapsedTime            int                    `json:"elapsedTime"`
	TotalMarksObtained     float64                `json:"totalMarksObtained"`
	IsStepCompleted        bool                   `json:"isStepCompleted"`
	StepStatus             string                 `json:"stepStatus"`
	StepAttemptNumber      int                    `json:"stepAttemptNumber"`
	Ordinality             int                    `json:"ordinality"`
	QuestionAttemptDetails QuestionAttemptDetails `json:"questionAttemptDetails"`
}

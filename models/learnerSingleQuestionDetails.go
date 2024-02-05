package models

//go:generate ffjson $GOFILE

import "time"

//LearnerSingleQuestionDetails LearnerSingleQuestionDetails
type LearnerSingleQuestionDetails struct {
	ProgramID                    string                         `json:"programId"`
	ECourseID                    string                         `json:"eCourseId"`
	ECoursePatternID             string                         `json:"eCoursePatternId"`
	LearnerID                    string                         `json:"learnerId"`
	SessionAggregationID         string                         `json:"sessionAggregationId"`
	AggregationID                string                         `json:"aggregationId"`
	QuestionID                   string                         `json:"questionID"`
	QuestionType                 string                         `json:"questionType"`
	SingleQuestionAttemptDetails []SingleQuestionAttemptDetails `json:"singleQuestionExamAttemptDetails"`
}

//SingleQuestionAttemptDetails SingleQuestionAttemptDetails
type SingleQuestionAttemptDetails struct {
	LanguageID        string          `json:"languageID"`
	AttemptNumber     int             `json:"attemptNumber"`
	AttemptDate       time.Time       `json:"attemptDate"`
	TotalCharsDeleted int             `json:"totalCharsDeleted"`
	WrongWordsCount   int             `json:"wrongWordsCount"`
	TypedText         string          `json:"typedText"`
	TotalCharsTyped   int             `json:"totalCharsTyped"`
	SubmittedAnswer   SubmittedAnswer `json:"submittedAnswer"`
	FeedbackReceived  string          `json:"feedbackReceived"`
	QuestionStatus    string          `json:"questionStatus"`
	TypingAccuracy    float64         `json:"typingAccuracy"`
	WordsPerMinute    int             `json:"wordsPerMinute"`
}

//Single Question New Structure
type LearnerSingleQuestionSessionAggregationWiseAttemptDetails struct {
	ProgramID            string               `json:"programId"`
	ECourseID            string               `json:"eCourseId"`
	ECoursePatternID     string               `json:"eCoursePatternId"`
	LearnerID            string               `json:"learnerId"`
	SessionAggregationID string               `json:"sessionAggregationId"`
	AggregationDetails   []AggregationDetails `json:"aggregationDetails"`
}

type AggregationDetails struct {
	AggregationID                string                         `json:"aggregationId"`
	QuestionID                   string                         `json:"questionID"`
	QuestionType                 string                         `json:"questionType"`
	SingleQuestionAttemptDetails []SingleQuestionAttemptDetails `json:"singleQuestionExamAttemptDetails"`
}

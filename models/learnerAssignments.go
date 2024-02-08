package models

//go:generate ffjson $GOFILE

import "time"

type LearnerAssignments struct {
	LearnerID            string               `json:"learnerId"`
	ProgramID            string               `json:"programId"`
	ECourseID            string               `json:"eCourseId"`
	CoursePatternID      string               `json:"coursePatternId"`
	BatchID              string               `json:"batchId"`
	UploadFlag           bool                 `json:"uploadFlag"`
	UploadDate           time.Time            `json:"uploadDate"`
	LastModified         time.Time            `json:"lastModified"`
	SessionAggregationID string               `bson:"sessionAggregationID"`
	AssignmentsDetails   []AssignmentsDetails `json:"assignmentsDetails"`
}

type AssignmentsDetails struct {
	AssignmentID                string                  `json:"assignmentId"`
	AssignmentStatementID       string                  `json:"assignmentStatementId"`
	AttemptNo                   int                     `json:"attemptNo"`
	AttemptStartDate            time.Time               `json:"attemptStartDate"`
	AttemptEndDate              time.Time               `json:"attemptEndDate"`
	UploadedRelativeFilePath    string                  `json:"uploadedRelativeFilePath"`
	Rating                      string                  `json:"rating"`
	MarksObtained               float64                 `json:"marksObtained"`
	IsShared                    bool                    `json:"isShared"`
	LanguageID                  string                  `json:"languageId"`
	AssignmentStatus            string                  `json:"assignmentStatus"`
	AssignmentPortfolioStatus   string                  `json:"assignmentPortfolioStatus"`
	FileChecksum                string                  `json:"fileChecksum"`
	IsFileUploaded              bool                    `json:"isFileUploaded"`
	FileSize                    int                     `json:"fileSize"`
	CreateBy                    string                  `json:"createBy"`
	CreatedOn                   time.Time               `json:"createdOn"`
	ModifiedBy                  string                  `json:"modifiedBy"`
	ModifiedOn                  time.Time               `json:"modifiedOn"`
	CommentsDetails             []CommentsDetails       `json:"commentsDetails"`
	ExhibitionDetails           []ExhibitionDetails     `json:"exhibitionDetails"`
	EvaluationDetails           []EvaluationDetails     `json:"evaluationDetails"`
	PeerEvaluationDetails       []PeerEvaluationDetails `json:"peerEvaluationDetails"`
	Duration                    string                  `json:"duration"`
	EvidenceFilePath            string                  `json:"evidenceFilePath"`
	IsEvidenceFileUploaded      bool                    `json:"isEvidenceFileUploaded"`
	EvidenceCDNFilePath         string                  `json:"evidenceCDNFilePath"`
	IsEvidenceFileUploadedOnCDN bool                    `json:"isEvidenceFileUploadedOnCDN"`
}

type CommentsDetails struct {
	AssignmentComment string `json:"assignmentComment"`
	ExhibitionComment string `json:"exhibitionComment"`
	CommentedBy       string `json:"commentedBy"`
}

type ExhibitionDetails struct {
	ExhibitionStartDate string `json:"exhibitionStartDate"`
	ExhibitionEndDate   string `json:"exhibitionEndDate"`
	IconFile            string `json:"iconFile"`
}

type EvaluationDetails struct {
	EvaluatorID            string    `json:"evaluatorId"`
	AssignmentCriteriaID   string    `json:"assignmentCriteriaId"`
	AssignmentCriteriaName string    `json:"assignmentCriteriaName"`
	EvaluatorResponse      string    `json:"evaluatorResponse"`
	EvaluatorComments      string    `json:"evaluatorComments"`
	EvaluatedMarks         float64   `json:"evaluatedMarks"`
	EvaluationDate         time.Time `json:"evaluationDate"`
}

type PeerEvaluationDetails struct {
	PeerEvaluatorLearnerID string    `json:"peerEvaluatorLearnerId"`
	AssignmentCriteriaID   string    `json:"assignmentCriteriaId"`
	PeerEvaluatorComments  string    `json:"peerEvaluatorComments"`
	PeerEvaluatedMarks     float64   `json:"peerEvaluatedMarks"`
	PeerEvaluatorType      string    `json:"peerEvaluatorType"`
	PeerEvaluatedRating    string    `json:"peerEvaluatedRating"`
	PeerRating             float64   `json:"peerRating"`
	EvaluatorResponse      string    `json:"evaluatorResponse"`
	EvaluationDate         time.Time `json:"evaluationDate"`
}

type ProcessFolioMetaData struct {
	ProgramID     string          `json:"programId"`
	EcourseID     string          `json:"eCourseId"`
	BatchID       string          `json:"batchId"`
	ProcessFolios ProcessFolioMap `json:"processFolios"`
}

type ProcessFolioMap map[string][]ProcessFolioDetails
type ProcessFolioDetails struct {
	FilePath             string `json:"filePath" validate:"required"`
	CloudUploadFilePath  string `json:"cloudUploadFilePath"`
	LocalUploadStatus    bool   `json:"localUploadStatus"`
	CloudUploadStatus    string `json:"cloudUploadStatus"`
	AssignmentID         string `json:"assignmentId" validate:"required"`
	BatchId              string `json:"batchId" validate:"required"`
	LocalUploadDate      string `json:"localUploadDate"`
	CloudUploadDate      string `json:"cloudUploadDate"`
	SessionAggregationID string `json:"sessionAggregationId"`
	IsDeleted            bool   `json:"isDeleted"`
}

type BatchData struct {
	ProgramID string `json:"programId"`
	ECourseID string `json:"eCourseId"`
	BatchID   string `json:"batchId"`
}

type LearnerDetailsForProcessFolio struct {
	LearnerID string `json:"learnerId" validate:"required"`
	ProgramID string `json:"programId" validate:"required"`
	ECourseID string `json:"eCourseId" validate:"required"`
	BatchID   string `json:"batchId" validate:"required"`
}

package models

//go:generate ffjson $GOFILE

import (
	"time"
)

// LearnerAllocation LearnerAllocation
type LearnerAllocation struct {
	LearnerDetails        LearnerDetails      `json:"learnerDetails,omitempty"`
	AllocationDetails     []AllocationDetails `json:"allocationDetails,omitempty"`
	ExamAllocationDetails []AllocationDetails `json:"examAllocationDetails,omitempty"`
	AppTrack              AppTrack            `json:"appTrack,omitempty"`
}

// LearnerDetails LearnerDetails
type LearnerDetails struct {
	LearnerID                          string             `json:"learnerId,omitempty" validate:"required"`
	LearnerPrefUITheme                 string             `json:"learnerPrefUITheme,omitempty"`
	LearnerUILanguageID                string             `json:"learnerUILanguageId,omitempty"`
	PendingMessages                    int                `json:"pendingMessages,omitempty"`
	FirstLogin                         *time.Time         `json:"firstLogin,omitempty"`
	LatestLogin                        *time.Time         `json:"latestLogin,omitempty"`
	ResetEnableBiometricDate           *time.Time         `json:"resetEnableBiometricDate,omitempty"`
	AboutMe                            string             `json:"aboutMe,omitempty"`
	PhotoPath                          string             `json:"photoPath,omitempty"`
	SolarPhotoPath                     string             `json:"solarPhotoPath,omitempty"`
	SignPath                           string             `json:"signPath,omitempty"`
	SolarSignPath                      string             `json:"solarSignPath,omitempty"`
	LeaderBoardDetails                 LeaderBoardDetails `json:"leaderBoardDetails,omitempty"`
	RealFeelActions                    []RealFeelActions  `json:"realFeelActions,omitempty" validate:"dive,required"`
	LastBiometricAttendance            time.Time          `json:"lastBiometricAttendance"`
	LastBiometricAttendanceWithVersion string             `json:"lastBiometricAttendanceWithVersion,omitempty"`
	PassPhraseDetails                  PassPhraseDetails  `json:"passPhraseDetails,omitempty"`
	BiometricData                      BiometricData      `json:"biometricData,omitempty"`
	BiometricPunchDayCount             int                `json:"biometricPunchDayCount,omitempty"`
	FaceData                           FaceData           `json:"faceData,omitempty"`
}

// BiometricData BiometricData
type BiometricData struct {
	ResetCount             int                                    `json:"resetCount,omitempty"`
	CreatedBy              string                                 `json:"createdBy,omitempty"`
	CreatedOn              *time.Time                             `json:"createdOn,omitempty"`
	ModifiedBy             string                                 `json:"modifiedBy,omitempty"`
	IsBiometricRestored    bool                                   `json:"isBiometricRestored,omitempty"`
	LastModifiedDate       *time.Time                             `json:"lastModifiedDate,omitempty"`
	FIData                 string                                 `json:"fiData,omitempty"`
	BiometricAttendanceMap map[string]*BiometricAttendanceDetails `json:"biometricAttendanceMap,omitempty"`
}

// Biometric Attendance details
type BiometricAttendanceDetails struct {
	StartTime   time.Time `json:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
	PunchCount  int64     `json:"punchCount,omitempty"`
	DurationMin float64   `json:"durationMin,omitempty"`
	DailyPunch  string    `json:"dailyPunch,omitempty"`
}

// RealFeelActions RealFeelActions
type RealFeelActions struct {
	LastRealFeelActionTime *time.Time `json:"lastRealFeelActionTime,omitempty"`
	RealFeelActionCount    int        `json:"realFeelActionCount,omitempty"`
	RealFeelActionName     string     `json:"realFeelActionName,omitempty" validate:"required"`
}

// LeaderBoardDetails LeaderBoardDetails
type LeaderBoardDetails struct {
	CurrentLevel        int            `json:"currentLevel,omitempty"`
	CurrentPoints       int            `json:"currentPoints,omitempty"`
	CurrentRank         int            `json:"currentRank,omitempty"`
	CurrentStatus       string         `json:"currentStatus,omitempty"`
	UpcomingLevelName   string         `json:"upcomingLevelName,omitempty"`
	UpcomingLevelPoints int            `json:"upcomingLevelPoints,omitempty"`
	PointsEarned        []PointsEarned `json:"pointsEarned,omitempty"`
}

// PointsEarned PointsEarned
type PointsEarned struct {
	ActivityID      string     `json:"activityId,omitempty"`
	ActivityName    string     `json:"activityName,omitempty"`
	PointsEarned    int        `json:"pointsEarned,omitempty"`
	CurrentPoints   int        `json:"currentPoints,omitempty"`
	CurrentRank     int        `json:"currentRank,omitempty"`
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
}

// AllocationDetails AllocationDetails
type AllocationDetails struct {
	/**** LearnerScoresID value will be set in case of "manualTransfer" learner marks data transfer */
	LearnerScoresID string `json:"-"`
	/***/
	LearnerCourseID                    string                     `json:"learnerCourseId,omitempty"`
	LearningCenterID                   string                     `json:"learningCenterId,omitempty"`
	LearningCenterCode                 string                     `json:"learningCenterCode,omitempty"`
	LearningCenterName                 string                     `json:"learningCenterName,omitempty"`
	BatchID                            string                     `json:"batchId,omitempty"`
	BatchName                          string                     `json:"batchName,omitempty"`
	ProgramID                          string                     `json:"programId,omitempty"`
	OfferingID                         string                     `json:"offeringId,omitempty"`
	OfferingName                       string                     `json:"offeringName,omitempty"`
	ProgramName                        string                     `json:"programName,omitempty"`
	ECourseID                          string                     `json:"eCourseId,omitempty"`
	ECourseName                        string                     `json:"eCourseName,omitempty"`
	ECourseImage                       string                     `json:"eCourseImage,omitempty"`
	ECourseVersion                     string                     `json:"eCourseVersion,omitempty"`
	ProcessStorePath                   string                     `json:"processStorePath,omitempty"`
	DownloadDate                       *time.Time                 `json:"downloadDate,omitempty"`
	LearningStartDate                  *time.Time                 `json:"learningStartDate,omitempty"`
	LearningStartedOn                  *time.Time                 `json:"learningStartedOn,omitempty"`
	HasBlocked                         bool                       `json:"hasBlocked,omitempty"`
	BlockDate                          *time.Time                 `json:"blockDate,omitempty"`
	LearnerCoursePrefLanguageCode      string                     `json:"learnerCoursePrefLanguageCode,omitempty"`
	LearnereCourseAttendance           int                        `json:"learnereCourseAttendance,omitempty"`
	CompletionPercentage               float64                    `json:"completionPercentage,omitempty"`
	ExpectedTimeSpent                  int                        `json:"expectedTimeSpent,omitempty"`
	SolvedChallengeCount               int                        `json:"solvedChallengeCount,omitempty"`
	SolvedDIYCount                     int                        `json:"solvedDIYCount,omitempty"`
	TimeSpent                          int                        `json:"timeSpent,omitempty"`
	CompletedAggregationCount          int                        `json:"completedAggregationCount,omitempty"`
	SolvedAssignmentCount              int                        `json:"solvedAssignmentCount,omitempty"`
	CompletedSessionCount              int                        `json:"completedSessionCount,omitempty"`
	Last20DaysTimeSpent                []Last20DaysTimeSpent      `json:"last20DaysTimeSpent,omitempty"`
	CurrentSession                     string                     `json:"currentSession,omitempty"`
	CurrentSessionAggregationID        string                     `json:"currentSessionAggregationId,omitempty"`
	SessionID                          string                     `json:"sessionId,omitempty"`
	AggregationID                      string                     `json:"aggregationId,omitempty"`
	SequentialAggregationID            string                     `json:"sequentialAggregationId,omitempty"`
	NextSequentialAggregationID        string                     `json:"nextSequentialAggregationId,omitempty"`
	NextSequentialSessionAggregationID string                     `json:"nextSequentialSessionAggregationId,omitempty"`
	PrevSequentialAggregationID        string                     `json:"prevSequentialAggregationId,omitempty"`
	PrevSequentialSessionAggregationID string                     `json:"prevSequentialSessionAggregationId,omitempty"`
	VisitedTime                        *time.Time                 `json:"visitedTime,omitempty"`
	LearnerVerificationStatus          string                     `json:"learnerVerificationStatus,omitempty"`
	Barcode                            string                     `json:"barcode,omitempty"`
	HasReceivedBook                    bool                       `json:"hasReceivedBook,omitempty"`
	ReceivedDate                       *time.Time                 `json:"receivedDate,omitempty"`
	Comment                            string                     `json:"comment,omitempty"`
	HaseCourseEnd                      bool                       `json:"haseCourseEnd,omitempty"`
	LearnereCourseEndedOn              *time.Time                 `json:"learnereCourseEndedOn,omitempty"`
	PerformanceStatement               string                     `json:"performanceStatement,omitempty"`
	LastUpdateDateTime                 *time.Time                 `json:"lastUpdateDateTime,omitempty"`
	UploadDate                         *time.Time                 `json:"uploadDate,omitempty"`
	UploadFlag                         bool                       `json:"uploadFlag,omitempty"`
	CoursePatternID                    string                     `json:"coursePatternId,omitempty"`
	LearningScoreDetails               []LearningScoreDetails     `json:"learningScoreDetails,omitempty"`
	RestoreScoreDetails                []RestoreScoreDetails      `json:"restoreScoreDetails,omitempty"`
	SessionCompletionDetails           []SessionCompletionDetails `json:"sessionCompletionDetails,omitempty"`
	LearnerEvents                      []LearnerEvents            `json:"learnerEvents,omitempty"`
	CourseMetadata                     ProgramCourse              `json:"courseMetadata,omitempty"`
	IsRestored                         bool                       `json:"isRestored"`
	ClassRoomAttendedCount             int                        `json:"classRoomAttendedCount"`
	NumberOfQuestionsAnswered          int                        `json:"numberOfQuestionsAnswered"`
	IsSessionWiseCalculation           bool                       `json:"isSessionWiseCalculation"`
	// Exam Related Attributes
	ExamID                      string             `json:"examId,omitempty"`
	ExamName                    string             `json:"examName,omitempty"`
	ExamImage                   string             `json:"examImage,omitempty"`
	ExamPatternID               string             `json:"examPatternId,omitempty"`
	ExamVersion                 string             `json:"examVersion,omitempty"`
	LearnerExamPrefLanguageCode string             `json:"learnerExamPrefLanguageCode,omitempty"`
	ExamEventName               string             `json:"examEventName,omitempty"`
	ExamScheduledFrom           *time.Time         `json:"examScheduledFrom,omitempty"`
	ExamScheduledTo             *time.Time         `json:"examScheduledTo,omitempty"`
	Description                 string             `json:"description,omitempty"`
	ExamBodyName                string             `json:"examBodyName,omitempty"`
	HallTicketStatus            string             `json:"hallTicketStatus,omitempty"`
	Result                      string             `json:"result,omitempty"`
	AttemptType                 string             `json:"attemptType,omitempty"`
	ExamBriefDetails            []ExamBriefDetails `json:"examBriefDetails"`
	LastLearningDateTime        *time.Time         `json:"lastLearningDateTime,omitempty"`

	// Elective Course related attributes
	IsElectiveCourse       bool                `json:"isElectiveCourse"`
	IsBundleUploadRequired bool                `json:"isBundleUploadRequired,omitempty"`
	IsCourseCompleted      bool                `json:"isCourseCompleted,omitempty"`
	CourseStatus           string              `json:"courseStatus,omitempty"`
	CourseSelectedOn       *time.Time          `json:"courseSelectedOn,omitempty"`
	CourseCompletedOn      *time.Time          `json:"courseCompletedOn,omitempty"`
	CourseExpiredOn        *time.Time          `json:"courseExpiredOn,omitempty"`
	ParentCourseDetails    ParentCourseDetails `json:"parentCourseDetails,omitempty"`
	IsCertificateGenerated bool                `json:"isCertificateGenerated,omitempty"`
	AverageTimeSpent       int                 `json:"averageTimeSpent,omitempty"`
}

type ExamBriefDetails struct {
	AggregationID       string          `json:"aggregationId,omitempty"`
	AttemptNumber       int             `json:"attemptNumber,omitempty"`
	IsExamCompleted     bool            `json:"isExamCompleted,omitempty"`
	DurationInSec       int             `json:"durationInSec,omitempty"`
	TotalItems          int             `json:"totalItems,omitempty"`
	TotalMarks          float64         `json:"totalMarks,omitempty"`
	ExamStartedOn       *time.Time      `json:"examStartedOn,omitempty"`
	ExamEndedOn         *time.Time      `json:"examEndedOn,omitempty"`
	ExamSubmittedOn     *time.Time      `json:"examSubmittedOn,omitempty"`
	ElapsedTime         int             `json:"elapsedTime,omitempty"`
	AttemptedItemCount  int             `json:"attemptedItemCount,omitempty"`
	ExamEvaluatedOn     *time.Time      `json:"examEvaluatedOn,omitempty"`
	TotalMarksObtained  float64         `json:"totalMarksObtained,omitempty"`
	CorrectItemCount    int             `json:"correctItemCount,omitempty"`
	WrongItemCount      int             `json:"wrongItemCount,omitempty"`
	SectionBriefDetails []SectionDetail `json:"sectionBriefDetails,omitempty"`
	ExamCurrentStatus   string          `json:"examCurrentStatus,omitempty"`
	LastUpdatedOn       *time.Time      `json:"lastUpdatedOn,omitempty"`
	// NOTE: CET
	CloudEvidenceUploadStatus string  `json:"cloudEvidenceUploadStatus"`
	LocalEvidenceUploadStatus string  `json:"localEvidenceUploadStatus"`
	EvidenceUploadedPath      string  `json:"evidenceUploadedPath"`
	MinimumPassingMarks       float64 `json:"minimumPassingMarks"`
}

type SectionDetail struct {
	SectionID                      string  `json:"sectionId"`
	SectionName                    string  `json:"sectionName"`
	Ordinality                     int     `json:"ordinality"`
	DurationInSec                  int     `json:"durationInSec"`
	TotalMarks                     float64 `json:"totalMarks"`
	TotalItems                     int     `json:"totalItems"`
	ElapsedTime                    int     `json:"elapsedTime"`
	AttemptedItemCount             int     `json:"attemptedItemCount"`
	CorrectItemCount               int     `json:"correctItemCount"`
	WrongItemCount                 int     `json:"wrongItemCount"`
	TotalMarksObtained             float64 `json:"totalMarksObtained"`
	IsSectionCompleted             bool    `json:"isSectionCompleted"`
	SectionStatus                  string  `json:"sectionStatus"`
	IsSectionTalentZone            bool    `json:"isSectionTalentZone"`
	TalentZoneQuestionDisplayCount int     `json:"talentZoneQuestionDisplayCount"`
}

type Last20DaysTimeSpent struct {
	TimeSpentDate *time.Time `json:"timeSpentDate,omitempty"`
	TimeSpent     int64      `json:"timeSpent,omitempty"`
}

// LearningScoreDetails LearningScoreDetails
type LearningScoreDetails struct {
	PassingHeadCode        string       `json:"passingHeadCode,omitempty"`
	PassingHeadName        string       `json:"passingHeadName,omitempty"`
	PassingHeadDisplayName string       `json:"passingHeadDisplayName,omitempty"`
	SkipAbsentSessionMarks bool         `json:"skipAbsentSessionMarks,omitempty"`
	Score                  float64      `json:"score,omitempty"`
	AdditionalScore        float64      `json:"additionalScore,omitempty"`
	ScoreDetails           ScoreDetails `json:"scoreDetails,omitempty"`
	MaxMarks               float64      `json:"maxMarks,omitempty"`
	TypingAcuracy          float64      `json:"typingAccuracy,omitempty"`
}

// ScoreDetails
type ScoreDetails struct {
	TotalScore      float64 `json:"totalScore,omitempty"`
	AdditionalMarks float64 `json:"additionalScore,omitempty"`
}

type RestoreScoreDetails struct {
	PassingHeadCode string  `json:"passingHeadCode" db:"PASSINGHEADCODE"`
	RestoreScore    float64 `json:"restoreScore"`
}

// SessionCompletionDetails SessionCompletionDetails
type SessionCompletionDetails struct {
	AggregationID                              string                    `json:"aggregationId,omitempty"`
	AggregationName                            string                    `json:"aggregationName,omitempty"`
	CompletedOn                                *time.Time                `json:"completedOn,omitempty"`
	IsCompleted                                bool                      `json:"isCompleted,omitempty"`
	SessionStatus                              string                    `json:"sessionStatus,omitempty"`
	IsClassRoomAttempted                       bool                      `json:"isClassRoomAttempted,omitempty"`
	AggregationDuration                        int                       `json:"aggregationDuration,omitempty"`
	TotalAggregationDuration                   int                       `json:"totalAggregationDuration,omitempty"`
	ActivityCount                              int                       `json:"activityCount,omitempty"`
	Ordinality                                 int                       `json:"ordinality,omitempty"`
	PerformanceRating                          int                       `json:"performanceRating,omitempty"`
	ChallengeCount                             int                       `json:"challengeCount,omitempty"`
	HasAdditionalCriteriaAggregationID1Visited bool                      `json:"hasAdditionalCriteriaAggregationId1Visited,omitempty"`
	HasAdditionalCriteriaAggregationID2Visited bool                      `json:"hasAdditionalCriteriaAggregationId2Visited,omitempty"`
	MaxViewedInLanguageCode                    string                    `json:"maxViewedInLanguageCode,omitempty"`
	LastUpdateDateTime                         *time.Time                `json:"lastUpdateDateTime,omitempty"`
	SkillBank                                  []SkillBank               `json:"skillBank,omitempty"`
	SessionMarks                               float32                   `json:"sessionMarks,omitempty"`
	IsSysEntry                                 bool                      `json:"isSysEntry,omitempty"`
	SessionClassroomDetails                    []SessionClassroomDetails `json:"sessionClassroomDetails"`
	PassingHeadWiseMarks                       []LearningScoreDetails    `json:"passingHeadWiseMarks,omitempty"`
	TotalClassroomDuration                     int                       `json:"totalClassroomDuration,omitempty"`
	OfflineContentStatus                       OfflineContentStatus      `json:"offlineContentStatus,omitempty"`
}

// LearnerEvents LearnerEvents
type LearnerEvents struct {
	EventCompletedTime *time.Time `json:"eventCompletedTime,omitempty"`
	EventID            string     `json:"eventId,omitempty"`
	TriggerTime        *time.Time `json:"triggerTime,omitempty"`
}

// SkillBank SkillBank
type SkillBank struct {
	SkillLevelAchieved string `json:"skillLevelAchieved,omitempty"`
	SkillName          string `json:"skillName,omitempty"`
}

// PassPhraseDetails PassPhraseDetails
type PassPhraseDetails struct {
	Phrase                string     `json:"phrase,omitempty"`
	Attempts              int        `json:"attempts,omitempty"`
	FailureCount          int        `json:"failureCount,omitempty"`
	Average               float64    `json:"average,omitempty"`
	LastAttemptPercentage float64    `json:"lastAttemptPercentage,omitempty"`
	RegisteredTappy       string     `json:"registeredTappy,omitempty"`
	UnlockedByLFCount     int        `json:"unlockedByLFCount,omitempty"`
	LastUnlockedByLFDate  *time.Time `json:"lastUnlockedByLFDate,omitempty"`
}

// SessionClassroomDetails SessionClassroomDetails
type SessionClassroomDetails struct {
	AggregationID string     `json:"aggregationId,omitempty"`
	CreatedDate   *time.Time `json:"createdDate"`
	ModifiedDate  *time.Time `json:"modifiedDate"`
}

// AppTrack AppTrack
type AppTrack struct {
	TotalTime int       `json:"TotalActiveTime,omitempty"`
	AppData   []AppData `json:"AppData,omitempty"`
	WebData   []WebData `json:"WebData"`
}

type AppData struct {
	ProcessName string `json:"ProcessName,omitempty"`
	AppName     string `json:"AppName,omitempty"`
	UsedTime    int    `json:"UsedTime,omitempty"`
}

type WebData struct {
	Url      string `json:"url,omitempty"`
	UsedTime int    `json:"usedTime,omitempty"`
}

type OfflineContentStatus struct {
	ProgramID            string `json:"programId"`
	ECourseID            string `json:"eCourseId"`
	ECoursePatternID     string `json:"eCoursePatternId"`
	LanguageID           string `json:"languageId"`
	SessionAggregationID string `json:"sessionAggregationId"`
	OflDataStatus        string `json:"oflDataStatus"`
	LearnerID            string `json:"learnerId"`
	DownloadPath         string `json:"downloadPath"`
	FileSize             int64  `json:"fileSize"`
}

type ParentCourseDetails struct {
	ProgramID        string `json:"programId,omitempty"`
	ECourseID        string `json:"eCourseId,omitempty"`
	ECoursePatternID string `json:"eCoursePatternId,omitempty"`
	BatchID          string `json:"batchId,omitempty"`
	SessionID        string `json:"sessionId,omitempty"`
	AggregationID    string `json:"aggregationId,omitempty"`
	LanguageID       string `json:"languageId,omitempty"`
	// SelectedOn       *time.Time `json:"selectedOn,omitempty"`
	ECMNodeDuration int `json:"ecmNodeDuration,omitempty"`
	BufferDuration  int `json:"bufferDuration,omitempty"`
}

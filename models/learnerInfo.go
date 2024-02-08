package models

import "time"

// LearnerInfo -
type LearnerInfo struct {
	LearnerID                string               `json:"learnerId" validate:"required"`
	LoginID                  string               `json:"loginId" validate:"required"`
	LoginPassword            string               `json:"loginPassword"`
	LearnerCode              string               `json:"learnerCode" validate:"required"`
	LegalFullName            string               `json:"legalFullName" validate:"required"`
	Gender                   string               `json:"gender" validate:"required"`
	DateOfBirth              time.Time            `json:"dateOfBirth" validate:"required"`
	MobileNo                 string               `json:"mobileNo"`
	EmailID                  string               `json:"emailId"`
	PhotoPath                string               `json:"photoPath"`
	SignPath                 string               `json:"signPath"`
	CurrentAddress           string               `json:"currentAddress"`
	PhysicallyChallenge      string               `json:"physicallyChallenge"`
	HasBraille               bool                 `json:"hasBraille"`
	MobileNo2                string               `json:"mobileNo2"`
	MobileNo2Relation        string               `json:"mobileNo2Relation"`
	ApprovalStatus           string               `json:"approvalStatus"`
	ActionBy                 string               `json:"actionBy"`
	RejectionDate            time.Time            `json:"rejectionDate"`
	RejectionReason          string               `json:"rejectionReason"`
	FingerPrintStatus        string               `json:"fingerPrintStatus"`
	CourseAllocation         []CourseAllocation   `json:"courseAllocation" validate:"dive"`
	ExamAllocation           []ExamAllocation     `json:"examAllocation,omitempty" validate:"dive"`
	ResetEnableBiometricDate *time.Time           `json:"resetEnableBiometricDate,omitempty"`
	FaceData                 FaceData             `json:"faceData,omitempty"`
	LearnerBiometricData     LearnerBiometricData `json:"learnerBiometricData,omitempty"`
	// JvID                     string               `json:"jvID,omitempty"`
}

type ExamAllocation struct {
	//This is SOLAR attribute
	LearnerCourseID             string                `json:"learnerCourseId" validate:"required"`
	LearningCenterID            string                `json:"learningCenterId" validate:"required"`
	LearningCenterCode          string                `json:"learningCenterCode" validate:"required"`
	LearningCenterName          string                `json:"learningCenterName" validate:"required"`
	ProgramID                   string                `json:"programId" validate:"required"`
	ProgramName                 string                `json:"programName" validate:"required"`
	ExamID                      string                `json:"examId" validate:"required"`
	ExamName                    string                `json:"examName" validate:"required"`
	SolarCourseID               string                `json:"solarCourseId" validate:"required"`
	SolarCourseName             string                `json:"solarCourseName" validate:"required"`
	BatchID                     string                `json:"batchId" validate:"required"`
	BatchName                   string                `json:"batchName" validate:"required"`
	BatchDuration               string                `json:"batchDuration" validate:"required"`
	ExamEventName               string                `json:"examEventName" validate:"required"`
	ExamFromDate                time.Time             `json:"examFromDate" validate:"required"`
	ExamToDate                  time.Time             `json:"examToDate" validate:"required"`
	Description                 string                `json:"description,omitempty"`
	ExamBodyName                string                `json:"examBodyName,omitempty"`
	HallTicketStatus            string                `json:"hallTicketStatus,omitempty"`
	Result                      string                `json:"result,omitempty"`
	AttemptType                 string                `json:"attemptType,omitempty"`
	EvidenceMap                 map[string][]Evidence `json:"evidences,omitempty"`
	ExamPatternID               string                `json:"examPatternId" validate:"required"`
	LearnerExamPrefLanguageCode string                `json:"learnerExamPrefLanguageCode" validate:"required"`
	ExamBriefDetails            []ExamBriefDetails    `json:"examBriefDetails,omitempty"`
}

type CourseAllocation struct {
	//This is SOLAR attribute
	LearnerCourseID    string `json:"learnerCourseId" validate:"required"`
	LearningCenterID   string `json:"learningCenterId" validate:"required"`
	LearningCenterCode string `json:"learningCenterCode" validate:"required"`
	LearningCenterName string `json:"learningCenterName" validate:"required"`
	ProgramID          string `json:"programId" validate:"required"`
	ProgramName        string `json:"programName" validate:"required"`
	ECourseID          string `json:"eCourseId" validate:"required"`
	ECourseName        string `json:"eCourseName" validate:"required"`
	SolarCourseID      string `json:"solarCourseId" validate:"required"`
	SolarCourseName    string `json:"solarCourseName"`
	BatchID            string `json:"batchId" validate:"required"`
	BatchName          string `json:"batchName" validate:"required"`
	//CourseBatchConfigurations CourseBatchConfigurations `json:"batchConfigurations,omitempty"`
	CoursePatternID           string                `json:"coursePatternId" validate:"required"`
	LearnereCourseStartDate   time.Time             `json:"learnereCourseStartDate" validate:"required"`
	LearnereCourseEndDate     time.Time             `json:"learnereCourseEndDate" validate:"required"`
	ECourseAttempt            int                   `json:"eCourseAttempt"`
	ECourseLanguagePreference string                `json:"eCourseLanguagePreference" validate:"required"`
	BatchDuration             string                `json:"batchDuration" validate:"required"`
	LearningMode              int                   `json:"learningMode"`
	LearningLimit             int                   `json:"learningLimit"`
	ExamEventName             string                `json:"examEventName"`
	ExamFromDate              time.Time             `json:"examFromDate"`
	ExamToDate                time.Time             `json:"examToDate"`
	Description               string                `json:"description"`
	ExamBodyName              string                `json:"examBodyName"`
	HallTicketStatus          string                `json:"hallTicketStatus"`
	Result                    string                `json:"result"`
	AttemptType               string                `json:"attemptType"`
	LeaderBoardScore          int                   `json:"leaderBoardScore"`
	LeaderBoardGrade          string                `json:"leaderBoardGrade"`
	GlobalRank                int                   `json:"globalRank"`
	LearnerBookPrefLanguage   string                `json:"learnerBookPrefLanguage"`
	RestoreScoreDetails       []RestoreScoreDetails `json:"restoreScoreDetails,omitempty"`
	AggregationID             string                `json:"visitedAggregationId,omitempty"`
	SequentialAggregationID   string                `json:"sequentialVisitedAggregationId,omitempty"`
	VisitedTime               time.Time             `json:"visitedTime,omitempty"`
	IsRestored                bool                  `json:"isRestored"`
	SolarScoreDetails         []SolarScoreDetails   `json:"solarScoreDetails,omitempty"`
	EvidenceMap               map[string][]Evidence `json:"evidences"`
	PerDaySessionLimit        int                   `json:"perDaySessionLimit,omitempty"`
	LearningCenterStaff       []LearningCenterStaff `json:"centerStaff,omitempty"`
	LearningCenterAddress     string                `json:"learningCenterAddress,omitempty"`
	DistrictID                string                `json:"districtId,omitempty"`
	DistrictName              string                `json:"districtName,omitempty"`
	TalukaID                  string                `json:"talukaId,omitempty"`
	TalukaName                string                `json:"talukaName,omitempty"`
	City                      string                `json:"city,omitempty"`
	PinCode                   string                `json:"pincode,omitempty"`
}

type Evidence struct {
	AggregationID string        `json:"aggregationId"`
	AssignmentID  string        `json:"assignmentId"`
	TypeCode      string        `json:"typeCode"`
	FileDetails   []FileDetails `json:"fileDetails"`
}

type FileDetails struct {
	FileExtension string `json:"fileExtension"`
	FilePath      string `json:"filePath"`
	EvidenceDate  string `json:"evidenceDate"`
}

//SolarScoreDetails -  SolarScoreDetails
type SolarScoreDetails struct {
	PassingHeadCode        string  `json:"passingHeadCode,omitempty"`
	PassingHeadName        string  `json:"passingHeadName,omitempty"`
	PassingHeadDisplayName string  `json:"passingHeadDisplayName,omitempty"`
	Score                  float64 `json:"score"`
	MaxMarks               float64 `json:"maxMarks,omitempty"`
}

type SectionBriefDetail struct {
	SectionId          string  `json:"sectionId,omitempty"`
	SectionName        string  `json:"sectionName,omitempty"`
	SectionalMarks     int     `json:"sectionalMarks,omitempty"`
	DurationInSec      int     `json:"durationInSec,omitempty"`
	ElapsedTime        int     `json:"elapsedTime,omitempty"`
	TotalItems         int     `json:"totalItems,omitempty"`
	AttemptedItemCount int     `json:"attemptedItemCount,omitempty"`
	CorrectItemCount   int     `json:"correctItemCount,omitempty"`
	WrongItemCount     int     `json:"wrongItemCount,omitempty"`
	TotalMarksObtained float64 `json:"totalMarksObtained,omitempty"`
	TotalMarks         float64 `json:"totalMarks,omitempty"`
	IsSectionCompleted bool    `json:"isSectionCompleted,omitempty"`
}

//FaceData FaceData
type FaceData struct {
	ResetCount                  int         `json:"resetCount,omitempty"`
	CreatedBy                   string      `json:"createdBy,omitempty"`
	ModifiedBy                  string      `json:"modifiedBy,omitempty"`
	LastModifiedDate            *time.Time  `json:"lastModifiedDate,omitempty"`
	DRCCHumanFaceDetectedCount  int         `json:"drccHumanFaceDetectedCount,omitempty"`
	LocalSDCPhotoPath           string      `json:"localSDCPhotoPath,omitempty"`
	SDCPhotoPath                string      `json:"sdcPhotoPath,omitempty"`
	SDCPhotoBase64              string      `json:"sdcPhotoBase64,omitempty"`
	ConfidenceWithDRCCPhoto     float64     `json:"confidenceWithDRCCPhoto,omitempty"`
	ConfidenceWithSDCPhoto      float64     `json:"confidenceWithSDCPhoto,omitempty"`
	SDCPhotoDescriptors         interface{} `json:"sdcPhotoDescriptors,omitempty"`
	DRCCPhotoDescriptors        interface{} `json:"drccPhotoDescriptors,omitempty"`
	HashCode                    string      `json:"hashCode,omitempty"`
	IsVerified                  bool        `json:"isVerified,omitempty"`
	LastFaceDetectionAttendance *time.Time  `json:"lastFaceDetectionAttendance,omitempty"`
	FaceDetectionDayCount       int         `json:"faceDetectionDayCount,omitempty"`
}

type LearnerBiometricData struct {
	IsVerified bool       `json:"isVerified"`
	FiData     string     `json:"fiData"`
	CreatedOn  *time.Time `json:"createdOn"`
}

// CourseBatchConfigurations -
// type CourseBatchConfigurations struct {
// 	DayStartRequired                  bool               `json:"dayStartRequired"`
// 	PerDaySessionLimit                int                `json:"perDaySessionLimit"`
// 	BatchDurationInDays               int                `json:"batchDurationInDays"`
// 	ActionDates                       []ActionDates      `json:"actionDates"`
// 	SecBrowserConfig                  []SecBrowserConfig `json:"secBrowserConfig,omitempty"`
// 	DisableWithinSessionSequentiality bool               `json:"disableWithinSessionSequentiality,omitempty"`
// 	IsClassroomRequired               bool               `json:"isClassroomRequired,omitempty"`
// }

// LearningCenterStaff -
type LearningCenterStaff struct {
	StaffID           int    `json:"staffId,omitempty"`
	StaffPostID       int    `json:"staffPostId,omitempty"`
	LoginID           string `json:"loginId,omitempty"`
	StaffStatus       string `json:"staffStatus,omitempty"`
	StaffPin          string `json:"staffPin,omitempty"`
	LoginPassword     string `json:"loginPassword,omitempty"`
	LearningCenterID  string `json:"learningCenterId,omitempty"`
	PostName          string `json:"postName,omitempty"`
	StaffName         string `json:"staffName,omitempty"`
	StaffPhoto        string `json:"staffPhoto,omitempty"`
	Gender            string `json:"gender,omitempty"`
	MobileNo          string `json:"mobileNo,omitempty"`
	EmailID           string `json:"emailId,omitempty"`
	DisplayOrdinality int    `json:"displayOrdinality,omitempty"`
}

type LearnerBasicDtls struct {
	LearnerID  string `json:"learnerId"`
	CenterCode string `json:"centerCode"`
}

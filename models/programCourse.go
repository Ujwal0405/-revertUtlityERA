package models

import "time"

//go:generate ffjson $GOFILE

//ProgramCourse structure
type ProgramCourse struct {
	EContentDuration      int                     `json:"eContentDuration,omitempty"`
	TotalAggregations     int                     `json:"totalAggregations,omitempty"`
	EContentBlockDuration int                     `json:"eContentBlockDuration,omitempty"`
	HasPublished          bool                    `json:"hasPublished,omitempty"`
	ProcessStorePath      string                  `json:"processStorePath,omitempty"`
	HasArchived           bool                    `json:"hasArchived,omitempty"`
	ECourseDescription    string                  `json:"eCourseDescription,omitempty"`
	ECourseVersion        string                  `json:"eCourseVersion,omitempty"`
	ReadyToDownload       bool                    `json:"readyToDownload,omitempty"`
	MultiLogin            bool                    `json:"multiLogin,omitempty"`
	EditAccessLevel       string                  `json:"editAccessLevel,omitempty"`
	HasEasyPoints         bool                    `json:"hasEasyPoints,omitempty"`
	HasCourseDetail       bool                    `json:"hasCourseDetail,omitempty"`
	LearningMode          string                  `json:"learningMode,omitempty"`
	CourseView            string                  `json:"courseView,omitempty"`
	IconFilePath          string                  `json:"iconFilePath,omitempty"`
	IconFileChecksum      string                  `json:"iconFileChecksum,omitempty"`
	DefaultPendencySize   int                     `json:"defaultPendencySize,omitempty"`
	HasSBEnabled          bool                    `json:"hasSBEnabled,omitempty"`
	CourseSessionDetails  []CourseSessionDetails  `json:"courseSessionDetails,omitempty"`
	LearningPendency      []LearningPendency      `json:"learningPendency,omitempty"`
	ECourseDisplay        []ECourseDisplay        `json:"eCourseDisplay,omitempty"`
	CourseMetadata        []CourseMetadata        `json:"courseMetadata,omitempty"`
	MarkingScheme         []MarkingSchemeDetail   `json:"markingScheme,omitempty"`
	CourseIntro           CourseIntro             `json:"courseIntro,omitempty"`
	WeekWiseCourseDetails []WeekWiseCourseDetails `json:"weekWiseCourseDetails,omitempty"`
	EcourseImage          string                  `json:"eCourseImage,omitempty"`
	ECourseBgColourCode   string                  `json:"eCourseBgColourCode,omitempty"`
	ExamMetdata           ExamMetadata            `json:"examMetadata,omitempty"`
	// Elective Course related attributes
	IsElectiveCourse            bool                   `json:"isElectiveCourse,omitempty"`
	IntroFilePath               string                 `json:"introFilePath,omitempty"`
	CourseDuration              int                    `json:"courseDuration,omitempty"`
	BufferDuration              int                    `json:"bufferDuration,omitempty"`
	IsBundleUploadRequired      bool                   `json:"isBundleUploadRequired,omitempty"`
	Sectors                     []Sector               `json:"sectors,omitempty"`
	ElectiveCourseDetails       []ProgramCourseDetails `json:"electiveCourseDetails,omitempty"`
	CourseProcessConfig         CourseProcessConfig    `json:"courseProcessConfig,omitempty"`
	IsCourseCertificateRequired bool                   `json:"isCourseCertificateRequired,omitempty"`
	JobDetails                  JobDetails             `json:"jobDetails,omitempty"`
}

// ExamMetadata -
type ExamMetadata struct {
	ExamIcon      string `json:"examIcon"`
	ExamID        string `json:"examId"`
	ExamName      string `json:"examName"`
	ExamPatternID string `json:"examPatternId"`
	ExamVersion   string `json:"examVersion"`
	LanguageID    string `json:"languageID"`
	LanguageName  string `json:"languageName"`
	MetadataName  string `json:"metadataName"`
}

//WeekWiseCourseDetails structure
type WeekWiseCourseDetails struct {
	WeekNumber   string `json:"weekNumber"`
	IconFilePath string `json:"iconFilePath"`
}

//CourseSessionDetails CourseSessionDetails
type CourseSessionDetails struct {
	SessionNumber               string   `json:"sessionNumber,omitempty"`
	Ordinality                  int32    `json:"ordinality,omitempty"`
	ToBeCompletedInWeek         string   `json:"toBeCompletedInWeek,omitempty"`
	ToBeCompletedOnDay          string   `json:"toBeCompletedOnDay,omitempty"`
	ExpiringOnDayIfNotCompleted string   `json:"expiringOnDayIfNotCompleted,omitempty"`
	SessionName                 string   `json:"sessionName,omitempty"`
	AggregationID               string   `json:"aggregationId,omitempty"`
	AggregationName             string   `json:"aggregationName,omitempty"`
	SessionHighlights           []string `json:"sessionHighlights,omitempty"`
	SessionTreeFilePath         string   `json:"sessionTreeFilePath,omitempty"`
	FileChecksum                string   `json:"fileChecksum,omitempty"`
	FileSize                    int      `json:"fileSize,omitempty"`
}

//LearningPendency LearningPendency
type LearningPendency struct {
	PendingSize int    `json:"pendingSize,omitempty"`
	BatchID     string `json:"batchId,omitempty"`
}

//MarkingSchemeDetail - markingScheme for eCourse object
type MarkingSchemeDetail struct {
	HasLinear       bool    `json:"hasLinear,omitempty"`
	PassingHeadCode string  `json:"passingHeadCode,omitempty"`
	MaxMarks        float64 `json:"maxMarks,omitempty"`
	MinMarks        float64 `json:"minMarks,omitempty"`
	EligibilityMode string  `json:"eligibilityMode,omitempty"`
}

//MarkingScheme for 8_7 - ecourse marking scheme and details object.
//MarkingScheme MarkingScheme
type MarkingScheme struct {
	ProgramID            string                 `json:"programId,omitempty"`
	ECourseID            string                 `json:"courseId,omitempty"`
	CoursePatternID      string                 `json:"coursePatternId,omitempty"`
	PassingHeadDetails   []PassingHeadDetails   `json:"passingHeadDetails,omitempty"`
	MarkingSchemeDetails []MarkingSchemeDetails `json:"markingSchemeDetails,omitempty"`
}

// PassingHeadDetails -
type PassingHeadDetails struct {
	LanguageID             string  `json:"languageId,omitempty"`
	PassingHeadName        string  `json:"passingHeadName,omitempty"`
	PassingHeadCode        string  `json:"passingHeadCode,omitempty"`
	PassingHeadDisplayText string  `json:"passingHeadDisplayText,omitempty"`
	SkipAbsentSessionMarks bool    `json:"skipAbsentSessionMarks,omitempty"`
	ExternalPassingHead    string  `json:"externalPassingHead,omitempty"`
	CalculationMethod      string  `json:"calculationMethod,omitempty"`
	HasLinear              bool    `json:"hasLinear,omitempty"`
	EligibilityMode        string  `json:"eligibilityMode,omitempty"`
	MinMarks               float64 `json:"minMarks,omitempty"`
	MaxMarks               float64 `json:"maxMarks,omitempty"`
	HasDeleted             bool    `json:"hasDeleted,omitempty"`
	DisableTransfer        bool    `json:"disableTransfer,omitempty"`
	HideInLearnerLogin     bool    `json:"hideInLearnerLogin,omitempty"`
	UpdateInParentCourse   bool    `json:"updateInParentCourse,omitempty"`
}

// MarkingSchemeDetails -
type MarkingSchemeDetails struct {
	AggregationID   string  `json:"aggregtionID,omitempty"`
	PassingHeadCode string  `json:"passingHeadCode,omitempty"`
	Marks           float64 `json:"marks,omitempty"`
	OutOfMarks      float64 `json:"outOfMarks,omitempty"`
}

//CourseMetadata - CourseMetadata
type CourseMetadata struct {
	ECourseMetadataID        string `json:"eCourseMetadataId,omitempty"`
	MetadataName             string `json:"metadataName,omitempty"`
	LanguageID               string `json:"languageId,omitempty"`
	ECourseID                string `json:"eCourseId,omitempty"`
	MetadataFileRelativePath string `json:"metadataFileRelativePath,omitempty"`
	EcourseImage             string `json:"EcourseImage,omitempty"`
	Ordinality               int    `json:"ordinality,omitempty"`
	FileChecksum             string `json:"fileChecksum,omitempty"`
	FileSize                 string `json:"fileSize,omitempty"`
	MetadataType             string `json:"metadataType,omitempty"`
}

//ECourseDisplay ECourseDisplay
type ECourseDisplay struct {
	LanguageID         string `json:"languageId,omitempty"`
	LanguageName       string `json:"languageName,omitempty"`
	ECourseDisplayName string `json:"eCourseDisplayName,omitempty"`
	HasDeleted         bool   `json:"hasDeleted,omitempty"`
	TreeFileChecksum   string `json:"treeFileChecksum,omitempty"`
	FileSize           string `json:"fileSize,omitempty"`
}

//CourseIntro CourseIntro
type CourseIntro struct {
	CourseIntro1 string `json:"courseIntro1,omitempty"`
	CourseIntro2 string `json:"courseIntro2,omitempty"`
	CourseIntro3 string `json:"courseIntro3,omitempty"`
	CourseIntro4 string `json:"courseIntro4,omitempty"`
	CourseIntro5 string `json:"courseIntro5,omitempty"`
}

//Sector structure
type Sector struct {
	SectorID          string `json:"sectorID,omitempty"`
	SectorName        string `json:"sectorName,omitempty"`
	SectorDescription string `json:"sectorDescription,omitempty"`
	IconFilePath      string `json:"iconFilePath,omitempty"`
}

// CourseProcessConfig - while saving to file []CourseConfigMapping will be used
// -1 value of programId and eCourseId will specify defult configuration
type CourseProcessConfig struct {
	ConfigurationName    string                 `json:"configurationName,omitempty"`
	ProcessConfiguration []ProcessConfiguration `json:"processConfiguration,omitempty"`
	SessionConfig        SessionConfig          `json:"sessionConfig,omitempty"`
}

type ProgramCourseDetails struct {
	ProgramID       string `json:"programId,omitempty"`
	ECourseID       string `json:"courseId,omitempty"`
	CoursePatternID string `json:"coursePatternId,omitempty"`
}

type JobDetails struct {
	AvailableJobs        string     `json:"availableJobs,omitempty"`
	AverageBidValue      string     `json:"averageBidValue,omitempty"`
	AvailableFreelancers string     `json:"availableFreelancers,omitempty"`
	CreatedOn            *time.Time `json:"createdOn,omitempty"`
	ModifiedOn           *time.Time `json:"modifiedOn,omitempty"`
}

type ProcessConfiguration struct {
	EventID string  `json:"eventId"`
	Steps   []Steps `json:"steps"`
}

type Steps struct {
	Function string `json:"function"`
	OnError  string `json:"onError"`
}

type SessionConfig struct {
	OverallMax                     string `json:"overallMax,omitempty"`
	PendingMax                     string `json:"pendingMax,omitempty"`
	PendingLength                  string `json:"pendingLength,omitempty"`
	CurrentMax                     string `json:"currentMax,omitempty"`
	MaxMissedClassroom             string `json:"maxMissedClassroom,omitempty"`
	MaxSessionStartbyLFInClassroom string `json:"maxSessionStartbyLFInClassroom,omitempty"`
}

package models

import "os"

const (
	LEARNER_COURSE_ALLOCATION_COLLECTION_NAME           = "Learner_Course_Allocation"
	LEARNING_SESSION_COMPLETION_DETAILS_COLLECTION_NAME = "Learning_Session_CompletionDtls"
	LEARNER_SESSION_VISIT_DETAILS_COLLECTION_NAME       = "3_8_LearnerSessionVisitDtls"
	LEARNER_INFORMATION_COLLECTION_NAME                 = "Learner_Information"
	LEARNER_ASSIGNMENT_DETAILS_COLLECTION_NAME          = "learnerAssignmentDtls"
	LEARNER_TEST_DETAILS_COLLECTION_NAME                = "Learner_TestDtls"
	LEARNING_TEST_ATTEMP_DETAILS_COLLECTION_NAME        = "Test_AttemptDtls"
	LEARNER_QUESTION_ATTEMP_DETAILS_COLLECTION_NAME     = "Test_questionAttemptDtls"
)

var (
	LOCAL_CREATE_FILES                          = "lc/"
	CENTER_CODE_ABBREVIATION                    = "cc_"
	CENTERCODE                                  = ""
	REGIONCODE                                  = ""
	FILE_PREFIX                                 = ""
	DBROOTPATH                                  = "D:/NewERAMigration"
	FILE_SEPARATOR                              = "/"
	CENTER_DIRECTORY_NAME                       = CENTER_CODE_ABBREVIATION + CENTERCODE + FILE_SEPARATOR
	CENTER_LOCAL_CREATE_FILES                   = LOCAL_CREATE_FILES + CENTER_DIRECTORY_NAME
	LEARNER_LOCAL_CREATE_FILES                  = CENTER_LOCAL_CREATE_FILES + LEARNER_DIRECTORY_NAME
	LEARNER_DIRECTORY_NAME                      = "learners/"
	COURSE_ACTIVITY_DIRECTORY_NAME              = "courseactivity/"
	BUNDLE_DIRECTORY_NAME                       = "bundles"
	FILE_EXTENTION                              = ".json"
	LC_LEARNER_ALLOCATION_DETAILS_MAPPING_CODES = FILE_PREFIX + "3_7_"
	LC_LEARNER_COURSE_VISIT_MAPPING_CODES       = FILE_PREFIX + "3_8_"
	LC_LEARNER_ASSIGNMENTS_MAPPING_CODES        = FILE_PREFIX + "3_12_"
	LC_LEARNER_QUESTION_ATTEMPTS_MAPPING_CODES  = FILE_PREFIX + "3_13_"
	LC_LEARNER_SPL_EXAM_ATTEMPTS_MAPPING_CODES  = FILE_PREFIX + "3_16_"
	FILE_UPLOAD_3_8_FILE                        = "3_8"
	FILE_UPLOAD_3_12_FILE                       = "3_12"
	FILE_UPLOAD_3_13_FILE                       = "3_13"
	FILE_UPLOAD_3_14_FILE                       = "3_14"
	FILE_UPLOAD_3_16_FILE                       = "3_16"
	FILE_UPLOAD_3_64_FILE                       = "3_64"
	EXCLUDE_3_7_FROM_BUNDLE_CODE                = "3_7_"
	// LC_LEARNER_FINAL_EXAM_ATTEMPTS_MAPPING_CODES = FILE_PREFIX + "3_64_"
)

func Get3_7FilePath(learnerId string) string {
	return DBROOTPATH + "/3_7FileGenerationFailedLearnerIds.txt"
}

// GetLearnerCourseActivityDirectoryPath Get Learner Local course Directory Path
func GetLearnerCourseActivityDirectoryPath(learnerID string) string {
	return DBROOTPATH + LEARNER_LOCAL_CREATE_FILES + learnerID + FILE_SEPARATOR + COURSE_ACTIVITY_DIRECTORY_NAME
}
func GetBundleDirectoryPath(learnerID string) string {
	return DBROOTPATH + LEARNER_LOCAL_CREATE_FILES + learnerID + FILE_SEPARATOR + BUNDLE_DIRECTORY_NAME + FILE_SEPARATOR
}
func GetCourseActivityDirPath(centerCode, learnerId string) string {
	return DBROOTPATH + FILE_SEPARATOR + LOCAL_CREATE_FILES + CENTER_CODE_ABBREVIATION + centerCode + FILE_SEPARATOR + LEARNER_DIRECTORY_NAME + learnerId + FILE_SEPARATOR + COURSE_ACTIVITY_DIRECTORY_NAME
}

func GetBundleFilePath(bundleName, learnerID, programID, eCourseID string) string {
	directoryPath := GetBundleDirectoryPath(learnerID)
	return directoryPath + bundleName + "_" + learnerID + "_" + programID + "_" + eCourseID + FILE_EXTENTION
}

// GetCourseActivityFilePath GetCourseActivityFilePath
func GetCourseActivityFilePath(courseActivityDirPath string, fileName string) string {
	return courseActivityDirPath + fileName
}

// GetLearnerAllocationDetailsFilePathByID Get Learner Allocation Details File Path by ID
func GetLearnerAllocationDetailsFilePathByID(learnerID string) string {
	return DBROOTPATH + LEARNER_LOCAL_CREATE_FILES + learnerID + "/" + LC_LEARNER_ALLOCATION_DETAILS_MAPPING_CODES + learnerID + FILE_EXTENTION
}

// GetLearnerCourseVisitFilePath Get Learner Course Visit File Path
func GetLearnerCourseVisitFilePath(centerCode, learnerId, programID, eCourseID, sessionID string) string {
	courseActivityPath := GetCourseActivityDirPath(centerCode, learnerId)
	return courseActivityPath + LC_LEARNER_COURSE_VISIT_MAPPING_CODES + learnerId + "_" + programID + "_" + eCourseID + "_" + sessionID + FILE_EXTENTION
}

// GetLearnerAllocationFileLCPath -
func GetLearnerAllocationFileLCPath(centerCode, learnerID string) string {
	learnerDirPath := GetLearnerDirLCPath(centerCode, learnerID)
	fileName := LC_LEARNER_ALLOCATION_DETAILS_MAPPING_CODES + learnerID
	return learnerDirPath + fileName + FILE_EXTENTION
}

// GetLearnerDirLCPath -
func GetLearnerDirLCPath(centerCode, learnerID string) string {
	centerPath := GetCenterLCPath(centerCode)
	return centerPath + LEARNER_DIRECTORY_NAME + learnerID + FILE_SEPARATOR
}

func GetCenterLCPath(centerCode string) string {
	return DBROOTPATH + FILE_SEPARATOR + LOCAL_CREATE_FILES + CENTER_CODE_ABBREVIATION + centerCode + FILE_SEPARATOR
}

// GetLearnerAssignmentDetailsFilePath GetLearnerAssignmentDetailsFilePath
func GetLearnerAssignmentDetailsFilePath(learnerID, centerCode, programID, eCourseID, batchID, sessionAggregationID string) string {
	courseActivityPath := GetCourseActivityDirPath(centerCode, learnerID)
	learnerAssignmentPath := courseActivityPath + LC_LEARNER_ASSIGNMENTS_MAPPING_CODES + learnerID + "_" + programID + "_" + eCourseID + "_" + batchID + "_" + sessionAggregationID + FILE_EXTENTION
	return learnerAssignmentPath
}

func FileAvailabilityCheck(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if fileInfo == nil && err != nil {
		return false
	}
	return true
}

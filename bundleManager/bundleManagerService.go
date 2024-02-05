package bundleManager

// import (
// 	"bytes"
// 	"path/filepath"
// 	"reverseUtilityNewERA/server/models"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/pquerna/ffjson/ffjson"

// 	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/dalhelper"
// 	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/filehelper"
// 	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
// 	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/securityhelper"
// )

// // BundleFile BundleFile
// type BundleFile struct {
// 	Data     interface{}    `json:"data"`
// 	Metadata BundleMetadata `json:"metadata"`
// }

// // BundleMetadata BundleMetadata
// type BundleMetadata struct {
// 	Xxhash   uint64    `json:"xxhash"`
// 	DateTime time.Time `json:"dateTime"`
// 	Version  string    `json:version"`
// }

// // GenerateBundle bundleName == 3_8|| 3_12 || 3_13 || 3_14 || 3_64
// func GenerateBundle(learnerID string, bundleName string) {
// 	courseActivityDirPath := models.GetLearnerCourseActivityDirectoryPath(learnerID)
// 	fileHelperObject := filehelper.FileHelperServiceObject{}
// 	fileList, searchError := fileHelperObject.FileSearch(bundleName+"*", courseActivityDirPath)

// 	if searchError != nil {
// 		logginghelper.LogError("error occured while searching :", bundleName, " file inside directory :", courseActivityDirPath, " ", searchError)
// 	}

// 	directoryPath := models.GetBundleDirectoryPath(learnerID)
// 	mapOfDataUpper := make(map[string](map[string]BundleFile))

// 	for _, filePath := range fileList {

// 		// bundleFileName := directoryPath + bundleName + "_" + learnerID + "_" + programID + "_" + eCourseID + model.FILE_EXTENTION
// 		bundleFileName := getBundleMapUpperKey(filepath.Base(filePath), learnerID) + models.FILE_EXTENTION

// 		mapOfData, upok := mapOfDataUpper[bundleFileName]
// 		if !upok {
// 			bundleFilePath := directoryPath + bundleFileName
// 			// check for is Bundle exist
// 			if models.FileAvailabilityCheck(bundleFilePath) {
// 				//  if exist read the bundle
// 				bundleData, getError := dalhelper.GetDataFromFDB(bundleFilePath)
// 				if getError != nil {
// 					logginghelper.LogError("error occured while fething data from  :", bundleFilePath, " ", getError)
// 				}
// 				unmarshalError := ffjson.Unmarshal(bundleData, &mapOfData)
// 				if unmarshalError != nil {
// 					logginghelper.LogError("error occured while unmarshaling data for file ", bundleFileName, unmarshalError)
// 					// create new bundle map if existing file is currupted
// 					mapOfData = make(map[string]BundleFile)
// 				}
// 			} else {
// 				mapOfData = make(map[string]BundleFile)
// 			}
// 		}
// 		// check for already exist
// 		fileName := filepath.Base(filePath)

// 		// get bundle upper map key name
// 		dataFromMap, ok := mapOfData[fileName]

// 		if ok {
// 			fileStat, fileStatErr := filehelper.FileInfo(filePath)

// 			if fileStatErr != nil {
// 				logginghelper.LogError("error occured while getting File Info ", fileStatErr)
// 			} else if dataFromMap.Metadata.DateTime == fileStat.ModTime() {
// 				continue
// 			}
// 		}

// 		learnerData, getError := dalhelper.GetDataFromFDB(filePath)
// 		if getError != nil {
// 			logginghelper.LogError("error occured while fething data from  :", filePath, " ", getError)
// 		}

// 		bundleFile := BundleFile{}
// 		bundleFile.Data = nil

// 		switch bundleName {
// 		case models.FILE_UPLOAD_3_8_FILE:

// 			object := models.LearnerVisit{}
// 			unmarshalError := ffjson.Unmarshal(learnerData, &object)
// 			if unmarshalError != nil {
// 				logginghelper.LogError("error occured while unmarshaling data for file ", filePath)
// 			} else {
// 				if object.ProgramID != "" {
// 					bundleFile.Data = object
// 				}
// 			}

// 		case models.FILE_UPLOAD_3_12_FILE:
// 			object := models.LearnerAssignments{}
// 			unmarshalError := ffjson.Unmarshal(learnerData, &object)
// 			if unmarshalError != nil {
// 				logginghelper.LogError("error occured while unmarshaling data for file ", filePath)
// 			} else {
// 				if object.ProgramID != "" {
// 					bundleFile.Data = object
// 				}
// 			}

// 		case models.FILE_UPLOAD_3_13_FILE:

// 			object := models.LearnerExamDetails{}
// 			unmarshalError := ffjson.Unmarshal(learnerData, &object)
// 			if unmarshalError != nil {
// 				logginghelper.LogError("error occured while unmarshaling data for file ", filePath)
// 			} else {
// 				if object.ProgramID != "" {
// 					bundleFile.Data = object
// 				}
// 			}
// 			//bundleFile.Data = object

// 		case models.FILE_UPLOAD_3_16_FILE:

// 			object := models.LearnerSPLExamDetails{}
// 			unmarshalError := ffjson.Unmarshal(learnerData, &object)
// 			if unmarshalError != nil {
// 				logginghelper.LogError("error occured while unmarshaling data for file ", filePath)
// 			} else {
// 				if object.ProgramID != "" {
// 					bundleFile.Data = object
// 				}
// 			}
// 		}
// 		//bundleFile.Data = object
// 		// case models.FILE_UPLOAD_3_14_FILE:
// 		// 	// umarshaling into both the structs as ffjson doesnt give unmarshal error
// 		// 	object := models.LearnerSingleQuestionDetails{}
// 		// 	unmarshalError := ffjson.Unmarshal(learnerData, &object)
// 		// 	if unmarshalError != nil {
// 		// 		logginghelper.LogError("error occured while unmarshaling data for file ", filePath)
// 		// 	} else {
// 		// 		// for old 3_14
// 		// 		if len(object.SingleQuestionAttemptDetails) > 0 {
// 		// 			if object.ProgramID != "" {
// 		// 				bundleFile.Data = object
// 		// 			}
// 		// 		} else {
// 		// 			// save 3_14 session aggregation wise
// 		// 			sessionAggregationWiseObject := models.LearnerSingleQuestionSessionAggregationWiseAttemptDetails{}
// 		// 			unmarshalError := ffjson.Unmarshal(learnerData, &sessionAggregationWiseObject)
// 		// 			if unmarshalError != nil {
// 		// 				logginghelper.LogError("error occured while unmarshaling data for file ", filePath)
// 		// 			} else {
// 		// 				if sessionAggregationWiseObject.ProgramID != "" {
// 		// 					bundleFile.Data = sessionAggregationWiseObject
// 		// 				}
// 		// 			}
// 		// 		}
// 		// 	}

// 		// case models.FILE_UPLOAD_3_64_FILE:
// 		// 	object := models.LearnerExamDetails{}
// 		// 	unmarshalError := ffjson.Unmarshal(learnerData, &object)
// 		// 	if unmarshalError != nil {
// 		// 		logginghelper.LogError("error occured while unmarshaling data for file ", filePath)
// 		// 	} else {
// 		// 		if len(object.ExamAttemptDetails) > 0 {
// 		// 			if object.ProgramID != "" {
// 		// 				bundleFile.Data = object
// 		// 			}
// 		// 		} else {
// 		// 			learnerExamLogObj := models.LearnerExamAttemptLogDetails{}
// 		// 			unmarshalError := ffjson.Unmarshal(learnerData, &learnerExamLogObj)
// 		// 			if unmarshalError != nil {
// 		// 				logginghelper.LogError("error occured while unmarshaling data for file ", filePath)
// 		// 			} else {
// 		// 				if learnerExamLogObj.ProgramID != "" {
// 		// 					bundleFile.Data = learnerExamLogObj
// 		// 				}
// 		// 			}
// 		// 		}
// 		// 	}
// 		// }

// 		// if bundleFile.Data do not contains any value i.e. nil, continue
// 		if bundleFile.Data == nil {
// 			continue
// 		}
// 		// Old hash code
// 		// hash, hashError := securityhelper.GetHashChecksumOfByteArray(learnerData)

// 		// New hash code
// 		hash, hashError := securityhelper.GetAttributeBasedHashWithRoundedModTime(filePath)

// 		if hashError != nil {
// 			logginghelper.LogError("error occured while hash of data ", hashError)
// 		} else {
// 			hashVal, _ := strconv.ParseUint(hash, 10, 64)
// 			bundleFile.Metadata.Xxhash = hashVal
// 		}

// 		fileStat, fileStatErr := filehelper.FileInfo(filePath)
// 		if fileStatErr != nil {
// 			logginghelper.LogError("error occured while getting File Info ", fileStatErr)
// 		}

// 		bundleFile.Metadata.DateTime = fileStat.ModTime()
// 		bundleFile.Metadata.Version = "1"

// 		mapKey := filepath.Base(filePath)
// 		if mapOfData != nil {
// 			mapOfData[mapKey] = bundleFile
// 			if mapOfDataUpper != nil {
// 				mapOfDataUpper[bundleFileName] = mapOfData
// 			}
// 		} else {
// 			continue
// 		}
// 	}
// 	if len(mapOfDataUpper) > 0 {
// 		directoryPath := models.GetBundleDirectoryPath(learnerID)

// 		for k, v := range mapOfDataUpper {

// 			// save 3_7 to bundle
// 			// get 3_7 from Lazy Writter Cache and save in to 3_8 bundle Map
// 			switch bundleName {
// 			case models.FILE_UPLOAD_3_8_FILE:

// 				bundleFile3_7 := BundleFile{}
// 				bundleFile3_7.Data = nil
// 				// get validate3_7 from LearnerId folder that we created 3_7 from mongo
// 				learnerAllocation, getError := GetLearnerAllocationDetailsDAO(learnerID)
// 				if getError != nil {
// 					logginghelper.LogError("error occured while getting data from cache", getError)
// 				}
// 				fileName3_7 := models.GetLearnerAllocationDetailsFilePathByID(learnerID)

// 				bundleFile3_7.Data = learnerAllocation

// 				// Old hash code
// 				// learnerDataByteArray, marshalError := ffjson.Marshal(learnerAllocation)
// 				// if marshalError != nil {
// 				// 	logginghelper.LogError("error occured while marshalling data ", marshalError)
// 				// }
// 				// hash, hashError := securityhelper.GetHashChecksumOfByteArray(learnerDataByteArray)

// 				// New hash code
// 				hash, hashError := securityhelper.GetAttributeBasedHashWithRoundedModTime(fileName3_7)

// 				if hashError != nil {
// 					logginghelper.LogError("error occured while hash of data ", hashError)
// 				} else {
// 					hashVal, _ := strconv.ParseUint(hash, 10, 64)
// 					bundleFile3_7.Metadata.Xxhash = hashVal
// 				}

// 				bundleFile3_7.Metadata.DateTime = time.Now()
// 				bundleFile3_7.Metadata.Version = "1"

// 				mapKey3_7 := filepath.Base(fileName3_7)
// 				if v != nil {
// 					v[mapKey3_7] = bundleFile3_7
// 				} else {
// 					continue
// 				}
// 			}
// 			objectBytes, marshalError := ffjson.Marshal(v)

// 			if marshalError != nil {
// 				logginghelper.LogError("error occured while marshaling data for file ")
// 				continue
// 			}

// 			saveError := dalhelper.SaveDataToFDB(directoryPath+"/"+k, objectBytes, true)
// 			if saveError != nil {
// 				logginghelper.LogError("error occured while saving data to file ", k)
// 			}
// 		}
// 	}
// }

// func getBundleMapUpperKey(filePath, learnerID string) string {
// 	// File extension check added to access only json files
// 	if filePath != "" && len(filePath) > 0 && strings.HasSuffix(filePath, models.FILE_EXTENTION) {
// 		var mapProgramCourseKey bytes.Buffer

// 		str := strings.Split(strings.Split(filePath, ".json")[0], "_")
// 		// Length check is added to avoid panic while reading course activity files
// 		if len(str) < 5 {
// 			logginghelper.LogError("getBundleMapUpperKey Error: filepath length is less than 5 with filepath: ", filePath)
// 			return ""
// 		}
// 		mapProgramCourseKey.WriteString(str[0] + "_" + str[1] + "_" + str[2] + "_" + str[3] + "_" + str[4])
// 		return mapProgramCourseKey.String()
// 	}
// 	return ""
// }

// // CrossCheckFileIsNotAvailable CrossCheckFileIsNotAvailable
// func CrossCheckFileIsNotAvailable(filePath string) bool {
// 	if !models.FileAvailabilityCheck(filePath) {
// 		return false
// 	}
// 	return true
// }

// func checkIfFileIsPresent(fileName string, existingFileList []string) bool {

// 	for index := 0; index < len(existingFileList); index++ {
// 		existingFileName := filepath.Base(existingFileList[index])

// 		if fileName == existingFileName {

// 			return true
// 		}
// 	}
// 	return false
// }

// // DeleteFileDataFromBundle deletes desired file data from bundle and writes map again to bundle directory
// func DeleteFileDataFromBundle(learnerID string, programID string, eCourseID string, fileName string, bundleName string) {
// 	directoryPath := models.GetBundleDirectoryPath(learnerID)

// 	mapOfData := make(map[string]BundleFile)

// 	// bndleFilePath := directoryPath + bundleName + model.FILE_EXTENTION
// 	// bndleFilePath := model.GetBundleFilePath(learnerID, directoryPath, bundleName)
// 	bndleFilePath := directoryPath + "/" + getBundleMapUpperKey(filepath.Base(fileName), learnerID) + models.FILE_EXTENTION

// 	bundleData, getError := dalhelper.GetDataFromFDB(bndleFilePath)
// 	if getError != nil {
// 		logginghelper.LogError("error occured while fething data from  :", bndleFilePath, " ", getError)
// 	}

// 	unmarshalError := ffjson.Unmarshal(bundleData, &mapOfData)
// 	if unmarshalError != nil {
// 		logginghelper.LogError("error occured while unmarshaling data for file ", bndleFilePath)
// 	}

// 	delete(mapOfData, fileName)

// 	objectBytes, marshalError := ffjson.Marshal(mapOfData)
// 	if marshalError != nil {
// 		logginghelper.LogError("error occured while marshaling data for file ")
// 	}

// 	// bundleFileName := directoryPath + bundleName + model.FILE_EXTENTION
// 	bundleFileName := models.GetBundleFilePath(bundleName, learnerID, programID, eCourseID)
// 	saveError := dalhelper.SaveDataToFDB(bundleFileName, objectBytes, true)
// 	if saveError != nil {
// 		logginghelper.LogError("error occured while saving data to file ", bundleFileName)
// 	}
// }

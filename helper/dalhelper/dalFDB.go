package dalhelper

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/securityhelper"
	linq "github.com/ahmetb/go-linq"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/filehelper"

	"time"
)

// ============================== Configuration =====================================
const (
	READ_ONLY = "RO"
)

var (
	NumberOfReads                 = 0
	NumberOfWrites                = 0
	NumberOfReadsWithoutSecurity  = 0
	NumberOfWritesWithoutSecurity = 0
)

var securityRequired = false
var defaultSecurityKey = []byte{}

type FDBModelType struct {
	ModelFilePath string
	ModelType     string // it can be RO,WO,UO
}

var fileAndArrayMap map[string][]byte
var hostAddressForDataFiles = ""
var splitString = "ERA_DB/"
var fDBModelTypeArray []FDBModelType
var mutex = &sync.Mutex{}

// SetFDBDownloadConfig Set remote path for receiving files if not found at local
func SetFDBFileAndArrayMap(fileNameArray []string) {
	if fileNameArray != nil && len(fileNameArray) > 0 {
		for _, fileName := range fileNameArray {
			fileAndArrayMap[fileName] = make([]byte, 1)
		}
	}
	//	fileAndArrayMap = fileAndArray
}

// SetFDBDownloadConfig Set remote path for receiving files if not found at local
func SetFDBDownloadConfig(remoteHost string, splitbystring string) {
	hostAddressForDataFiles = remoteHost
	splitString = splitbystring
}

// SetFDBDownloadConfig Set remote path for receiving files if not found at local
func SetFDBModelTypeArray(FDBModelTypeArray []FDBModelType) {
	fDBModelTypeArray = FDBModelTypeArray
}

// EnableSecurity Set remote path for receiving files if not found at local
func EnableSecurity(key []byte, initializationVector string) {
	securityhelper.SetSecurityConfig(key, initializationVector)
	defaultSecurityKey = key
	securityRequired = true
}

// DisableSecurity disables security
func DisableSecurity() {
	securityRequired = false
}

// GetSecurityStatus get status of security flag
func GetSecurityStatus() bool {
	return securityRequired
}

//============================== Configuration =====================================

func CheckModelTypeOfRequestedModel(filePath string) FDBModelType {
	searchedFDBArray := []FDBModelType{}
	linq.From(fDBModelTypeArray).
		WhereT(
			func(modelType FDBModelType) bool {
				return modelType.ModelFilePath == filePath
			},
		).
		ToSlice(&searchedFDBArray)

	if searchedFDBArray != nil && len(searchedFDBArray) < 0 {
		return searchedFDBArray[0]
	}
	return FDBModelType{}

}

// GetDataFromFDB gets data from FDB
func GetDataFromFDB(filePath string) ([]byte, error) {
	//Following checking will be done for the provided file path, which will let us know
	// whether the file is Read only/Write only/Update only
	mutex.Lock()
	NumberOfReads = NumberOfReads + 1
	mutex.Unlock()

	// if isDebugMode {
	// 	lazyCacheObject := PerformanceAnalyser[filePath]
	// 	lazyCacheObject.DISK_READ_COUNT++
	// 	PerformanceAnalyser[filePath] = lazyCacheObject
	// }
	data, err := filehelper.ReadFile(filePath)
	if err != nil {
		if hostAddressForDataFiles == "" {
			logginghelper.LogError("Call dalhelper.SetRemoteHostPath(remoteHost) to set remote file path ")
			return nil, err
		}
		relativeFilePath := strings.SplitAfter(filePath, splitString)
		remotePath := hostAddressForDataFiles + relativeFilePath[1]
		resp, httpError := http.Get(remotePath)
		if httpError != nil || resp.StatusCode == 404 {

			logginghelper.LogError("Error while receiving data from ", remotePath, " ", httpError, " Status code :", resp.StatusCode)
			return nil, httpError
		}
		data, readError := ioutil.ReadAll(resp.Body)
		if readError != nil {
			logginghelper.LogError("Error while reading data from response body ", readError)
			return nil, readError
		}
		defer resp.Body.Close()
		saveError := SaveDataToFDB(filePath, data, true)
		if saveError != nil {
			logginghelper.LogError("Error while saving data to ", filePath, " ", saveError)
			return nil, saveError
		}

	}

	if securityRequired {
		// decryptedData, decryptionError := securityhelper.AESDecryptDefault(data)
		keyBytes := GetKeyWithFileNameAndDefaultKey(filePath)
		decryptedData, decryptionError := securityhelper.AESDecrypt(data, keyBytes)
		if decryptionError != nil {
			logginghelper.LogError("Error while decrypting data of ", filePath, " : ", decryptionError)
			return nil, decryptionError
		}
		decompressedData, decompressError := securityhelper.Decompress(decryptedData)
		if decompressError != nil {
			logginghelper.LogError("Error while decompressing data of ", filePath, " : ", decompressError)
			return nil, decompressError
		}
		return decompressedData, nil
	}
	return data, err
}

// SaveDataToFDB saves data to FDB
func SaveDataToFDB(filePath string, data []byte, makeDir bool) error {
	mutex.Lock()
	NumberOfWrites = NumberOfWrites + 1
	mutex.Unlock()
	// if isDebugMode {
	// 	lazyCacheObject := PerformanceAnalyser[filePath]
	// 	lazyCacheObject.DISK_WRITE_COUNT++
	// 	PerformanceAnalyser[filePath] = lazyCacheObject
	// }
	if makeDir {
		os.MkdirAll(path.Dir(filePath), os.ModePerm)
	}

	if securityRequired {
		// encryptedData, encryptionError := securityhelper.AESEncryptDefault(data)

		keyBytes := GetKeyWithFileNameAndDefaultKey(filePath)

		compressedText, compressionError := securityhelper.Compress(data)
		if compressionError != nil {
			logginghelper.LogError("Error while compressing data of ", filePath, " : ", compressionError)
			return compressionError
		}
		encryptedData, encryptionError := securityhelper.AESEncrypt(compressedText, keyBytes)
		if encryptionError != nil {
			logginghelper.LogError("Error while encrypting data of ", filePath, " : ", encryptionError)
			return encryptionError
		}
		fileWriteError := filehelper.WriteFile(filePath, encryptedData)
		if fileWriteError != nil {
			logginghelper.LogError("Error while writing data to ", filePath, " : ", fileWriteError)
		}
		return fileWriteError
	}

	err := filehelper.WriteFile(filePath, data)
	if err != nil {
		logginghelper.LogError("Error while writing data to ", filePath, " : ", err)
	}
	return err
}

// GetKeyWithFileNameAndDefaultKey generates key using file name + Default key
func GetKeyWithFileNameAndDefaultKey(filePath string) []byte {
	fileName := filepath.Base(filePath)
	fileNameBytes := []byte(fileName)
	fileNameBytes = append(fileNameBytes, defaultSecurityKey...)
	keyBytes := securityhelper.Get128BitHash(fileNameBytes)
	return keyBytes[:]
}

// AppendDataToFDB apppends data to FDB
func AppendDataToFDB(filePath string, data []byte) error {
	count, err := filehelper.AppendFile(filePath, string(data))
	logginghelper.LogInfo("Count returned from append function ", count)
	return err
}

// DeleteFileFromFDB deletes file to FDB
// BUG: Remove time.now as it add :: into file name which do not work in windows
func SoftDeleteFileFromFDB(filePath string) error {
	t := time.Now()

	newFilePath := filePath + "_deleted_" + t.Format("20060102150405")
	err := filehelper.RenameFile(filePath, newFilePath)

	return err

}

// HardDeleteFileFromFDB
func HardDeleteFileFromFDB(filePath string) error {
	// newFilePath := filePath + "_deleted_" + time.Now().String()
	err := filehelper.DeleteFile(filePath)
	return err
}

// GetDataFromFDBWithoutSecurity gets data from FDB
func GetDataFromFDBWithoutSecurity(filePath string) ([]byte, error) {
	//Following checking will be done for the provided file path, which will let us know
	// whether the file is Read only/Write only/Update only
	mutex.Lock()
	NumberOfReadsWithoutSecurity = NumberOfReadsWithoutSecurity + 1
	mutex.Unlock()
	data, err := filehelper.ReadFile(filePath)
	if err != nil {
		if hostAddressForDataFiles == "" {
			logginghelper.LogError("Call dalhelper.SetRemoteHostPath(remoteHost) to set remote file path ")
			return nil, err
		}
		relativeFilePath := strings.SplitAfter(filePath, splitString)
		remotePath := hostAddressForDataFiles + relativeFilePath[1]
		resp, httpError := http.Get(remotePath)
		if httpError != nil || resp.StatusCode == 404 {

			logginghelper.LogError("Error while receiving data from ", remotePath, " ", httpError, " Status code :", resp.StatusCode)
			return nil, httpError
		}
		data, readError := ioutil.ReadAll(resp.Body)
		if readError != nil {
			logginghelper.LogError("Error while reading data from response body ", readError)
			return nil, readError
		}
		defer resp.Body.Close()
		saveError := SaveDataToFDB(filePath, data, true)
		if saveError != nil {
			logginghelper.LogError("Error while saving data to ", filePath, " ", saveError)
			return nil, saveError
		}

	}

	return data, err
}

// SaveDataToFDBWithoutSecurity saves data to FDB
func SaveDataToFDBWithoutSecurity(filePath string, data []byte, makeDir bool) error {
	mutex.Lock()
	NumberOfWritesWithoutSecurity = NumberOfWritesWithoutSecurity + 1
	mutex.Unlock()
	if makeDir {
		os.MkdirAll(path.Dir(filePath), os.ModePerm)
	}

	err := filehelper.WriteFile(filePath, data)
	return err
}

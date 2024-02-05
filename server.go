package main

import (
	"context"
	"fmt"
	"github/revert_utility_era2/db"
	"github/revert_utility_era2/helper/dalhelper"
	"github/revert_utility_era2/migration"
	"net/url"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

var (
	SecurityKey = "MKCLSecurity$#@!"
	IV          = "AAAAAAAAAAAAAAAA"
)

// J3a@Wg_vsL6e9y$$ // bsdm production
// q{$Kg#+8J(j^x^q8 // mkcl production
// J&BDqYhW?5$nh7p6 // hkcl production
// tm#XY7E]-Cx!%8c3 // homedemy production
// BSDMSecurity$#@! // bsdm staging
// MKCLSecurity$#@! // mkcl stagingSSS
// HOMEDEMYSecurity$#@! // homedemy staging

func main() {
	fmt.Println("Application Started")
	dalhelper.EnableSecurity([]byte(SecurityKey), IV)
	var hostname, username, password, dbName string
	hostname = "10.15.20.80:27017"
	username = "reader"
	password = url.QueryEscape("reader#123")
	dbName = "ERALive_2"

	if err := db.InitMongoDB(hostname, username, password, dbName); err != nil {
		logginghelper.LogFatal(err)
		return
	}

	LearnerID := "2024302"
	_, err := migration.ProcessLearner3_7File(context.Background(), LearnerID)
	if err != nil {
		logginghelper.LogError(err)
		return
	}

	fmt.Println("Application Stop")
}

// func WriteData(learnerID string, data interface{}) error {
// 	filePath := models.GetLearnerAllocationDetailsFilePathByID(learnerID)
// 	filehelper.CreateDirectoryRecursive(filePath)
// 	// filePath := fmt.Sprintf("./3_7_%s.json", learnerID)
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	// Create and open the file for writing
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// Write the JSON data to the file
// 	_, err = file.Write(jsonData)
// 	if err != nil {
// 		return err
// 	}

// 	// Optionally, you can also output a message or log
// 	// to confirm the file has been written successfully
// 	logginghelper.LogInfo("JSON data written to", filePath)
// 	return nil
// }

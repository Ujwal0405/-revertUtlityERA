package main

import (
	"context"
	"fmt"
	"github/revert_utility_era2/db"
	"github/revert_utility_era2/helper/dalhelper"
	"github/revert_utility_era2/migration"
	"net/url"
	"sync"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

var (
	SecurityKey      = "MKCLSecurity$#@!"
	IV               = "AAAAAAAAAAAAAAAA"
	MapMutex         = &sync.Mutex{}
	LearnerCenterMap map[string]string
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
	hostname = "10.15.20.251:27017"
	username = "reader"
	password = url.QueryEscape("reader#123")
	dbName = "ERALive_2"

	if err := db.InitMongoDB(hostname, username, password, dbName); err != nil {
		logginghelper.LogFatal(err)
		return
	}

	LearnerID := []string{"12121212"}

	for _, lId := range LearnerID {
		err := PrepareLearnerMetabse(lId)
		if err != nil {
			logginghelper.LogError(err)
			return
		}
	}

	// _, err := migration.ProcessLearner3_7File(context.Background(), LearnerID)
	// if err != nil {
	// 	logginghelper.LogError(err)
	// 	return
	// }
	// _, err = migration.ProcessLearner3_12File(context.Background(), LearnerID)
	// if err != nil {
	// 	logginghelper.LogError(err)
	// 	return
	// }

	fmt.Println("Application Stop")
}

func PrepareLearnerMetabse(lId string) error {

	learnerBasicData, err := migration.ProcessLearner3_7File(context.Background(), lId)
	if err != nil {
		logginghelper.LogError(err)
		return err
	}
	_, err = migration.ProcessLearner3_12File(context.Background(), lId, learnerBasicData.CenterCode)
	if err != nil {
		logginghelper.LogError(err)
		return err
	}
	if err := migration.ProcessLearner3_8Files(context.Background(), lId, learnerBasicData.CenterCode); err != nil {
		logginghelper.LogError(err)
		return err
	}
	return nil
}

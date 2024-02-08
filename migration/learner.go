package migration

import (
	"context"
	"github/revert_utility_era2/db"
	"github/revert_utility_era2/models"
	"github/revert_utility_era2/utils"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"go.mongodb.org/mongo-driver/bson"
)

type LearnerDetails struct {
	LearnerId  string `json:"learnerId"`
	CenterCode string `json:"centerCode"`
}

func ProcessLearner3_7File(ctx context.Context, learnerID string) (models.LearnerBasicDtls, error) {
	logginghelper.LogDebug("IN: ProcessLearner3_7File")
	learnerAllocation := models.LearnerAllocation{}

	centerCode := ""
	learnerDetails, err := getLearnerDetails(ctx, learnerID)
	if err != nil {
		return models.LearnerBasicDtls{}, err
	}

	learnerAllocation.LearnerDetails = learnerDetails

	allocations, err := getLearnerAllocations(ctx, learnerID)
	if err != nil {
		return models.LearnerBasicDtls{}, err
	}

	if len(allocations) > 0 {
		centerCode = allocations[0].LearningCenterCode
	}

	learnerAllocation.AllocationDetails = allocations

	filePath := models.GetLearnerAllocationFileLCPath(centerCode, learnerID)
	err = utils.WriteData(filePath, learnerAllocation)

	if err != nil {
		return models.LearnerBasicDtls{}, err
	}
	logginghelper.LogDebug("OUT: ProcessLearner3_7File")
	return models.LearnerBasicDtls{
		LearnerID:  learnerID,
		CenterCode: centerCode,
	}, nil
}

func getLearnerDetails(ctx context.Context, learnerID string) (models.LearnerDetails, error) {

	filter := bson.M{
		"learnerId": learnerID,
	}
	ld := models.MongoLearnerDetails{}
	// err := db.ERALiveDB.Collection(LearnerSessionCompletionDetailsCollection).FindOne(ctx, filter).Decode(&ld)
	err := db.ERALiveDB.Collection(models.LEARNING_SESSION_COMPLETION_DETAILS_COLLECTION_NAME).FindOne(ctx, filter).Decode(&ld)

	if err != nil {
		return models.LearnerDetails{}, err
	}

	return models.LearnerDetails{
		LearnerID: ld.LearnerId,
		PhotoPath: ld.PhotoPath,
	}, nil
}

func getLearnerCourseAllocationFromMongo(ctx context.Context, learnerID string) ([]models.MongoLearnerCourceAllocation, error) {
	filter := bson.M{
		"learnerId": learnerID,
	}

	courceAllocationscursor, err := db.ERALiveDB.Collection(models.LEARNER_COURSE_ALLOCATION_COLLECTION_NAME).Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer courceAllocationscursor.Close(ctx)

	mcourceAllocations := []models.MongoLearnerCourceAllocation{}

	if err := courceAllocationscursor.All(ctx, &mcourceAllocations); err != nil {
		return nil, err
	}

	return mcourceAllocations, nil
}

func getSessionCompletionDetailsFromMongo(ctx context.Context, allocation models.MongoLearnerCourceAllocation) ([]models.MongoSessionCompletionDetails, error) {
	filter := bson.M{
		"learnerId":        allocation.LearnerId,
		"programId":        allocation.ProgramId,
		"eCourseId":        allocation.ECourseId,
		"eCoursePatternId": allocation.ECoursePatternId,
	}
	logginghelper.LogInfo("filter:", filter)
	cursor, err := db.ERALiveDB.Collection(models.LEARNING_SESSION_COMPLETION_DETAILS_COLLECTION_NAME).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	sessionCompletionDetails := []models.MongoSessionCompletionDetails{}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &sessionCompletionDetails); err != nil {
		return nil, err
	}

	return sessionCompletionDetails, nil
}

func getLearnerAllocations(ctx context.Context, learnerID string) ([]models.AllocationDetails, error) {

	allocations := []models.AllocationDetails{}
	mallocations, err := getLearnerCourseAllocationFromMongo(ctx, learnerID)

	if err != nil {
		return nil, err
	}

	for _, al := range mallocations {
		sessionCompletionDetails, err := getSessionCompletionDetailsFromMongo(ctx, al)
		if err != nil {
			logginghelper.LogError(err)
			continue
		}
		logginghelper.LogInfo("session Completion details:", len(sessionCompletionDetails))
		allocations = append(allocations, prepareLearnerAllocationObject(al, sessionCompletionDetails))
	}

	return allocations, nil
}

func prepareLearnerAllocationObject(mallocation models.MongoLearnerCourceAllocation, msessions []models.MongoSessionCompletionDetails) models.AllocationDetails {
	allocationDetails := models.AllocationDetails{
		LearnerCourseID:                    mallocation.LearnerCourseId,
		LearningCenterID:                   mallocation.LearningCenterId,
		LearningCenterCode:                 mallocation.LearningCenterCode,
		LearningCenterName:                 mallocation.LearningCenterName,
		BatchID:                            utils.ToString(mallocation.ConfirmationMonth + mallocation.ConfirmationYear),
		BatchName:                          mallocation.BatchName,
		ProgramID:                          mallocation.ProgramId,
		ProgramName:                        mallocation.ProgramName,
		ECourseID:                          mallocation.ECourseId,
		ECourseName:                        mallocation.ECourseName,
		ECourseImage:                       mallocation.ECourseImage,
		ECourseVersion:                     mallocation.ECourseVersion,
		LearningStartDate:                  utils.TimestampToTime(mallocation.LearningStartDate),
		LearnerCoursePrefLanguageCode:      mallocation.LearnerCoursePrefLanguageCode,
		CompletionPercentage:               mallocation.LearningDetails.CompletionPercentage,
		TimeSpent:                          mallocation.LearningDetails.TimeSpent,
		CompletedAggregationCount:          mallocation.LearningDetails.CompletedAggregationCount,
		SolvedAssignmentCount:              mallocation.LearningDetails.SolvedAssignmentCount,
		CompletedSessionCount:              mallocation.LearningDetails.CompletedSessionCount,
		CurrentSession:                     mallocation.LearningDetails.CurrentSessionId,
		CurrentSessionAggregationID:        mallocation.LearningDetails.CurrentSessionId,
		AggregationID:                      mallocation.LearningDetails.CurrentNodeAggregationId,
		NextSequentialAggregationID:        mallocation.LearningDetails.NextSequentialAggregationId,
		NextSequentialSessionAggregationID: mallocation.LearningDetails.NextSequentialSessionId,
		VisitedTime:                        utils.TimestampToTime(mallocation.LearningDetails.VisitedTime),
		CoursePatternID:                    mallocation.ECoursePatternId,
		LearningScoreDetails:               prepareLearningScoreDetails(mallocation.LearningDetails.LearningScoreDetails),
		SessionCompletionDetails:           prepareLearnerSessionCompletionDetails(msessions),
	}
	return allocationDetails
}

func prepareLearnerSessionCompletionDetails(msessions []models.MongoSessionCompletionDetails) []models.SessionCompletionDetails {
	sessions := []models.SessionCompletionDetails{}

	for _, ms := range msessions {
		sessions = append(sessions,
			models.SessionCompletionDetails{
				AggregationID:            ms.AggregationId,
				CompletedOn:              utils.TimestampToTime(ms.CompletedOn),
				IsCompleted:              ms.IsCompleted,
				SessionStatus:            ms.SessionStatus,
				TotalAggregationDuration: ms.TotalAggregationDuration,
				LastUpdateDateTime:       utils.TimestampToTime(ms.LastUpdateDateTime),
				PassingHeadWiseMarks:     preparePassingHeadWiseMakrsDetails(ms.PassingHeadWiseMarks),
			},
		)
	}
	return sessions
}

func preparePassingHeadWiseMakrsDetails(mpassingHeads []models.MongoPassingHeadWiseMark) []models.LearningScoreDetails {
	passingHeads := []models.LearningScoreDetails{}

	for _, mph := range mpassingHeads {
		passingHeads = append(passingHeads,
			models.LearningScoreDetails{
				PassingHeadCode: mph.PassingHeadCode,
				Score:           mph.Score,
				ScoreDetails: models.ScoreDetails{
					TotalScore:      mph.ScoreDetails.TotalScore,
					AdditionalMarks: mph.ScoreDetails.AdditionalMarks,
				},
			},
		)
	}
	return passingHeads
}

func prepareLearningScoreDetails(learningScoreDetails []models.MongoLearningScoreDetails) []models.LearningScoreDetails {

	scds := []models.LearningScoreDetails{}

	for _, ls := range learningScoreDetails {
		scds = append(scds,
			models.LearningScoreDetails{
				PassingHeadCode:        ls.PassingHeadCode,
				PassingHeadName:        ls.PassingHeadName,
				PassingHeadDisplayName: ls.PassingHeadDisplayName,
				Score:                  ls.Score,
				ScoreDetails: models.ScoreDetails{
					TotalScore:      ls.ScoreDetails.TotalScore,
					AdditionalMarks: ls.ScoreDetails.AdditionalMarks,
				},
				MaxMarks: ls.MaxMarks,
			},
		)
	}
	return scds
}

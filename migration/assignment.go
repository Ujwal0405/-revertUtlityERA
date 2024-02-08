package migration

import (
	"context"
	"github/revert_utility_era2/db"
	"github/revert_utility_era2/models"
	"github/revert_utility_era2/utils"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"go.mongodb.org/mongo-driver/bson"
)

func ProcessLearner3_12File(ctx context.Context, learnerId, centerCode string) (models.LearnerAssignments, error) {
	logginghelper.LogDebug("IN : ProcessLearner3_12File")
	learnerAssignments := models.LearnerAssignments{}

	assignments, err := getLearnerAssignments(ctx, learnerId)
	if err != nil {
		return learnerAssignments, err
	}

	for _, learnerAssignments := range assignments {
		filePath := models.GetLearnerAssignmentDetailsFilePath(learnerAssignments.LearnerID, centerCode, learnerAssignments.ProgramID, learnerAssignments.ECourseID, learnerAssignments.BatchID, learnerAssignments.SessionAggregationID)
		err = utils.WriteData(filePath, learnerAssignments)

		if err != nil {
			return learnerAssignments, err
		}
	}

	return learnerAssignments, nil
}

func getLearnerAssignments(ctx context.Context, learnerId string) ([]models.LearnerAssignments, error) {
	learnerAssignments := []models.LearnerAssignments{}
	mAssignments, err := getLearnerAssignmentsFromMongo(ctx, learnerId)

	if err != nil {
		return nil, err
	}
	filters := filterAssignmentBySessionId(mAssignments)
	la := models.LearnerAssignments{}
	for _, ads := range filters {
		for _, ad := range ads {

			la.LearnerID = ad.LearnerId
			la.ProgramID = ad.ProgramId
			la.ECourseID = ad.ECourseId
			la.CoursePatternID = ad.ECoursePatternId
			la.SessionAggregationID = ad.SessionAggregationId
			la.BatchID = utils.ToString(ad.ConfirmationMonth + ad.ConfirmationYear)

			logginghelper.LogInfo("assignment Details:", len(ads))
			la.AssignmentsDetails = append(la.AssignmentsDetails, prepareLearnerAssignmentObject(ad))
		}
		learnerAssignments = append(learnerAssignments, la)

	}
	return learnerAssignments, nil
}
func filterAssignmentBySessionId(mAssignments []models.MongoAssignmentSubmission) map[string][]models.MongoAssignmentSubmission {
	filters := map[string][]models.MongoAssignmentSubmission{}
	for _, mAd := range mAssignments {
		_, ok := filters[mAd.AssignmentId]
		if !ok {
			filters[mAd.AssignmentId] = []models.MongoAssignmentSubmission{}
		}
		filters[mAd.AssignmentId] = append(filters[mAd.AssignmentId], mAd)

	}
	return filters
}
func getLearnerAssignmentsFromMongo(ctx context.Context, learnerId string) ([]models.MongoAssignmentSubmission, error) {
	filter := bson.M{
		"learnerId": learnerId,
	}
	assignmentCursor, err := db.ERALiveDB.Collection(models.LEARNER_ASSIGNMENT_DETAILS_COLLECTION_NAME).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer assignmentCursor.Close(ctx)
	massignments := []models.MongoAssignmentSubmission{}
	if err := assignmentCursor.All(ctx, &massignments); err != nil {
		return nil, err
	}
	return massignments, nil
}

func prepareLearnerAssignmentObject(mAd models.MongoAssignmentSubmission) models.AssignmentsDetails {
	return models.AssignmentsDetails{
		AssignmentID:             mAd.AssignmentId,
		AttemptNo:                mAd.AttemptNo,
		AttemptStartDate:         *utils.TimestampToTime(mAd.AttemptStartDate),
		AttemptEndDate:           *utils.TimestampToTime(mAd.AttemptEndDate),
		UploadedRelativeFilePath: mAd.UploadedRelativeFilePath,
		MarksObtained:            mAd.MarksObtained,
		LanguageID:               mAd.LanguageId,
		AssignmentStatus:         mAd.AssignmentStatus,
		IsFileUploaded:           mAd.IsFileUploaded,
		EvaluationDetails:        mAd.EvaluationDetails,
	}
}

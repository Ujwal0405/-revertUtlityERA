package migration

import (
	"context"
	"github/revert_utility_era2/db"
	"github/revert_utility_era2/models"
	"github/revert_utility_era2/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func ProcessLearner3_8Files(ctx context.Context, learnerID, centerCode string) error {
	// Construct an aggregation pipeline to group sessions by ID
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"learnerId": learnerID},
		},
		bson.M{
			"$group": bson.M{
				"_id":     "$sessionId",
				"details": bson.M{"$first": "$$ROOT"},
			},
		},
	}

	// Execute the aggregation pipeline
	cursor, err := db.ERALiveDB.Collection(models.LEARNER_SESSION_VISIT_DETAILS_COLLECTION_NAME).Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Iterate over the sessions in the aggregation result
	for cursor.Next(ctx) {
		var result struct {
			ID      string              `bson:"_id"`
			Details models.LearnerVisit `bson:"details"`
		}

		if err := cursor.Decode(&result); err != nil {
			return err
		}

		// Fetch all visit details for the current session
		visitDetails, err := fetchVisitDetailsForSession(ctx, result.Details)
		if err != nil {
			return err
		}

		// Assign the fetched visit details to the current session
		result.Details.VisitsDetail = visitDetails

		// Create a file path for the JSON file based on the session ID
		filePath := models.GetLearnerCourseVisitFilePath(centerCode, learnerID, result.Details.ProgramID, result.Details.ECourseID, result.Details.SessionID)
		// Write the session details to the JSON file
		err = utils.WriteData(filePath, result.Details)
		if err != nil {
			return err
		}
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}

func fetchVisitDetailsForSession(ctx context.Context, session models.LearnerVisit) ([]models.VisitsDetail, error) {
	filter := bson.M{
		"learnerId": session.LearnerID,
		"programId": session.ProgramID,
		"eCourseId": session.ECourseID,
		"sessionId": session.SessionID,
	}

	// Fetch all visit details for the session
	cursor, err := db.ERALiveDB.Collection(models.LEARNER_SESSION_VISIT_DETAILS_COLLECTION_NAME).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var visitDetails []models.VisitsDetail

	// Iterate over the results and append each visit detail to visitDetails
	for cursor.Next(ctx) {
		var visitDetail models.VisitsDetail
		if err := cursor.Decode(&visitDetail); err != nil {
			return nil, err
		}
		visitDetails = append(visitDetails, visitDetail)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return visitDetails, nil
}

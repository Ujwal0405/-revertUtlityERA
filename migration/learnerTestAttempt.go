package migration

// package migration

// import (S
// 	"context"
// 	"github/revert_utility_era2/db"
// 	"github/revert_utility_era2/models"
// 	"github/revert_utility_era2/utils"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// // db.getCollection('Test_AttemptDtls').aggregate([
// //     {
// //         "$match": {"learnerId": "2024302"}
// //     },
// //     {
// //         "$group": {
// //             "_id": "$aggregationId",
// //             "details": {"$push": "$$ROOT"}
// //         }
// //     }
// // ]);

// func ProcessLearner3_13Files(ctx context.Context, learnerId string) error {
// 	pipeline := bson.A{
// 		bson.M{
// 			"$match": bson.M{"learnerId": learnerId},
// 		},
// 		bson.M{
// 			"$group": bson.M{
// 				"_id":     "$aggregationId",
// 				"details": bson.M{"$first": "$$ROOT"},
// 			},
// 		},
// 	}

// 	// Execute the aggregation pipeline
// 	cursor, err := db.ERALiveDB.Collection(LEARNING_TEST_ATTEMP_DETAILS_COLLECTION_NAME).Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return err
// 	}
// 	defer cursor.Close(ctx)

// 	// Iterate over the sessions in the aggregation result
// 	for cursor.Next(ctx) {
// 		var result struct {
// 			ID      string                    `bson:"_id"`
// 			Details models.LearnerExamDetails `bson:"details"`
// 		}

// 		if err := cursor.Decode(&result); err != nil {
// 			return err
// 		}

// 		filePath := "D:/ERA_N/" + result.ID + "_visit.json"
// 		// Write the session details to the JSON file
// 		err = utils.WriteData(filePath, result.Details)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // func GetExamAttemptDetails(ctx context.Context, learnerId string) error {
// // 	pipeline := bson.A{
// // 		bson.M{
// // 			"$match": bson.M{"learnerId": learnerId},
// // 		},
// // 		bson.M{
// // 			"$group": bson.M{
// // 				"_id":     "$aggregationId",
// // 				"details": bson.M{"$push": "$$ROOT"},
// // 			},
// // 		},
// // 	}

// // 	// Execute the aggregation pipeline
// // 	cursor, err := db.ERALiveDB.Collection(LEARNING_TEST_ATTEMP_DETAILS_COLLECTION_NAME).Aggregate(ctx, pipeline)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	defer cursor.Close(ctx)

// // 	// Iterate over the sessions in the aggregation result
// // 	for cursor.Next(ctx) {
// // 		var result struct {
// // 			ID      string                    `bson:"_id"`
// // 			Details models.ExamAttemptDetails `bson:"details"`
// // 		}
// // 		if err := cursor.Decode(&result); err != nil {
// // 			return err
// // 		}

// // 		fmt.Println("Result", result)
// // 	}

// // 	if err := cursor.Err(); err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }

// // func getExamAttemptDetails(ctx context.Context, learnerID string) ([]models.ExamAttemptDetails, error) {
// // 	filter := bson.M{
// // 		"learnerId": learnerID,
// // 	}

// // 	testAttemptcursor, err := db.ERALiveDB.Collection(LEARNING_TEST_ATTEMP_DETAILS_COLLECTION_NAME).Find(ctx, filter)

// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	defer testAttemptcursor.Close(ctx)

// // 	mongoTestAttempts := []models.MongoTestAttemptDetails{}

// // 	if err := testAttemptcursor.All(ctx, &mongoTestAttempts); err != nil {
// // 		return nil, err
// // 	}

// // 	return mongoTestAttempts, nil
// // }

// // func prepareExamAttemptDetailsObject(mtestAttempt models.MongoTestAttemptDetails) models.ExamAttemptDetails {
// // 	examAttemptDetails := models.ExamAttemptDetails{
// // 		ExamID:                          mtestAttempt.ExamID,
// // 		AggregationID:                   mtestAttempt.AggregationID,
// // 		AttemptNumber:                   mtestAttempt.AttemptNumber,
// // 		QuestionAttemptDetails:          prepareQuestionAttemptDetails(mtestAttempt.QuestionAttemptDetails),
// // 		IsExamCompleted:                 mtestAttempt.IsExamCompleted,
// // 		TotalCorrectChars:               int(mtestAttempt.TotalCorrectChars),
// // 		Accuracy:                        mtestAttempt.Accuracy,
// // 		NetWordsPerMinute:               int(mtestAttempt.NetWordsPerMinute),
// // 		IsSectionTimed:                  mtestAttempt.IsSectionTimed,
// // 		DurationInSec:                   int(mtestAttempt.DurationInSec),
// // 		ElapsedTime:                     int(mtestAttempt.ElapsedTime),
// // 		TotalMarksObtained:              mtestAttempt.TotalMarksObtained,
// // 		ExamStartedOn:                   utils.TimestampToTime(mtestAttempt.ExamStartedOn),
// // 		ExamEndedOn:                     utils.TimestampToTime(mtestAttempt.ExamEndedOn),
// // 		TotalMarks:                      mtestAttempt.TotalMarks,
// // 		TotalItems:                      mtestAttempt.TotalItems,
// // 		LanguageID:                      mtestAttempt.LanguageID,
// // 		Percentage:                      mtestAttempt.Percentage,
// // 		AttemptedItemCount:              mtestAttempt.AttemptedItemCount,
// // 		CorrectItemCount:                mtestAttempt.CorrectItemCount,
// // 		WrongItemCount:                  mtestAttempt.WrongItemCount,
// // 		ExamCurrentStatus:               mtestAttempt.ExamCurrentStatus,
// // 		IsQuestionPaperGenerated:        mtestAttempt.IsQuestionPaperGenerated,
// // 		EvaluationID:                    mtestAttempt.EvaluationId,
// // 		EvaluationServerStatus:          mtestAttempt.EvaluationServerStatus,
// // 		LocalEvidenceUploadStatus:       mtestAttempt.LocalEvidenceUploadStatus,
// // 		EvidenceUploadedPath:            mtestAttempt.EvidenceUploadedPath,
// // 		MinimumPassingMarks:             float64(mtestAttempt.MinimumPassingMarks),
// // 		EvidenceUploadAcknowledgementID: mtestAttempt.EvidenceUploadAcknowledgementID,
// // 		ExamEndCause:                    mtestAttempt.ExamEndCause,
// // 	}
// // 	return examAttemptDetails
// // }

// // func GetQuestionAttemptDetailsFromMongo(ctx context.Context, learnerID string) ([]models.QuestionAttemptDetails, error) {
// // 	pipeline := bson.A{
// // 		bson.M{
// // 			"$match": bson.M{
// // 				"learnerId":     learnerID,
// // 				"programId":     programId,
// // 				"eCourseId":     eCourseID,
// // 				"testAttemptId": testAttempt,
// // 			},
// // 		},
// // 	}
// // 	questionAttemptCursor, err := db.ERALiveDB.Collection(LEARNER_QUESTION_ATTEMP_DETAILS_COLLECTION_NAME).Aggregate(ctx, pipeline)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	defer questionAttemptCursor.Close(ctx)

// // 	var questionAttempts []models.QuestionAttemptDetails
// // 	if err := questionAttemptCursor.All(ctx, &questionAttempts); err != nil {
// // 		return nil, err
// // 	}

// // 	return questionAttempts, nil
// // }

// // func prepareQuestionAttemptDetails(questions []models.MongoTestQuestionAttemptDetails) []models.QuestionAttemptDetails {
// // 	var questionAttemptDetails []models.QuestionAttemptDetails

// // 	for _, q := range questions {
// // 		questionAttemptDetails = append(questionAttemptDetails, models.QuestionAttemptDetails{
// // 			QuestionID:                q.QuestionID,
// // 			Ordinality:                q.Ordinality,
// // 			SectionID:                 q.SectionId,
// // 			SectionName:               q.SectionName,
// // 			QuestionType:              q.QuestionType,
// // 			LanguageID:                q.LanguageID,
// // 			AttemptDate:               utils.TimestampToTime(q.AttemptDate),
// // 			QuestionSequenceNo:        q.QuestionSequenceNo,
// // 			WrongCharsCount:           q.WrongCharsCount,
// // 			TypedText:                 q.TypedText,
// // 			CorrectWordsCount:         q.CorrectWordsCount,
// // 			WordsPerMinute:            q.WordsPerMinute,
// // 			TimeElapsed:               q.TimeElapsed,
// // 			TypingAccuracy:            q.TypingAccuracy,
// // 			MarksObtained:             q.MarksObtained,
// // 			SubmittedAnswer:           q.SubmittedAnswer,
// // 			FeedbackReceived:          q.FeedbackReceived,
// // 			QuestionStatus:            q.QuestionStatus,
// // 			QuestionViewedCount:       q.QuestionViewedCount,
// // 			IsQuestionPublished:       q.IsQuestionPublished == "true",
// // 			Tag:                       q.Tag,
// // 			QuestionDifficultyLevel:   q.QuestionDifficultyLevel,
// // 			TotalMarks:                q.TotalMarks,
// // 			DurationInSec:             q.DurationInSec,
// // 			GrossWPM:                  q.GrossWPM,
// // 			TotalDeleteCount:          q.TotalDeleteCount,
// // 			WrongWordCount:            int(q.WrongWordCount),
// // 			TotalKeyStrokes:           int(q.TotalKeyStrokes),
// // 			NetWPM:                    q.NetWPM,
// // 			FeedbackReason:            q.FeedbackReason,
// // 			FeedbackComment:           q.FeedbackComment,
// // 			SimulatorCurrentStepID:    q.SimulatorCurrentStepId,
// // 			SimulatorCurrentControlID: q.SimulatorCurrentControlId,
// // 			SubmitCount:               q.SubmitCount,
// // 		})
// // 	}
// // 	return questionAttemptDetails
// // }

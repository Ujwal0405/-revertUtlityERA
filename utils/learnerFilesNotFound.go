package utils

import (
	"bufio"
	"fmt"
	"github/revert_utility_era2/models"
	"os"
)

func WriteFailedLearnerID3_7(lId string) {
	filePath := models.Get3_7FilePath(lId)

	// Check if the learner ID already exists in the file
	if learnerIDExists(filePath, lId) {
		fmt.Println("Learner ID", lId, "already exists in the file:", filePath)
		return
	}

	// Open the file for writing, create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating the file:", err)
		return
	}
	defer file.Close()

	// Write learner ID
	if _, err := file.WriteString(lId + "\n"); err != nil {
		fmt.Println("Error writing learner ID to file:", err)
	}
}

// Function to check if the learner ID already exists in the file
func learnerIDExists(filePath, learnerID string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false // Assume file does not exist, so learner ID does not exist
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == learnerID {
			return true // Learner ID exists in the file
		}
	}
	return false // Learner ID does not exist in the file
}

package utils

import (
	"encoding/json"

	"github/revert_utility_era2/helper/dalhelper"
)

func WriteData(filePath string, data interface{}) error {
	// dir := filepath.Dir(filePath)
	// filehelper.CreateDirectoryRecursive(dir)
	// filePath := fmt.Sprintf("./3_7_%s.json", learnerID)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := dalhelper.SaveDataToFDB(filePath, jsonData, true); err != nil {
		return err
	}

	return nil
}

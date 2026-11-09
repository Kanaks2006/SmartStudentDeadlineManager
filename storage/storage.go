package storage

import (
	"encoding/json"
	"os"

	"smartstudentdeadline/models"
)

// SaveTasks stores all tasks in a JSON file.
func SaveTasks(tasks []models.Task) error {

	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadTasks reads previously saved tasks from the JSON file.
func LoadTasks() ([]models.Task, error) {

	var tasks []models.Task

	data, err := os.ReadFile("tasks.json")

	if os.IsNotExist(err) {
		return tasks, nil
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

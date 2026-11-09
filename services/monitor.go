package services

import (
	"fmt"
	"strings"
	"time"

	"smartstudentdeadline/models"
)

// DeadlineResult stores the result of one deadline check.
type DeadlineResult struct {
	TaskID   int
	Title    string
	Status   string
	TimeLeft string
}

// checkTaskDeadline checks one task's deadline.
func checkTaskDeadline(task models.Task, results chan DeadlineResult) {

	dueDate, err := time.ParseInLocation("02-01-2006", task.DueDate, time.Local)

	if err != nil {
		results <- DeadlineResult{
			TaskID: task.ID,
			Title:  task.Title,
			Status: "Invalid Date",
		}
		return
	}

	now := time.Now().In(time.Local)

	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		time.Local,
	)

	daysLeft := int(dueDate.Sub(today).Hours() / 24)

	var status string
	var timeLeft string

	if daysLeft < 0 {
		status = "OVERDUE"
		timeLeft = fmt.Sprintf("%d days overdue", -daysLeft)
	} else if daysLeft == 0 {
		status = "DUE TODAY"
		timeLeft = "Deadline is today"
	} else {
		status = "UPCOMING"
		timeLeft = fmt.Sprintf("%d days remaining", daysLeft)
	}

	// Send the result through the channel.
	results <- DeadlineResult{
		TaskID:   task.ID,
		Title:    task.Title,
		Status:   status,
		TimeLeft: timeLeft,
	}
}

// MonitorDeadlines checks all pending tasks concurrently.
func MonitorDeadlines(tasks []models.Task) {

	fmt.Println("\n========== CONCURRENT DEADLINE MONITOR ==========")

	pendingCount := 0

	for _, task := range tasks {
		if strings.EqualFold(task.Status, "Pending") {
			pendingCount++
		}
	}

	if pendingCount == 0 {
		fmt.Println("No pending tasks to monitor!")
		return
	}

	// Create a channel to collect deadline results.
	results := make(chan DeadlineResult, pendingCount)

	// Start one goroutine for every pending task.
	for _, task := range tasks {
		if strings.EqualFold(task.Status, "Pending") {
			go checkTaskDeadline(task, results)
		}
	}

	fmt.Println("Checking deadlines concurrently...")

	// Receive results from all goroutines.
	for i := 0; i < pendingCount; i++ {

		result := <-results

		fmt.Printf("Task ID : %d\n", result.TaskID)
		fmt.Printf("Title   : %s\n", result.Title)
		fmt.Printf("Status  : %s\n", result.Status)
		fmt.Printf("Time    : %s\n", result.TimeLeft)
		fmt.Println("------------------------------------")
	}

	fmt.Println("All deadline checks completed!")
	fmt.Println("=================================================")
}

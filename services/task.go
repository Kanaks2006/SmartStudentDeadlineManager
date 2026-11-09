package services

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"

	"smartstudentdeadline/models"
)

func readInput(reader *bufio.Reader, message string) string {
	fmt.Print(message)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Checks whether a date is valid.
func isValidDate(date string) bool {
	_, err := time.Parse("02-01-2006", date)
	return err == nil
}

// Accepts only Low, Medium or High priority.
func getValidPriority(reader *bufio.Reader) string {
	for {
		priority := readInput(reader, "Enter Priority (Low/Medium/High): ")

		switch strings.ToLower(priority) {
		case "low":
			return "Low"
		case "medium":
			return "Medium"
		case "high":
			return "High"
		default:
			fmt.Println("Invalid priority! Please enter Low, Medium or High.")
		}
	}
}

// Accepts only Pending or Completed status.
func getValidStatus(reader *bufio.Reader) string {
	for {
		status := readInput(reader, "New Status (Pending/Completed): ")

		switch strings.ToLower(status) {
		case "pending":
			return "Pending"
		case "completed":
			return "Completed"
		default:
			fmt.Println("Invalid status! Enter Pending or Completed.")
		}
	}
}

// AddTask collects and saves task details.
func AddTask(reader *bufio.Reader, taskID int) models.Task {
	var task models.Task

	fmt.Println("\n========== ADD NEW TASK ==========")

	task.ID = taskID

	for {
		task.Title = readInput(reader, "Enter Task Title: ")
		if task.Title != "" {
			break
		}
		fmt.Println("Task title cannot be empty!")
	}

	task.Description = readInput(reader, "Enter Description: ")
	task.Subject = readInput(reader, "Enter Subject: ")
	task.TaskType = readInput(reader, "Enter Task Type: ")

	for {
		task.DueDate = readInput(reader, "Enter Due Date (DD-MM-YYYY): ")

		if isValidDate(task.DueDate) {
			break
		}

		fmt.Println("Invalid date! Please enter a real date in DD-MM-YYYY format.")
	}

	task.Priority = getValidPriority(reader)
	task.Status = "Pending"
	task.CalculatedPriority = CalculatePriority(task.DueDate)

	fmt.Println("\nTask added successfully!")
	fmt.Printf("Task ID             : %d\n", task.ID)
	fmt.Printf("Title               : %s\n", task.Title)
	fmt.Printf("Due Date            : %s\n", task.DueDate)
	fmt.Printf("Original Priority   : %s\n", task.Priority)
	fmt.Printf("Calculated Priority : %s\n", task.CalculatedPriority)
	fmt.Printf("Status              : %s\n", task.Status)

	return task
}

// ViewTasks displays all saved tasks.
func ViewTasks(tasks []models.Task) {
	fmt.Println("\n========== ALL YOUR TASKS ==========")

	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}

	for _, task := range tasks {
		calculated := task.CalculatedPriority

		// Calculate priority for older saved tasks if missing.
		if calculated == "" {
			calculated = CalculatePriority(task.DueDate)
		}

		fmt.Printf("\nTask ID             : %d\n", task.ID)
		fmt.Printf("Title               : %s\n", task.Title)
		fmt.Printf("Description         : %s\n", task.Description)
		fmt.Printf("Subject             : %s\n", task.Subject)
		fmt.Printf("Task Type           : %s\n", task.TaskType)
		fmt.Printf("Due Date            : %s\n", task.DueDate)
		fmt.Printf("Original Priority   : %s\n", task.Priority)
		fmt.Printf("Calculated Priority : %s\n", calculated)
		fmt.Printf("Status              : %s\n", task.Status)
		fmt.Println("------------------------------------")
	}

	fmt.Printf("Total Tasks: %d\n", len(tasks))
}

// UpdateTask modifies an existing task.
func UpdateTask(reader *bufio.Reader, tasks []models.Task) {
	fmt.Println("\n========== UPDATE TASK ==========")

	if len(tasks) == 0 {
		fmt.Println("No tasks available to update!")
		return
	}

	idText := readInput(reader, "Enter the Task ID you want to update: ")
	taskID, err := strconv.Atoi(idText)

	if err != nil {
		fmt.Println("Invalid Task ID!")
		return
	}

	for i := range tasks {
		if tasks[i].ID == taskID {
			fmt.Println("\nTask found! Press Enter to keep the existing value.")

			value := readInput(reader, "New Title: ")
			if value != "" {
				tasks[i].Title = value
			}

			value = readInput(reader, "New Description: ")
			if value != "" {
				tasks[i].Description = value
			}

			value = readInput(reader, "New Subject: ")
			if value != "" {
				tasks[i].Subject = value
			}

			value = readInput(reader, "New Task Type: ")
			if value != "" {
				tasks[i].TaskType = value
			}

			for {
				value = readInput(reader, "New Due Date: ")

				if value == "" {
					break
				}

				if isValidDate(value) {
					tasks[i].DueDate = value
					break
				}

				fmt.Println("Invalid date! Enter a real date in DD-MM-YYYY format.")
			}

			for {
				value = readInput(reader, "New Priority (Low/Medium/High): ")

				if value == "" {
					break
				}

				switch strings.ToLower(value) {
				case "low":
					tasks[i].Priority = "Low"
				case "medium":
					tasks[i].Priority = "Medium"
				case "high":
					tasks[i].Priority = "High"
				default:
					fmt.Println("Invalid priority! Enter Low, Medium or High.")
					continue
				}
				break
			}

			for {
				value = readInput(reader, "New Status (Pending/Completed): ")

				if value == "" {
					break
				}

				switch strings.ToLower(value) {
				case "pending":
					tasks[i].Status = "Pending"
				case "completed":
					tasks[i].Status = "Completed"
				default:
					fmt.Println("Invalid status! Enter Pending or Completed.")
					continue
				}
				break
			}

			tasks[i].CalculatedPriority = CalculatePriority(tasks[i].DueDate)

			fmt.Println("\nTask updated successfully!")
			fmt.Printf("Title               : %s\n", tasks[i].Title)
			fmt.Printf("Due Date            : %s\n", tasks[i].DueDate)
			fmt.Printf("Original Priority   : %s\n", tasks[i].Priority)
			fmt.Printf("Calculated Priority : %s\n", tasks[i].CalculatedPriority)
			fmt.Printf("Status              : %s\n", tasks[i].Status)
			return
		}
	}

	fmt.Println("Task ID not found!")
}

// DeleteTask removes a task using its ID.
func DeleteTask(reader *bufio.Reader, tasks []models.Task) []models.Task {
	fmt.Println("\n========== DELETE TASK ==========")

	if len(tasks) == 0 {
		fmt.Println("No tasks available to delete!")
		return tasks
	}

	idText := readInput(reader, "Enter the Task ID you want to delete: ")
	taskID, err := strconv.Atoi(idText)

	if err != nil {
		fmt.Println("Invalid Task ID!")
		return tasks
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("\nTask deleted successfully!")
			return tasks
		}
	}

	fmt.Println("Task ID not found!")
	return tasks
}

// ViewDeadlines checks the status of all task deadlines.
func ViewDeadlines(tasks []models.Task) {
	fmt.Println("\n========== UPCOMING DEADLINES ==========")

	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	for _, task := range tasks {
		dueDate, err := time.ParseInLocation("02-01-2006", task.DueDate, time.Local)

		fmt.Printf("\nTask ID : %d\n", task.ID)
		fmt.Printf("Title   : %s\n", task.Title)
		fmt.Printf("Due Date: %s\n", task.DueDate)

		if err != nil {
			fmt.Println("Deadline Status: Invalid date format")
		} else if dueDate.Before(today) {
			fmt.Println("Deadline Status: OVERDUE")
		} else if dueDate.Equal(today) {
			fmt.Println("Deadline Status: DUE TODAY")
		} else {
			fmt.Println("Deadline Status: UPCOMING")
		}

		fmt.Printf("Task Status: %s\n", task.Status)
		fmt.Println("------------------------------------")
	}
}

// CalculatePriority automatically calculates task urgency.
func CalculatePriority(dueDateText string) string {
	dueDate, err := time.ParseInLocation("02-01-2006", dueDateText, time.Local)

	if err != nil {
		return "Unknown"
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	daysLeft := int(dueDate.Sub(today).Hours() / 24)

	if daysLeft <= 0 {
		return "Critical"
	} else if daysLeft <= 2 {
		return "High"
	} else if daysLeft <= 5 {
		return "Medium"
	}

	return "Low"
}

// ViewWorkload displays the student's pending workload.
func ViewWorkload(tasks []models.Task) {
	fmt.Println("\n========== PENDING WORKLOAD ==========")

	totalTasks := len(tasks)
	pending := 0
	completed := 0
	high := 0
	medium := 0
	low := 0
	critical := 0

	for _, task := range tasks {
		if strings.EqualFold(task.Status, "Pending") {
			pending++

			priority := task.CalculatedPriority
			if priority == "" {
				priority = CalculatePriority(task.DueDate)
			}

			switch strings.ToLower(priority) {
			case "critical":
				critical++
			case "high":
				high++
			case "medium":
				medium++
			case "low":
				low++
			}
		} else if strings.EqualFold(task.Status, "Completed") {
			completed++
		}
	}

	fmt.Printf("\nTotal Tasks             : %d\n", totalTasks)
	fmt.Printf("Pending Tasks           : %d\n", pending)
	fmt.Printf("Completed Tasks         : %d\n", completed)

	fmt.Println("\n----- AUTOMATIC PRIORITY -----")
	fmt.Printf("Critical Priority       : %d\n", critical)
	fmt.Printf("High Priority           : %d\n", high)
	fmt.Printf("Medium Priority         : %d\n", medium)
	fmt.Printf("Low Priority            : %d\n", low)

	fmt.Println("\n====================================")
}

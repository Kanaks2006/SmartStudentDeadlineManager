package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"smartstudentdeadline/models"
	"smartstudentdeadline/services"
	"smartstudentdeadline/storage"
)

func saveTasks(tasks []models.Task) {
	err := storage.SaveTasks(tasks)

	if err != nil {
		fmt.Println("Error saving tasks:", err)
	} else {
		fmt.Println("Tasks saved successfully!")
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	tasks, err := storage.LoadTasks()

	if err != nil {
		fmt.Println("Error loading tasks:", err)
		tasks = []models.Task{}
	}

	taskID := 1

	for _, task := range tasks {
		if task.ID >= taskID {
			taskID = task.ID + 1
		}
	}

	fmt.Println("Tasks loaded successfully!")

	for {

		fmt.Println("\n========================================")
		fmt.Println("     SMART STUDENT DEADLINE MANAGER")
		fmt.Println("========================================")

		fmt.Println("\n----------- MAIN MENU -----------")
		fmt.Println("1. Add New Task")
		fmt.Println("2. View All Tasks")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. View Upcoming Deadlines")
		fmt.Println("6. View Pending Workload")
		fmt.Println("7. Concurrent Deadline Monitoring")
		fmt.Println("8. Exit")

		fmt.Print("\nEnter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":
			task := services.AddTask(reader, taskID)
			tasks = append(tasks, task)
			taskID++
			saveTasks(tasks)

		case "2":
			services.ViewTasks(tasks)

		case "3":
			services.UpdateTask(reader, tasks)
			saveTasks(tasks)

		case "4":
			tasks = services.DeleteTask(reader, tasks)
			saveTasks(tasks)

		case "5":
			services.ViewDeadlines(tasks)

		case "6":
			services.ViewWorkload(tasks)

		case "7":
			services.MonitorDeadlines(tasks)

		case "8":
			fmt.Println("\nThank you for using Smart Student Deadline Manager!")
			return

		default:
			fmt.Println("Invalid choice! Enter a number between 1 and 8.")
		}
	}
}


# Smart Student Task & Deadline Manager

A command-line academic task management application developed using **Go (Golang)** to help students organize academic activities, manage assignments, track deadlines, and monitor pending workloads.

## Overview

Managing multiple academic responsibilities can be challenging. This application provides a simple and structured way to record tasks, monitor deadlines, and prioritize unfinished work.

The project demonstrates practical Go programming concepts, including modular programming, JSON-based data storage, error handling, and concurrent processing.

## Features

- **Add Tasks:** Create academic tasks with titles, descriptions, subjects, task types, due dates, and priorities.
- **View Tasks:** Display all saved academic tasks.
- **Update Tasks:** Modify existing task details and completion status.
- **Delete Tasks:** Remove tasks when they are no longer required.
- **Deadline Tracking:** Identify upcoming and overdue tasks.
- **Smart Priority Calculation:** Automatically calculate task priority based on the time remaining before its deadline.
- **Pending Workload:** View an overview of unfinished tasks.
- **Concurrent Deadline Monitoring:** Monitor multiple pending deadlines concurrently using goroutines and channels.
- **JSON Data Storage:** Save and retrieve task information using JSON file handling.

## Technologies & Concepts Used

- Go (Golang)
- Structures and Functions
- Packages and Modular Programming
- Slices
- JSON Encoding and Decoding
- File Handling
- Date and Time Operations
- Input Validation and Error Handling
- Goroutines and Channels
- Concurrent Processing

## Project Structure

```text
SmartStudentDeadlineManager/
├── main.go
├── go.mod
├── models/
│   └── task.go
├── services/
│   ├── task.go
│   └── monitor.go
└── storage/
    └── storage.go
```

## Getting Started

### Prerequisites

- Go installed on your computer.
- Visual Studio Code or another code editor.

### Installation

1. Clone this repository:

   ```bash
   git clone YOUR_REPOSITORY_URL
   ```

2. Navigate to the project directory:

   ```bash
   cd SmartStudentDeadlineManager
   ```

3. Run the application:

   ```bash
   go run .
   ```

4. Follow the options displayed in the terminal menu.

The application creates a local `tasks.json` file to store task information.

## Learning Objectives

This project was developed to apply Go programming concepts through a practical application, focusing on modular code organization, data persistence, task management, and concurrency.

## Author

**Kanak Sharma**  
B.Sc. Computer Science

**Project Type:** Academic Capstone Project  
**Programming Language:** Go (Golang)

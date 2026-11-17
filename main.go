// Command student-result-processing is a modular CLI application that lets
// a user enter students and their subject marks, computes each student's
// total, percentage, grade and pass/fail status, and shows class statistics.
//
// The program is split into packages under internal/ so that data types,
// grading rules, business logic, storage and user interface each live in
// their own module — see README.md for the full layout.
package main

import (
	"fmt"
	"os"

	"github.com/kesavaramk/student-result-processing/internal/models"
	"github.com/kesavaramk/student-result-processing/internal/processor"
	"github.com/kesavaramk/student-result-processing/internal/storage"
	"github.com/kesavaramk/student-result-processing/internal/ui"
)

// dataFile is the path used to persist students between runs.
const dataFile = "data/students.json"

func main() {
	p := processor.NewProcessor()

	// Attempt to load any previously saved data so the class survives a
	// restart. LoadFromFile returns an empty slice on first run.
	if students, err := storage.LoadFromFile(dataFile); err == nil {
		p.Students = students
	} else {
		fmt.Println("Warning: could not load saved data:", err)
	}

	ui.Header("Student Result Processing System")
	fmt.Println("Records loaded:", p.Count())

	for {
		showMenu()
		switch ui.Prompt("Enter your choice: ") {
		case "1":
			addStudent(p)
		case "2":
			ui.Header("All Student Results")
			ui.DisplayAll(p.Students)
			ui.Pause()
		case "3":
			searchStudent(p)
		case "4":
			showStatistics(p)
		case "5":
			saveData(p)
		case "6":
			fmt.Println("Thank you for using the Student Result Processing System.")
			return
		default:
			fmt.Println("Invalid choice. Please select an option from 1 to 6.")
		}
	}
}

// showMenu prints the available operations.
func showMenu() {
	ui.Header("Main Menu")
	fmt.Println("1. Add Student")
	fmt.Println("2. Display All Results")
	fmt.Println("3. Search Student by Roll Number")
	fmt.Println("4. Class Statistics")
	fmt.Println("5. Save Data to File")
	fmt.Println("6. Exit")
}

// addStudent collects a student's details and subject marks from the user,
// computes the result, and stores the record. It demonstrates loops for
// collecting a variable number of subjects and validation of inputs.
func addStudent(p *processor.Processor) {
	ui.Header("Add New Student")

	rollNo := ui.PromptInt("Enter Roll Number: ")
	if p.RollNoExists(rollNo) {
		fmt.Printf("A student with roll number %d already exists.\n", rollNo)
		ui.Pause()
		return
	}

	name := ui.Prompt("Enter Student Name: ")
	if name == "" {
		fmt.Println("Name cannot be empty.")
		ui.Pause()
		return
	}

	count := ui.PromptInt("Enter number of subjects: ")
	if count <= 0 {
		fmt.Println("There must be at least one subject.")
		ui.Pause()
		return
	}

	// Collect each subject and its marks in a slice, validating that marks
	// stay within the allowed range (0 to MaxMarksPerSubject).
	subjects := make([]models.Subject, 0, count)
	for i := 0; i < count; i++ {
		fmt.Printf("\n--- Subject %d ---\n", i+1)
		subjectName := ui.Prompt("  Subject name: ")
		if subjectName == "" {
			fmt.Println("  Subject name cannot be empty; skipping will retry.")
			i-- // retry this index
			continue
		}
		marks := ui.PromptFloatBounded(
			"  Marks (0-100): ", 0, models.MaxMarksPerSubject)
		subjects = append(subjects, models.Subject{Name: subjectName, Marks: marks})
	}

	student := models.Student{
		RollNo:   rollNo,
		Name:     name,
		Subjects: subjects,
	}

	p.CalculateResults(&student)
	p.AddStudent(student)

	fmt.Println("\nStudent added successfully. Result card:")
	ui.DisplayStudent(student)
	ui.Pause()
}

// searchStudent looks up a student by roll number and prints their card.
func searchStudent(p *processor.Processor) {
	ui.Header("Search Student")
	if p.Count() == 0 {
		fmt.Println("No records to search. Add a student first.")
		ui.Pause()
		return
	}
	rollNo := ui.PromptInt("Enter Roll Number to search: ")
	student, err := p.GetStudentByRollNo(rollNo)
	if err != nil {
		fmt.Println(err)
		ui.Pause()
		return
	}
	ui.DisplayStudent(*student)
	ui.Pause()
}

// showStatistics prints summary information about the whole class.
func showStatistics(p *processor.Processor) {
	ui.Header("Class Statistics")
	if p.Count() == 0 {
		fmt.Println("No records available.")
		ui.Pause()
		return
	}
	fmt.Printf("Total Students : %d\n", p.Count())
	fmt.Printf("Passed         : %d\n", p.PassCount())
	fmt.Printf("Failed         : %d\n", p.FailCount())
	fmt.Printf("Class Average  : %.2f%%\n", p.ClassAverage())

	if top, err := p.TopPerformer(); err == nil {
		fmt.Println("\nTop Performer:")
		ui.DisplayStudent(*top)
	}
	ui.Pause()
}

// saveData writes the current class to the data file so it can be reloaded.
func saveData(p *processor.Processor) {
	// Ensure the data directory exists before writing.
	if err := os.MkdirAll("data", 0755); err != nil {
		fmt.Println("Could not create data directory:", err)
		ui.Pause()
		return
	}
	if err := storage.SaveToFile(p.Students, dataFile); err != nil {
		fmt.Println("Failed to save data:", err)
	} else {
		fmt.Printf("Saved %d record(s) to %s\n", p.Count(), dataFile)
	}
	ui.Pause()
}

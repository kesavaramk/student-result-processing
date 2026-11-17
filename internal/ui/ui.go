// Package ui handles all interaction with the user through the terminal:
// reading input, validating it, and printing formatted output. Centralising
// I/O here keeps main.go small and lets the processor package stay free
// of any printing or scanning logic.
package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kesavaramk/student-result-processing/internal/models"
)

// reader is a package-level buffered reader wrapping os.Stdin. Using a
// single reader across calls keeps the input buffer consistent.
var reader = bufio.NewReader(os.Stdin)

// Line reads a single line from standard input and trims surrounding
// whitespace. The error is ignored because EOF on stdin is a normal way
// to leave the program.
func Line() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Prompt prints a prompt and returns the trimmed line the user enters.
func Prompt(label string) string {
	fmt.Print(label)
	return Line()
}

// PromptInt reads an integer using strconv.Atoi, retrying until the user
// enters a valid value. This demonstrates control flow (loop + condition)
// and the strconv package for string-to-number conversion.
func PromptInt(label string) int {
	for {
		value, err := strconv.Atoi(Prompt(label))
		if err == nil {
			return value
		}
		fmt.Println("  Please enter a valid whole number.")
	}
}

// PromptFloat reads a floating-point number, retrying on invalid input.
func PromptFloat(label string) float64 {
	for {
		value, err := strconv.ParseFloat(Prompt(label), 64)
		if err == nil {
			return value
		}
		fmt.Println("  Please enter a valid number.")
	}
}

// PromptFloatBounded reads a float constrained to [min, max], retrying
// until the value is within range. Used to keep marks between 0 and 100.
func PromptFloatBounded(label string, min, max float64) float64 {
	for {
		value := PromptFloat(label)
		if value >= min && value <= max {
			return value
		}
		fmt.Printf("  Value must be between %.0f and %.0f.\n", min, max)
	}
}

// Pause waits for the user to press Enter before continuing, which keeps
// output on screen long enough to read in menu-driven programs.
func Pause() {
	fmt.Print("\nPress Enter to continue...")
	Line()
}

// Header prints a titled banner so each screen is clearly delimited.
func Header(title string) {
	bar := strings.Repeat("=", 56)
	fmt.Println()
	fmt.Println(bar)
	fmt.Println(center(title, 56))
	fmt.Println(bar)
}

// DisplayStudent prints the full result card for a single student.
func DisplayStudent(s models.Student) {
	fmt.Printf("Roll No    : %d\n", s.RollNo)
	fmt.Printf("Name       : %s\n", s.Name)
	fmt.Println("Subjects   :")
	for _, sub := range s.Subjects {
		fmt.Printf("   %-20s : %6.2f\n", sub.Name, sub.Marks)
	}
	fmt.Printf("Total      : %6.2f\n", s.Total)
	fmt.Printf("Percentage : %6.2f%%\n", s.Percentage)
	fmt.Printf("Grade      : %s\n", s.Grade)
	fmt.Printf("Status     : %s\n", s.Status)
}

// DisplayAll prints every student as a formatted table. The table layout
// demonstrates formatted printing with fmt.Printf and width specifiers.
func DisplayAll(students []models.Student) {
	if len(students) == 0 {
		fmt.Println("No student records to display.")
		return
	}
	fmt.Printf("%-8s %-20s %8s %10s %6s %6s\n",
		"RollNo", "Name", "Total", "Percentage", "Grade", "Status")
	fmt.Println(strings.Repeat("-", 64))
	for _, s := range students {
		fmt.Printf("%-8d %-20s %8.2f %9.2f%% %6s %6s\n",
			s.RollNo, s.Name, s.Total, s.Percentage, s.Grade, s.Status)
	}
	fmt.Println(strings.Repeat("-", 64))
}

// center pads a string with spaces so it is centred within the given width.
// A small helper showing string manipulation with the strings package.
func center(s string, width int) string {
	if len(s) >= width {
		return s
	}
	pad := (width - len(s)) / 2
	return strings.Repeat(" ", pad) + s
}

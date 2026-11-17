// Package models defines the core data structures used across the
// application. Keeping the data types in their own package means every
// other package (grades, processor, storage, ui) can depend on these
// types without creating import cycles.
package models

// Subject represents a single subject studied by a student.
// Marks is stored as float64 so that decimal marks (e.g. 87.5) are supported.
type Subject struct {
	Name  string  `json:"name"`
	Marks float64 `json:"marks"`
}

// Student represents a student along with their marks and the computed
// result fields. The computed fields (Total, Percentage, Grade, Status)
// are filled in by the processor after the raw marks are entered.
type Student struct {
	RollNo     int       `json:"rollNo"`
	Name       string    `json:"name"`
	Subjects   []Subject `json:"subjects"`   // slice of subjects
	Total      float64   `json:"total"`      // sum of all subject marks
	Percentage float64   `json:"percentage"` // average across subjects
	Grade      string    `json:"grade"`      // letter grade (A+..F)
	Status     string    `json:"status"`     // "Pass" or "Fail"
}

// MaxMarksPerSubject is the maximum marks a single subject can carry.
// Defined as a constant so it can be reused and changed in one place.
const MaxMarksPerSubject = 100.0

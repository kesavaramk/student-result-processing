// Package processor holds the business logic of the application: it stores
// the list of students, computes their results, and produces statistics
// such as the class average and the top performer. This is the "engine"
// of the program; the ui package is the "dashboard" that drives it.
package processor

import (
	"fmt"

	"github.com/kesavaramk/student-result-processing/internal/grades"
	"github.com/kesavaramk/student-result-processing/internal/models"
)

// Processor manages the collection of students and the operations performed
// on them. Using a struct with methods keeps related state and behaviour
// together, which is Go's idiomatic way of grouping a "service".
type Processor struct {
	Students []models.Student
}

// NewProcessor is a constructor function that returns a Processor with an
// initialised (empty but non-nil) slice. Constructors keep object creation
// consistent and centralised.
func NewProcessor() *Processor {
	return &Processor{
		Students: []models.Student{},
	}
}

// AddStudent appends a student to the collection. Slices in Go grow with
// the built-in append function, which handles reallocation for us.
func (p *Processor) AddStudent(s models.Student) {
	p.Students = append(p.Students, s)
}

// CalculateResults computes the Total, Percentage, Grade and Status fields
// of a student in place (the receiver pointer lets us mutate the struct).
// It demonstrates a for-range loop over a slice.
func (p *Processor) CalculateResults(s *models.Student) {
	var total float64
	for _, subject := range s.Subjects {
		total += subject.Marks
	}
	s.Total = total
	if count := len(s.Subjects); count > 0 {
		s.Percentage = total / float64(count)
	}
	s.Grade = grades.CalculateGrade(s.Percentage)
	s.Status = grades.StatusText(s.Percentage)
}

// GetStudentByRollNo returns a pointer to the student whose roll number
// matches, or an error if none is found. Returning a pointer lets the
// caller inspect the stored record directly.
func (p *Processor) GetStudentByRollNo(rollNo int) (*models.Student, error) {
	for i := range p.Students {
		if p.Students[i].RollNo == rollNo {
			return &p.Students[i], nil
		}
	}
	return nil, fmt.Errorf("no student found with roll number %d", rollNo)
}

// RollNoExists reports whether a given roll number is already in use,
// which the UI uses to prevent duplicate entries.
func (p *Processor) RollNoExists(rollNo int) bool {
	for i := range p.Students {
		if p.Students[i].RollNo == rollNo {
			return true
		}
	}
	return false
}

// Count returns the number of students currently stored.
func (p *Processor) Count() int {
	return len(p.Students)
}

// ClassAverage returns the average percentage across all students.
// Returns 0 when there are no students to avoid a divide-by-zero.
func (p *Processor) ClassAverage() float64 {
	if len(p.Students) == 0 {
		return 0
	}
	var sum float64
	for _, s := range p.Students {
		sum += s.Percentage
	}
	return sum / float64(len(p.Students))
}

// PassCount returns how many students passed.
func (p *Processor) PassCount() int {
	count := 0
	for _, s := range p.Students {
		if s.Status == "Pass" {
			count++
		}
	}
	return count
}

// FailCount returns how many students failed.
func (p *Processor) FailCount() int {
	return p.Count() - p.PassCount()
}

// TopPerformer returns a pointer to the student with the highest
// percentage, or an error when the class is empty.
func (p *Processor) TopPerformer() (*models.Student, error) {
	if len(p.Students) == 0 {
		return nil, fmt.Errorf("no students recorded yet")
	}
	top := &p.Students[0]
	for i := range p.Students {
		if p.Students[i].Percentage > top.Percentage {
			top = &p.Students[i]
		}
	}
	return top, nil
}

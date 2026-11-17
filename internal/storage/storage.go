// Package storage handles persisting the list of students to disk and
// loading it back, using JSON as the file format. Keeping I/O separate
// from business logic means the processor never needs to know where or
// how the data is stored.
package storage

import (
	"encoding/json"
	"os"

	"github.com/kesavaramk/student-result-processing/internal/models"
)

// SaveToFile serialises the students slice to JSON and writes it to the
// file at filename. MarshalIndent produces human-readable, indented JSON
// which is convenient for inspection and version control. The file is
// created with permission 0644 (read/write for owner, read for others).
func SaveToFile(students []models.Student, filename string) error {
	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile reads the JSON file and deserialises it into a slice of
// students. If the file does not exist yet (first run), an empty slice
// and a nil error are returned so the application can start fresh.
func LoadFromFile(filename string) ([]models.Student, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// First run: no data file yet, start with an empty class.
			return []models.Student{}, nil
		}
		return nil, err
	}
	var students []models.Student
	if err := json.Unmarshal(data, &students); err != nil {
		return nil, err
	}
	return students, nil
}

// Package grades contains the grading rules used to convert a numeric
// percentage into a letter grade, a pass/fail decision, and grade points.
// Isolating these rules in one package makes them easy to tune without
// touching the rest of the program.
package grades

// PassPercentage is the minimum percentage required to pass.
const PassPercentage = 40.0

// CalculateGrade maps a percentage to a letter grade using a switch with
// case expressions evaluated top to bottom. This demonstrates Go's
// switch statement with conditional (expression-less) cases.
func CalculateGrade(percentage float64) string {
	switch {
	case percentage >= 90:
		return "A+"
	case percentage >= 80:
		return "A"
	case percentage >= 70:
		return "B"
	case percentage >= 60:
		return "C"
	case percentage >= 50:
		return "D"
	case percentage >= PassPercentage:
		return "E"
	default:
		return "F"
	}
}

// IsPass reports whether the given percentage clears the pass threshold.
// Returns a bool, demonstrating Go's bool return type.
func IsPass(percentage float64) bool {
	return percentage >= PassPercentage
}

// GradePoint converts a letter grade into a numeric grade point on a
// 10-point scale. Used for computing the GPA of a student.
func GradePoint(grade string) float64 {
	switch grade {
	case "A+":
		return 10.0
	case "A":
		return 9.0
	case "B":
		return 8.0
	case "C":
		return 7.0
	case "D":
		return 6.0
	case "E":
		return 5.0
	case "F":
		return 0.0
	default:
		return 0.0
	}
}

// StatusText returns a human-readable pass/fail label for a percentage.
func StatusText(percentage float64) string {
	if IsPass(percentage) {
		return "Pass"
	}
	return "Fail"
}

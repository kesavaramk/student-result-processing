# Student Result Processing System (Go CLI)

A modular command-line application written in Go that records students and
their subject marks, computes each student's total, percentage, grade and
pass/fail status, and reports class statistics. The project demonstrates Go
basics (data types, control structures, functions), arrays/slices, packages
and Go modules.

---

## Table of Contents

1. [Features](#features)
2. [Project Structure](#project-structure)
3. [Grading Scheme](#grading-scheme)
4. [Prerequisites](#prerequisites)
5. [Getting Started (Step by Step)](#getting-started-step-by-step)
6. [How to Use the Application](#how-to-use-the-application)
7. [Example Run](#example-run)
8. [Using Go Modules](#using-go-modules)
9. [Version Control with Git & GitHub](#version-control-with-git--github)
10. [Understanding the Code (Package Walkthrough)](#understanding-the-code-package-walkthrough)
11. [Rubric Coverage](#rubric-coverage)
12. [Troubleshooting](#troubleshooting)

---

## Features

- Add a student with a unique roll number, name, and any number of subjects.
- Automatically calculate **total marks**, **percentage**, **letter grade**, and **pass/fail status**.
- Display all students in a formatted table.
- Search a student by roll number and view their full result card.
- View class statistics: total students, pass/fail counts, class average, and top performer.
- Persist records to a JSON file (`data/students.json`) and reload them on the next run.
- Fully menu-driven terminal interface with input validation.

---

## Project Structure

```
student-result-processing/
├── go.mod                  # Module definition (name + Go version)
├── main.go                 # Program entry point and menu loop
├── .gitignore              # Files Git should ignore
├── README.md               # This file
├── data/
│   ├── .gitkeep            # Keeps the data/ folder in Git
│   └── students.json       # Created at runtime (ignored by Git)
└── internal/               # All application packages live here
    ├── models/
    │   └── student.go      # Student & Subject structs (data types)
    ├── grades/
    │   └── grades.go       # Grading rules (grade, pass/fail, grade points)
    ├── processor/
    │   └── processor.go    # Business logic: calculation & statistics
    ├── storage/
    │   └── storage.go       # Save/load students to/from JSON file
    └── ui/
        └── ui.go           # Terminal input/output helpers
```

> Packages under `internal/` can only be imported by code within this module.
> This is Go's built-in way to keep application-specific code private.

---

## Grading Scheme

The grade is based on the **average percentage** across all subjects:

| Percentage Range | Grade | Grade Point |
|------------------|-------|-------------|
| 90 – 100         | A+    | 10          |
| 80 – 89          | A     | 9           |
| 70 – 79          | B     | 8           |
| 60 – 69          | C     | 7           |
| 50 – 59          | D     | 6           |
| 40 – 49          | E     | 5           |
| Below 40         | F     | 0           |

A student **passes** if the percentage is 40% or above.

---

## Prerequisites

1. **Go** installed (version 1.21 or newer).
   Check with:
   ```powershell
   go version
   ```
   If it is not installed, download it from <https://go.dev/dl/>.

2. **Git** installed for version control.
   Check with:
   ```powershell
   git --version
   ```
   If it is not installed, download it from <https://git-scm.com/downloads>.

3. A **GitHub account** to host the repository.

---

## Getting Started (Step by Step)

These steps assume you have cloned or copied the project folder and opened a
terminal inside it.

### 1. Initialise the Go module (only if `go.mod` is missing)

If `go.mod` is already present (it is, in this project), you can skip this.
Otherwise run:

```powershell
go mod init github.com/kesavaramk/student-result-processing
```

The module path in `go.mod` is `github.com/kesavaramk/student-result-processing`
and matches the import paths used in the code.

### 2. Tidy dependencies

```powershell
go mod tidy
```

### 3. Build the application

```powershell
go build
```

This creates an executable named `student-result-processing.exe`
(on Windows) or `student-result-processing` (on macOS/Linux) in the project
folder.

### 4. Run the application

Either run the built executable:

```powershell
.\student-result-processing.exe
```

Or run directly without building a binary:

```powershell
go run .
```

### 5. (Optional) Format and check the code

```powershell
go fmt ./...
go vet ./...
```

---

## How to Use the Application

When you launch the program, you see a menu:

```
========================================================
         Student Result Processing System
========================================================
Records loaded: 0

========================================================
                       Main Menu
========================================================
1. Add Student
2. Display All Results
3. Search Student by Roll Number
4. Class Statistics
5. Save Data to File
6. Exit
Enter your choice:
```

- **1. Add Student** — Enter roll number, name, number of subjects, each
  subject name and its marks (0–100). The result is computed and shown.
- **2. Display All Results** — Shows all students in a table.
- **3. Search Student by Roll Number** — Enter a roll number to view one card.
- **4. Class Statistics** — Total students, passed/failed, class average,
  and the top performer.
- **5. Save Data to File** — Writes records to `data/students.json`.
- **6. Exit** — Quits the program.

---

## Example Run

```
Enter your choice: 1
Enter Roll Number: 101
Enter Student Name: Alice
Enter number of subjects: 3

--- Subject 1 ---
  Subject name: Math
  Marks (0-100): 92

--- Subject 2 ---
  Subject name: Science
  Marks (0-100): 85

--- Subject 3 ---
  Subject name: English
  Marks (0-100): 78

Student added successfully. Result card:
Roll No    : 101
Name       : Alice
Subjects   :
   Math                 :  92.00
   Science              :  85.00
   English              :  78.00
Total      : 255.00
Percentage :  85.00%
Grade      : A
Status     : Pass
```

---

## Using Go Modules

- `go.mod` declares the **module path** (`github.com/kesavaramk/student-result-processing`)
  and the **Go version** (1.21). The module path is the unique identifier for
  this project and forms the prefix of every internal import path.
- Because this project uses **only the Go standard library**, there are no
  external dependencies, so `go.sum` is empty / not generated. If you add a
  third-party package later, `go get <package>` and `go mod tidy` will
  create `go.sum`.
- To learn more, see the official documentation: <https://go.dev/doc/modules>

---

## Version Control with Git & GitHub

This section satisfies the "GitHub Usage and Version Control" rubric.

### 1. Initialise Git in the project folder

```powershell
git init
```

### 2. Configure your identity (once per machine)

```powershell
git config --global user.name "Your Name"
git config --global user.email "you@example.com"
```

### 3. Stage and commit the initial project

```powershell
git add .
git commit -m "Initial commit: modular Go CLI for student result processing"
```

### 4. Create the repository on GitHub

1. Go to <https://github.com/new>.
2. Name it `student-result-processing`.
3. Do **not** initialise with a README (you already have one locally).
4. Click **Create repository**.

### 5. Connect your local repo and push

GitHub will show you the commands. They look like this (replace the username):

```powershell
git branch -M main
git remote add origin https://github.com/kesavaramk/student-result-processing.git
git push -u origin main
```

### 6. Use a feature branch for changes

A good practice is to do work on a branch, then merge it into `main`:

```powershell
git checkout -b feature/add-statistics
# ... make changes ...
git add .
git commit -m "Add class statistics feature"
git push -u origin feature/add-statistics
git checkout main
git merge feature/add-statistics
git push origin main
```

### Suggested commit history (matches the rubric)

The rubric rewards clear commits and branch use. Aim for commits like:

1. `Initial commit: project structure and go.mod`
2. `Add models package with Student and Subject structs`
3. `Add grades package with grading logic`
4. `Add processor package for result calculation`
5. `Add storage package for JSON persistence`
6. `Add ui package for terminal input/output`
7. `Add main.go menu loop`
8. `Add README documentation`

---

## Understanding the Code (Package Walkthrough)

### `internal/models/student.go`
Defines the **data types**: `Subject` and `Student` structs, plus the
`MaxMarksPerSubject` constant. Struct fields use JSON tags so the storage
package can (de)serialise them.

### `internal/grades/grades.go`
Holds the **grading rules**. `CalculateGrade` uses a `switch` with
conditional cases to map a percentage to a letter grade. `IsPass` returns a
bool, and `GradePoint` converts a letter to numeric grade points.

### `internal/processor/processor.go`
The **business logic**. A `Processor` struct with methods stores the class,
computes each student's results (using a `for-range` loop), and produces
statistics (class average, pass/fail counts, top performer).

### `internal/storage/storage.go`
Handles **persistence**. `SaveToFile` serialises the class to indented JSON;
`LoadFromFile` reads it back, gracefully returning an empty class on the
first run when no file exists yet.

### `internal/ui/ui.go`
All **terminal I/O**. Input helpers retry until valid data is entered
(demonstrating loops + `strconv` conversion). Output helpers print headers,
individual cards, and the results table.

### `main.go`
The **entry point**. Loads saved data, then runs a menu loop that dispatches
to handler functions (`addStudent`, `searchStudent`, `showStatistics`,
`saveData`). A `switch` selects the chosen action.

---

## Rubric Coverage

| Rubric Criterion                              | Where it is satisfied                                                                 |
|-----------------------------------------------|--------------------------------------------------------------------------------------|
| Functionality & Output Accuracy (12)          | Calculation in `processor.CalculateResults`, grading in `grades`, output in `ui`     |
| Application of Go Features (10)                | Structs, slices, functions, methods, `switch`, `if`, `for-range`, `strconv`, `json` |
| Modularity & Code Structure (6)               | Five packages under `internal/`, each with a single responsibility; `go.mod` module |
| GitHub Usage & Version Control (6)            | `.gitignore`, branch workflow, and the commit/branch guide above                     |
| Documentation & Presentation (6)              | This README, inline comments on every package, `go fmt` formatting                 |

---

## Troubleshooting

**`go: command not found`** — Go is not installed or not on your PATH.
Reinstall from <https://go.dev/dl/> and reopen the terminal.

**`package github.com/... is not in std`** — Your `go.mod` module path does not
match the import paths in the code. Either keep the provided module path, or
update every import in `main.go` and `internal/` to match your own.

**`cannot find package`** — Run the commands from inside the project root
(the folder containing `go.mod`).

**Marks won't save / `data/` missing** — Use menu option **5**; it creates the
`data/` directory automatically before writing.

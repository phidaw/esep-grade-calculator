# ESEP Grade Calculator

This library is a simple utility to calculate your final grade for a course.

How to use the library
- To use the library, include its module (see path in go.mod) into your workspace and import the module into files that reference it.
- Then create an instance of the GradeCalculator struct with NewGradeCalculator(). This can be used to calculate the final grades for a single course.
- To get started, add grades to an instance with the AddGrade() method. Then, you can call an instance's GetFinalGrade() method to retrieve the final grade for the course, given as a letter grade.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade-calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade-calculator/
# collect a coverage profile
go test -coverprofile=coverage.txt ./...
# generate a report
go tool cover -func=coverage.txt
```
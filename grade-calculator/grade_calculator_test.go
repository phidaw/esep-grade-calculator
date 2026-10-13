package esepunittests

import "testing"

// ----------- GradeType.String() -----------

func TestGradeTypeToStringAssignment(t *testing.T) {
	expected_value := "assignment"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)

	actual_value := gradeCalculator.assignments[0].Type.String()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeToStringExam(t *testing.T) {
	expected_value := "exam"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("exam 1", 100, Exam)

	actual_value := gradeCalculator.exams[0].Type.String()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeToStringEssay(t *testing.T) {
	expected_value := "essay"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.essays[0].Type.String()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// ------- GetFinalGrade() -------

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 79, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 68, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 62, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)
	gradeCalculator.AddGrade("exam 1", 57, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

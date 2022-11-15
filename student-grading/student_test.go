package main

import (
	"reflect"
	"testing"

)

func TestCalculateScoreAndGrade(t *testing.T) {
	stud := student{"Aanam", "Shaikh", "Mumbai University", 78, 59, 69, 59}
	actualString := getAverageScore(stud)
	expectedString := 66.25
	if actualString != expectedString {
		t.Errorf("Expected String(%v) is not same as"+
			" actual string (%v)", expectedString, actualString)
	}
}

func TestCalculateGrade(t *testing.T) {
    // assert := assert.New(t)
	stud := student{"Aanam", "Shaikh", "Mumbai University", 78, 59, 69, 59}
	avg := getAverageScore(stud)
	actualString := calculateGrade(avg)
	expectedString := "C"

    // assert.Equal{actualString,expectedString,"Grades are not equal"}
	if actualString != expectedString {
		t.Errorf("Expected String(%v) is not same as"+
			" actual string (%v)", expectedString, actualString)
	}
}

func TestGetOverallTopper(t *testing.T) {

	students := make([]studentGrade, 0)
	students = append(students, studentGrade{student{"Aanam", "Shaikh", "Mumbai University", 99, 99, 99, 99}, 99.0, "A"})
	students = append(students, studentGrade{student{"Aanam", "Shaikh", "Mumbai University", 78, 59, 69, 59}, 0, "F"})
	students = append(students, studentGrade{student{"Aanam", "Shaikh", "Mumbai University", 78, 59, 69, 59}, 0, "F"})
	students = append(students, studentGrade{student{"Aanam", "Shaikh", "Mumbai University", 78, 59, 69, 59}, 0, "F"})

	topper := overallTopper(students)
	expectedTopper := studentGrade{student{"Aanam", "Shaikh", "Mumbai University", 99, 99, 99, 99}, 99.0, "A"}

	if topper != expectedTopper {
		t.Errorf("The expected topper is: %v, but got %v", expectedTopper, topper)
	}
}

func TestUniversityTopper(t *testing.T) {

	students := make([]studentGrade, 0)
	students = append(students, studentGrade{student{"Aanam", "Shaikh", "Mumbai University", 99, 99, 99, 99}, 99.0, "A"})
	students = append(students, studentGrade{student{"Aanam", "Shaikh", "Mumbai University", 78, 59, 69, 59}, 66.25, "C"})
	students = append(students, studentGrade{student{"Aanam", "Shaikh", "Pune University", 78, 59, 69, 59}, 66.25, "C"})
	students = append(students, studentGrade{student{"Aanam", "Shaikh", "Delhi University", 78, 59, 69, 59}, 66.25, "C"})

	topperStudent := make(map[string]studentGrade)
	topperStudent["Mumbai University"] = studentGrade{student{"Aanam", "Shaikh", "Mumbai University", 99, 99, 99, 99}, 99.0, "A"}
	topperStudent["Pune University"] = studentGrade{student{"Aanam", "Shaikh", "Pune University", 78, 59, 69, 59}, 66.25, "C"}
	topperStudent["Delhi University"] = studentGrade{student{"Aanam", "Shaikh", "Delhi University", 78, 59, 69, 59}, 66.25, "C"}

	topperPerUniversity := topperPerUniversity(students)

	if !reflect.DeepEqual(topperStudent, topperPerUniversity) {
		t.Errorf("Expected value %v, but got %v", topperStudent, topperPerUniversity)
	}
}

func TestParseCSV(t *testing.T) {
	actualList := make([]student, 0)
	actualList = append(actualList, student{"Aanam", "Shaikh", "Mumbai University", 45, 56, 78, 30})
	actualList = append(actualList, student{"Tom", "Fellon", "Pune University", 89, 56, 78, 34})
	actualList = append(actualList, student{"jane", "Doe", "Delhi University", 55, 96, 90, 37})
	actualList = append(actualList, student{"Tom", "smith", "Mumbai University", 59, 96, 78, 37})
	actualList = append(actualList, student{"John", "Doe", "Delhi University",55, 96, 98, 37})
	actualList = append(actualList, student{"John", "smith", "Delhi University",55, 76, 99, 97})

	expectedList := parseCSV("student-data.csv")
	
    if actualList[0]!=expectedList[0] {
        t.Errorf("Expected  Student(%v) is not same as"+
			" actual Student (%v)", expectedList[0], actualList[0])
    }

    if actualList[len(actualList)-1]!=expectedList[len(expectedList)-1] {
        t.Errorf("Expected Student(%v) is not same as"+
			" actual Student (%v)",expectedList[len(expectedList)-1],actualList[len(actualList)-1])
    }

    if len(actualList)!=len(expectedList){
        t.Errorf("Expected size(%v) is not same as"+
			" actual size (%v)",len(expectedList) ,len(expectedList))
    }

}

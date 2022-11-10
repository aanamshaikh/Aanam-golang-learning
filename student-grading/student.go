package main

import (
	"fmt"
)

type student struct {
	firstname  string
	lastname   string
	university string
	test1Score float64
	test2Score float64
	test3Score float64
	test4Score float64

	// finalScore float64
	// grade string
}

func (student student) getAverageScore() float64 {
	average := (student.test1Score + student.test2Score + student.test3Score + student.test4Score) / 4
	return float64(average)
}

func calculateGrade(average float64) string {

	if average > 90 {

		return "A"
	} else if average > 80 && average < 90 {
		return "B"
	} else if average > 60 && average < 80 {
		return "C"
	} else if average > 35 && average < 60 {
		return "D"
	} else {
		return "F"
	}

}

func overallTopper(students []student) student {
	topper := student{"", "", "", 0, 0, 0, 0}
	var highestScore float64

	for i := 0; i < len(students); i++ {

		if students[i].getAverageScore() > highestScore {
			topper = students[i]
		}

	}
	return topper
}

var universityTopper = map[string]student{}
func topperPerUniversity(students []student) {

	topper := student{"", "", "", 0, 0, 0, 0}
	fmt.Println(topper)
	var maxScore float64 = 0

	for j := 0; j < len(students); j++ {
		if students[j].university == "Mumbai University" {

			if students[j].university == "Mumbai University" && maxScore < students[j].getAverageScore() {
				topper = students[j]
			}
			universityTopper["Mumbai University"]=topper
			
		}

		if students[j].university == "Delhi University" {

			if students[j].university == "Delhi University" && maxScore < students[j].getAverageScore() {
				topper = students[j]
			}
			universityTopper["Delhi University"]=topper
			
		}
		if students[j].university == "Pune University" {

			if students[j].university == "Pune University" && maxScore < students[j].getAverageScore() {
				topper = students[j]
			}
			universityTopper["Pune University"]=topper
		}

	}
	
}

func main() {

	var studentsList []student
	studentsList = append(studentsList, student{"Aanam", "Shaikh", "Mumbai University", 45, 56, 78, 35})
	studentsList = append(studentsList, student{"Harry", "Styles", "Delhi University", 89, 13, 49, 65})
	studentsList = append(studentsList, student{"Tom", "Gilford", "Mumbai University", 90, 13, 45, 55})
	studentsList = append(studentsList, student{"John", "Smith", "Pune University", 98, 78, 78, 89})

	// for _, studentsList := range studentsList {
	// 	finalScore := studentsList.getAverageScore()
	// 	grade := calculateGrade(finalScore)
	// 	fmt.Printf("The final Score of Student %v %v is %v and the grade is %v\n", studentsList.firstname, studentsList.lastname, finalScore, grade)
	// }
	// topper := overallTopper(studentsList)
	// fmt.Printf("The overall Topper is %v %v with the score %v", topper.firstname, topper.lastname, topper.getAverageScore())

	topperPerUniversity(studentsList)
	fmt.Println(universityTopper)

}

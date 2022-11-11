package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
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
			universityTopper["Mumbai University"] = topper

		}

		if students[j].university == "Delhi University" {

			if students[j].university == "Delhi University" && maxScore < students[j].getAverageScore() {
				topper = students[j]
			}
			universityTopper["Delhi University"] = topper

		}
		if students[j].university == "Pune University" {

			if students[j].university == "Pune University" && maxScore < students[j].getAverageScore() {
				topper = students[j]
			}
			universityTopper["Pune University"] = topper
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

	// topperPerUniversity(studentsList)
	// fmt.Println(universityTopper)

	// read from csv and set to array
	// read from csv file and set the universities to array making it dynamic

	csvFile, _ := os.Open("student-data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var students []student
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}


		a, _ := strconv.ParseFloat(line[3], 64)
        fmt.Println(a,reflect.TypeOf(a))
		students = append(students, student{
			firstname:  line[0],
			lastname:   line[1],
			university: line[2],
			test1Score: a,
			// test2Score: line[4],
			// test3Score: line[5],
			// test4Score: line[6],
		})
		fmt.Println(students)

		
	}
}

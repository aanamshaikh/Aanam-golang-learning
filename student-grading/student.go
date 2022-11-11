package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

func ParseToFloat(number string) float64 {
	num, _ := strconv.ParseInt(number, 10, 64)
	// fmt.Println(number,num)
	return float64(num)
}

func parseCSV(filepath string) []student {
	csvFile, _ := os.Open(filepath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var students []student
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		students = append(students, student{
			firstname:  line[0],
			lastname:   line[1],
			university: line[2],
			test1Score: ParseToFloat(line[3]),
			test2Score: ParseToFloat(line[4]),
			test3Score: ParseToFloat(line[5]),
			test4Score: ParseToFloat(line[6]),
		})
	}
	return students
}

func getScoreAndGradeForStudent(students []student) {
	for _, students := range students {
		finalScore := students.getAverageScore()
		grade := calculateGrade(finalScore)
		fmt.Printf("The final Score of Student %v %v is %v and the grade is %v\n", students.firstname, students.lastname, finalScore, grade)
	}
}

func topperPerUniversity(students []student) map[string]student {
	studentList := make(map[string]student)
	listOfStudents:=students
	 sort.Slice(listOfStudents, func(i, j int) bool {
		return listOfStudents[i].university < listOfStudents[j].university
	})
	
	i:=0
	lastIndexUniv:=0

	for i=0;i<len(listOfStudents);i++{ 
			if listOfStudents[lastIndexUniv].university!=listOfStudents[i].university {
				lastIndexUniv=i
			}
			studentList[listOfStudents[i].university]=overallTopper(listOfStudents[lastIndexUniv:i])
	}
	studentList[listOfStudents[lastIndexUniv].university]=overallTopper(listOfStudents[lastIndexUniv:])
	
	return studentList
}

func main() {
	var studentsList []student = parseCSV("student-data.csv")

	getScoreAndGradeForStudent(studentsList)

	topper := overallTopper(studentsList)
	fmt.Printf("The overall Topper is %v %v with the score %v", topper.firstname, topper.lastname, topper.getAverageScore())
   
	topperPerUniversity(studentsList)
	fmt.Println("Students: ",topperPerUniversity(studentsList))
}

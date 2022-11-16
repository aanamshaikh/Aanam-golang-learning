package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	// "sort"
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
}

type studentGrade struct {
	student    student
	finalScore float64
	grade      string
}

func (s student) string() string {
	return fmt.Sprintf("%v %v %v %v %v %v %v", s.firstname, s.lastname, s.university, s.test1Score, s.test2Score, s.test2Score, s.test4Score)
}

func (sg studentGrade) string() string {
	return fmt.Sprintf("%v %v %v", sg.student.string(), sg.finalScore, sg.grade)
}
func getAverageScore(student student) float64 {
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

func parseCSV(filepath string) []student {
	csvFile, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return make([]student, 0)
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var students []student
	for {
		line, err := reader.Read()
		// refactor
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
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

func getScoreAndGrades(students []student) []studentGrade {

	studentGrades := make([]studentGrade, 0)
	for _, s := range students {
		finalScore := getAverageScore(s)
		grade := calculateGrade(finalScore)
		sg := studentGrade{student: s, finalScore: finalScore, grade: grade}
		studentGrades = append(studentGrades, sg)

	}
	return studentGrades
}

func overallTopper(students []studentGrade) studentGrade {
	var topper studentGrade

	if len(students) < 1 {
		return topper
	}

	topper = students[0]

	for i := 0; i < len(students); i++ {
		if topper.finalScore < students[i].finalScore {
			topper = students[i]
		}
	}
	return topper
}

func ParseToFloat(number string) float64 {
	num, _ := strconv.ParseInt(number, 10, 64)
	return float64(num)
}

func topperPerUniversity(students []studentGrade) map[string]studentGrade {
	topperPerUniversity := make(map[string]studentGrade, 0) // mum

	for _, s := range students { // Mum ,pune , delhi

		student, ok := topperPerUniversity[s.student.university] //mumbai
		if ok {
			if s.finalScore > student.finalScore {
				topperPerUniversity[s.student.university] = s
				continue
			}
		} 
		topperPerUniversity[s.student.university] = studentGrade{student: s.student, finalScore: s.finalScore, grade: s.grade}
		
	}
	return topperPerUniversity
}

func main() {
	var studentsList []student = parseCSV("student-data.csv")

	studentGrades := getScoreAndGrades(studentsList)
	for _, s := range studentGrades {

		fmt.Println(s.string())
		// fmt.Printf("student %v %v has secured %v marks with %v grade\n", s.student.firstname, s.student.lastname, s.finalScore, s.grade)
	}

	topper := overallTopper(studentGrades)
	fmt.Printf("The overall Topper is %v %v with the score %v \n", topper.student.firstname, topper.student.lastname, topper.finalScore)

	students := topperPerUniversity(studentGrades)
	for _, s := range students {
		fmt.Printf("The topper for %v is %v %v with the score %v\n", s.student.university, s.student.firstname, s.student.lastname, s.finalScore)
	}

}

package main

import "fmt"

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

func getAverageScore(student student) float64 {
average :=(student.test1Score + student.test2Score + student.test3Score +student.test4Score) /4
return float64(average)
}

func calculateGrade(average float64) string{

    if average > 90 {

        return "A"
    }else if average > 80 && average < 90 {
        return "B"
    }else if average >60  && average < 80 {
        return "C"
    }else if average >35  && average < 60 {
        return "D"
    }else{
        return "F"
    }
    
}

func overallTopper(students student) {

}
func main() {

	var studentsList []student
	studentsList = append(studentsList, student{"Aanam", "Shaikh", "Mumbai University", 45, 56, 78, 35})
	studentsList = append(studentsList, student{"Harry", "Styles", "Delhi University", 89, 13, 49, 65})
	studentsList = append(studentsList, student{"Tom", "Gilford", "Mumbai University", 90, 13, 45, 55})
	studentsList = append(studentsList, student{"John", "Smith", "Pune University", 98, 78, 78, 89})
	fmt.Println(studentsList)
    
    for _,studentsList:=range studentsList {
		var finalScore float64
		finalScore = getAverageScore(studentsList)
		grade := calculateGrade(finalScore)
        // fmt.Println(studentsList.firstname,grade)
        fmt.Printf("The final Score of Student %v %v is %v and the grade is %v\n",studentsList.firstname,studentsList.lastname,finalScore,grade)
	}

    

}

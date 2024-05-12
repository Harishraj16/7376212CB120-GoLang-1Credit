package main

import "fmt"

// Define a struct for student
type Student struct {
    ID     int
    Name   string
    Age    int
    Courses []string
}

func main() {
    courses := []string{"Math", "Science", "History", "English"}

    studentDB := make(map[int]Student)

    studentDB[101] = Student{ID: 101, Name: "Harish", Age: 20, Courses: []string{"Math", "Science"}}
    studentDB[102] = Student{ID: 102, Name: "Lokesh", Age: 22, Courses: []string{"History", "English"}}
    studentDB[103] = Student{ID: 103, Name: "Kamal", Age: 21, Courses: []string{"Science", "English"}}

    for id, student := range studentDB {
        fmt.Printf("Student ID: %d\n", id)
        fmt.Printf("Name: %s\n", student.Name)
        fmt.Printf("Age: %d\n", student.Age)
        fmt.Printf("Courses Enrolled: %v\n", student.Courses)
        fmt.Println("-------------------------")
    }

    fmt.Println("Available Courses:")
    for i, course := range courses {
        fmt.Printf("%d. %s\n", i+1, course)
    }
}

package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	rollnumber  int
	name        string
	address     string
	Studcourses []string
}

func NewStudent(rollno int, name string, address string, courses []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.Studcourses = courses
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, courses []string) *Student {
	st := NewStudent(rollno, name, address, courses)
	ls.list = append(ls.list, st)
	return st
}

func CalculateHash(stringToHash string) string {
	fmt.Println("String Recieved: ", stringToHash)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))

}


func (li *StudentList) Print() {
	no := 0
	max := len((li.list))
	for no < max {
		fmt.Println(strings.Repeat("=", 25), " List  ", no, strings.Repeat("=", 25), "\nstudent rollno:", li.list[no].rollnumber, "\nstudent name: ", li.list[no].name)
		fmt.Println("student address:", li.list[no].address)
		fmt.Println("student courses:", li.list[no].Studcourses, "\n")

		str := strconv.Itoa(li.list[no].rollnumber)
		str = str + li.list[no].address + li.list[no].name

		for j := 0; j < len(li.list[no].Studcourses); j++ {
			str = str + li.list[no].Studcourses[j]
		}
		//CalculateHash(str)
		output := CalculateHash(str)
		fmt.Println("\nHash: \n", output,"\n")
		//fmt.Println(str)

		no++
	}
}


func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAA", []string{"Web Programming", "Statistics"})
	student.CreateStudent(25, "Naveed", "BBBBBB", []string{"Calculus", "Chemistry"})

	student.Print()
	
}

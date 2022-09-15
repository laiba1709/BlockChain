package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")

	student.Print()
}

func (li *StudentList) Print() {

	no := 0
	
	max := len((li.list))
	for no < max {
	fmt.Println(strings.Repeat("=", 25)," List  ", no, strings.Repeat("=", 25),"\nstudent rollno:",li.list[no].rollnumber, "\nstudent name: ",li.list[no].name)
	fmt.Println("student address:",li.list[no].address,"\n")
	no ++
	}


}
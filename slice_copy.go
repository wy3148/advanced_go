package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {

	//go is using value copy
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, v.Name)
	}
}

func main() {
	pase_student()
}

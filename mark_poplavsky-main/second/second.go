package second

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"time"
)

type Student struct { // Our example struct, you can use "-" to ignore a field
	SureName    string `csv:"фамилия"`
	Name        string `csv:"имя"`
	DateOfBirth string `csv:"дата рождения"`
}

func Second() {
	classesFile, err := os.OpenFile("students.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer classesFile.Close()

	classes := make([]Student, 0)

	if err := gocsv.UnmarshalFile(classesFile, &classes); err != nil { // Load clients from file
		panic(err)
	}

	minS := findMin(classes)
	maxS := findMax(classes)

	fmt.Println(minS.Name, minS.SureName, minS.DateOfBirth)
	fmt.Println(maxS.Name, maxS.SureName, maxS.DateOfBirth)
}

func findMax(students []Student) Student {
	result := students[0]
	for _, j := range students {
		value, err := time.Parse("02.01.2006", j.DateOfBirth)
		if err != nil {
			fmt.Println("все сломалось - ", err.Error())
		}
		valueAtResult, err := time.Parse("02.01.2006", result.DateOfBirth)
		if time.Since(value) > time.Since(valueAtResult) {
			result = j
		}
	}
	return result
}

func findMin(students []Student) Student {
	result := students[0]
	for _, j := range students {
		value, err := time.Parse("02.01.2006", j.DateOfBirth)
		if err != nil {
			fmt.Println("все сломалось - ", err.Error())
		}
		valueAtResult, err := time.Parse("02.01.2006", result.DateOfBirth)
		if time.Since(value) < time.Since(valueAtResult) {
			result = j
		}
	}
	return result
}

package main

import (
	"fmt"
	"github.com/LockedThread/ValenciaRateMyProfessor/model"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("example.html")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", "example.html")
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

	courseArray := ParseTable(string(data))

	courseInstructorMap := make(map[string]model.Course)

	for i := range courseArray {
		course := courseArray[i]
		courseInstructorMap[course.Instructor] = course
	}

	for _ = range courseInstructorMap {

	}
}

func ParseTable(data string) []model.Course {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	courses := make([]model.Course, 1)

	table := doc.Find("table")

	table.Find("tr").Each(func(i int, rowSelection *goquery.Selection) {
		if i < 2 {
			return
		}

		tableData := rowSelection.Find("td")

		courses = append(courses, GetCourse(tableData))
	})
	return courses
}

func GetCourse(tableData *goquery.Selection) (course model.Course) {
	cont := true
	tableData.Each(func(i int, selection *goquery.Selection) {
		ParseTableRow(&course, &cont, i, selection)
	})

	if len(course.Title) == 0 {
		return
	}

	return course
}

func ParseTableRow(course *model.Course, cont *bool, i int, selection *goquery.Selection) () {
	text := selection.Text()

	if i == 0 {
		if len(text) == 2 {
			*cont = false
			return
		} else {
			*cont = true
		}
	}
	if *cont == false {
		return
	}

	switch i {
	case 1:
		course.Course = text
		break
	case 2:
		course.Subject = text
		break
	case 3:
		course.CRN = text
		break
	case 4:
		break
	case 5:
		course.CampusID = text
		break
	case 6:
		atoi, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Fatal(err)
		}
		course.Credits = atoi
		break
	case 7:
		course.Title = text
		break
	case 16:
		course.Instructor = text
		break
	}
}

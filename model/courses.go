package model

import (
	"fmt"
)

var Campuses = map[string]string{
	"WC":  "West",
	"EC":  "East",
	"DTC": "Downtown",
	"LNC": "Lake Nona",
	"PC":  "Poinciana",
	"WP":  "Winter Park",
	"OC":  "Osceola",
}

type Course struct {
	CRN        string
	Subject    string
	Course     string
	Title      string
	CampusID   string
	Credits    int
	Honors     bool
	Instructor string
}

func (c Course) String() string {
	return fmt.Sprintf("CRN: %s, Subject: %s, Course: %s, Title: %s, CampusID: %s, Credits: %d, Honors: %t, Instructor: %s", c.CRN, c.Subject, c.Course, c.Title, c.CampusID, c.Credits, c.Honors, c.Instructor)
}

func (c Course) CampusName() {

}

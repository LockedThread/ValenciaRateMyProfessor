package model

import (
	"fmt"
	"log"
	"strings"
)

type FullName struct {
	lastName      string
	middleInitial string
	firstName     string
}

func (f FullName) FormattedString() string {
	if len(f.middleInitial) > 0 {
		return fmt.Sprintf("%s %s %s", f.firstName, f.middleInitial, f.lastName)
	}
	return fmt.Sprintf("%s %s", f.firstName, f.lastName)
}

func GetFullNameFromString(name string) FullName {
	name = strings.Replace(name, " (P)", "", 1)
	name = strings.TrimSpace(name)
	//nameLength := len(name)

	fullNameStruct := FullName{}

	split := strings.Split(name, " ")

	cursorIndex := 0
	for i := range split {
		cursor := split[i]
		if len(cursor) > 0 {
			switch cursorIndex {
			case 0:
				fullNameStruct.firstName = cursor
				break
			case 1:
				fullNameStruct.middleInitial = cursor
				break
			case 2:
				fullNameStruct.lastName = cursor
				break
			default:
				log.Fatalf("Name fragment index too long, %s\n", cursor)
			}
			cursor = ""
			cursorIndex++
		}
	}

	/*
		currentFragmentIndex := 0
		var cursor string
		for i := range name {
			character := name[i]

			currentIsSpace := character == ' '

			if currentIsSpace {
				if len(cursor) != 0 {
					switch currentFragmentIndex {
					case 0:
						fullNameStruct.firstName = cursor
						log.Println(cursor)
						break
					case 1:
						fullNameStruct.middleInitial = cursor
						break
					case 2:
						fullNameStruct.lastName = cursor
						break
					default:
						log.Fatalln("Name fragment index too long")
					}
					currentFragmentIndex++
					cursor = ""
				}
				continue
			}
			cursor += string(character)
		}*/
	return fullNameStruct
}

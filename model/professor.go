package model

import "fmt"

type Professor struct {
	FullName FullName
}

func (p Professor) FormattedString() string {
	return fmt.Sprintf("%s", p.FullName)
}

func (p Professor) String() string {
	return fmt.Sprintf("FullName: %s", p.FullName)
}

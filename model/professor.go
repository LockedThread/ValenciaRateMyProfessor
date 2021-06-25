package model

import "fmt"

type Professor struct {
	FullName FullName
}

func (p Professor) FormattedString() string {
	return p.FullName.FormattedString()
}

func (p Professor) String() string {
	return fmt.Sprintf("FullName: %s", p.FullName)
}

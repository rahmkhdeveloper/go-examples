package model

import "fmt"

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	initials  string
}

func (u *User) GetMyInitials() (byte, byte) {
	f := u.FirstName[0]
	l := u.LastName[0]
	return f, l
}

func (u *User) WriteMyInitials(F byte, L byte) {
	u.initials = fmt.Sprintf("%s. %s.", string(F), string(L))
}

func (u *User) WriteMyNameAndInitials() (output string) {
	output = fmt.Sprintf("My name is %s %s (%s)", u.FirstName, u.LastName, u.initials)
	return
}

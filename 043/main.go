package main

import "fmt"

type User struct {
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
	fmt.Println(u.Name)
}

type Em struct {
	User
}

func main() {
	e := new(Em)
	e.SetName("22")
}

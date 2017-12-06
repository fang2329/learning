package test

import (
	"fmt"
	"testing"
)

type user struct {
	name string
	age  uint32
}

func (u *user) SetInfo(name string, age uint32) {
	u.name = name
	u.age = age
}

func (u *user) GetAge() uint32 {
	return u.age
}

func TestInfo(t *testing.T) {
	u := user{"Bob", 10}
	u.SetInfo("Alice", 23)
	age := u.GetAge()
	fmt.Println(age)
}

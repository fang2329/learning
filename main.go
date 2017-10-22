package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type User struct {
	Name int32
	Age  int32
}

type Admin struct {
	User
	Sex int32
}

func (u *User) Encode() []byte {

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, u)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return buf.Bytes()
}

func (u *User) Decode(b []byte) {
	buf := bytes.NewBuffer(b)

	err := binary.Read(buf, binary.LittleEndian, u)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
}

func main() {
	a := Admin{
		User: User{
			Name: 2001,
			Age:  23,
		},
		Sex: 1,
	}
	buf := a.User.Encode()
	a.User.Decode(buf)
	fmt.Println(buf)
	fmt.Println(a)

	buff := a.Encode()
	a.Decode(buff)
	fmt.Println(buff)
	fmt.Println(a)
}

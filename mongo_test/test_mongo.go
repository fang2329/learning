package main

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserInfo struct {
	Id   int
	Name string
	Age  int
	Sex  string
	Addr string
	Tel  string
}

type User struct {
	Users []UserInfo
}

const (
	URL = "127.0.0.1:27017"
)

var (
	ErrConnect  = errors.New("connect filed")
	ErrorHandle = errors.New("handle is nil")
	ErrorInsert = errors.New("insert failed")
	ErrorFind   = errors.New("find failed")
	ErrorUpdate = errors.New("update failed")
	ErrorRemove = errors.New("remove failed")
)

func new(id int, name string, age int, sex, addr, tel string) *UserInfo {
	return &UserInfo{
		Id:   id,
		Name: name,
		Age:  age,
		Sex:  sex,
		Addr: addr,
		Tel:  tel,
	}
}

func main() {
	session, err := mgo.Dial(URL)
	if err != nil {
		fmt.Println(ErrConnect)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("testdb")
	collection := db.C("user")

	user1 := new(1, "alice", 10, "woman", "123", "456")
	user2 := new(3, "BOB", 20, "man", "111", "222")
	user3 := new(2, "xiaoming", 23, "man", "111", "777")
	user4 := new(7, "xiaohua", 20, "woman", "777", "456")
	InserData(collection, user1, user2, user3, user4)

	user5 := UserInfo{}
	QueryOneData(collection, &user5)

	var userAll User
	var tmpuser UserInfo
	QueryManyData(collection, userAll, tmpuser)

	UpdateNameDataById(collection, 2, "xiaohong")
	UpdateNameDataById(collection, 3, "xiaolan")
	UpdateNameDataById(collection, 7, "mongo")
	var userAll2 User
	var tmpuser2 UserInfo
	QueryManyData(collection, userAll2, tmpuser2)

	UpdateData(collection, 2, "111")
	UpdateData(collection, 3, "222")
	UpdateData(collection, 7, "333")
	var userAll3 User
	var tmpuser3 UserInfo
	QueryManyData(collection, userAll3, tmpuser3)

	Removedata(collection, 1)
	var userAll4 User
	var tmpuser4 UserInfo
	QueryManyData(collection, userAll4, tmpuser4)

}

//insert one or more date
func InserData(c *mgo.Collection, users ...*UserInfo) {
	if c == nil {
		fmt.Println(ErrorHandle)
	}
	for _, user := range users {
		err := c.Insert(user)
		if err != nil {
			fmt.Println(ErrorInsert)
		}
	}

}

//query one
func QueryOneData(c *mgo.Collection, user *UserInfo) {
	err := c.Find(bson.M{"id": 1}).One(user)
	if err != nil {
		fmt.Println(ErrorFind)
	} else {
		fmt.Println(*user)
	}
}

//query many data
func QueryManyData(c *mgo.Collection, user User, oneuser UserInfo) {
	iter := c.Find(nil).Iter()
	for iter.Next(&oneuser) {
		fmt.Println("result: ", oneuser)
		user.Users = append(user.Users, oneuser)
	}
}

//update
func UpdateNameDataById(c *mgo.Collection, id int, name string) {
	err := c.Update(bson.M{"id": id}, bson.M{"$set": bson.M{"name": name}})
	if err != nil {
		fmt.Println(ErrorUpdate)
	}
}

func UpdateData(c *mgo.Collection, id int, telp string) {
	err := c.Update(bson.M{"id": id}, bson.M{"tel": telp})
	if err != nil {
		fmt.Println(ErrorUpdate)
	}
}

//remove
func Removedata(c *mgo.Collection, id int) {
	_, err := c.RemoveAll(bson.M{"id": id})
	if err != nil {
		fmt.Println(ErrorRemove)
	}
}

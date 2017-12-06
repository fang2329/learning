package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	server = "127.0.0.1:11211"
)

func main() {
	//create a handle
	mc := memcache.New(server)
	if mc == nil {
		fmt.Println("memcache New failed")
	}

	//set key-value
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	//get key's value
	it, _ := mc.Get("foo")
	if string(it.Key) == "foo" {
		fmt.Println("value is ", string(it.Value))
	} else {
		fmt.Println("Get failed")
	}
	///Add a new key-value
	mc.Add(&memcache.Item{Key: "foo", Value: []byte("bluegogo")})
	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println("Add failed")
	} else {
		if string(it.Key) == "foo" {
			fmt.Println("Add value is ", string(it.Value))
		} else {
			fmt.Println("Get failed")
		}
	}
	//replace a key's value
	mc.Replace(&memcache.Item{Key: "foo", Value: []byte("mobike")})
	it, err = mc.Get("foo")
	if err != nil {
		fmt.Println("Replace failed")
	} else {
		if string(it.Key) == "foo" {
			fmt.Println("Replace value is ", string(it.Value))
		} else {
			fmt.Println("Replace failed")
		}
	}
	//delete an exist key
	err = mc.Delete("foo")
	if err != nil {
		fmt.Println("Delete failed:", err.Error())
	}
	//incrby
	err = mc.Set(&memcache.Item{Key: "aaa", Value: []byte("1")})
	if err != nil {
		fmt.Println("Set failed :", err.Error())
	}
	it, err = mc.Get("foo")
	if err != nil {
		fmt.Println("Get failed ", err.Error())
	} else {
		fmt.Println("src value is:", it.Value)
	}
	value, err := mc.Increment("aaa", 7)
	if err != nil {
		fmt.Println("Increment failed")
	} else {
		fmt.Println("after increment the value is :", value)
	}
	//decrby
	value, err = mc.Decrement("aaa", 4)
	if err != nil {
		fmt.Println("Decrement failed", err.Error())
	} else {
		fmt.Println("after decrement the value is ", value)
	}
}

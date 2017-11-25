package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

/***********************************************************************
hset(key, field, value)：向名称为key的hash中添加元素field
hget(key, field)：返回名称为key的hash中field对应的value
hmget(key, (fields))：返回名称为key的hash中field i对应的value
hmset(key, (fields))：向名称为key的hash中添加元素field
hincrby(key, field, integer)：将名称为key的hash中field的value增加integer
hexists(key, field)：名称为key的hash中是否存在键为field的域
hdel(key, field)：删除名称为key的hash中键为field的域
hlen(key)：返回名称为key的hash中元素个数
hkeys(key)：返回名称为key的hash中所有键
hvals(key)：返回名称为key的hash中所有键对应的value
hgetall(key)：返回名称为key的hash中所有的键（field）及其对应的value
************************************************************************/

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connect failed", err.Error())
	}
	fmt.Println(reflect.TypeOf(c))
	defer c.Close()
	/////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("hset", "myhash", "bike1", "mobike")
	if err != nil {
		fmt.Println("haset failed", err.Error())
	}
	//////////////////////////////////////////////////////////////////////////////
	res, err := c.Do("hget", "myhash", "bike1")
	fmt.Println(reflect.TypeOf(res))
	if err != nil {
		fmt.Println("hget failed", err.Error())
	} else {
		fmt.Printf("hget value :%s\n", res.([]byte))
	}
	//////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("hmset", "myhash", "bike2", "bluegogo", "bike3", "xiaoming", "bike4", "xiaolan")
	if err != nil {
		fmt.Println("hmset error", err.Error())
	} else {
		value, err := redis.Values(c.Do("hmget", "myhash", "bike1", "bike2", "bike3", "bike4"))
		if err != nil {
			fmt.Println("hmget failed", err.Error())
		} else {
			fmt.Printf("hmget myhash's element :")
			for _, v := range value {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Printf("\n")
		}
	}
	////////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("hset", "myhash", "tmpnum", 20)
	if err != nil {
		fmt.Println("before hincrby failed", err.Error())
	} else {
		_, err = c.Do("hincrby", "myhash", "tmpnum", 10)
		if err != nil {
			fmt.Println("hincrby failed", err.Error())
		} else {
			res, err = c.Do("hget", "myhash", "tmpnum")
			if err != nil {
				fmt.Println("after hincrby failed", err.Error())
			} else {
				fmt.Printf("after hincrby value now is :%s", res.([]byte))

			}
		}
	}
	/////////////////////////////////////////////////////////////////////////////////
	isExist, err := c.Do("hexists", "myhash", "tmpnum")
	if err != nil {
		fmt.Println("hexist failed", err.Error())
	} else {
		fmt.Println("exist or not:", isExist)
	}
	/////////////////////////////////////////////////////////////////////////////////
	ilen, err := c.Do("hlen", "myhash")
	if err != nil {
		fmt.Println("hlen failed", err.Error())
	} else {
		fmt.Println("myhash's len is :", ilen)
	}
	/////////////////////////////////////////////////////////////////////////////////
	resKeys, err := redis.Values(c.Do("hkeys", "myhash"))
	if err != nil {
		fmt.Println("hkeys failed", err.Error())
	} else {
		fmt.Printf("myhash's keys is :")
		for _, v := range resKeys {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
	///////////////////////////////////////////////////////////////////////////////
	resValues, err := redis.Values(c.Do("hvals", "myhash"))
	if err != nil {
		fmt.Println("hvals failed", err.Error())
	} else {
		fmt.Printf("myhash's values is:")
		for _, v := range resValues {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
	/////////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("HDEL", "myhash", "tmpnum")
	if err != nil {
		fmt.Println("hdel failed", err.Error())
	}

	///////////////////////////////////////////////////////////////////////////////
	result, err := redis.Values(c.Do("hgetall", "myhash"))
	if err != nil {
		fmt.Println("hgetall failed", err.Error())
	} else {
		fmt.Printf("all keys and values are:")
		for _, v := range result {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
	//////////////////////////////////////////////////////////////////////////////

}

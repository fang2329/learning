package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

/*****************************************************************
sadd(key, member)：向名称为key的set中添加元素member
srem(key, member) ：删除名称为key的set中的元素member
spop(key) ：随机返回并删除名称为key的set中一个元素
smove(srckey, dstkey, member) ：移到集合元素
scard(key) ：返回名称为key的set的基数
sismember(key, member) ：member是否是名称为key的set的元素
sinter(key1, key2,…key N) ：求交集
sinterstore(dstkey, (keys)) ：求交集并将交集保存到dstkey的集合
sunion(key1, (keys)) ：求并集
sunionstore(dstkey, (keys)) ：求并集并将并集保存到dstkey的集合
sdiff(key1, (keys)) ：求差集
sdiffstore(dstkey, (keys)) ：求差集并将差集保存到dstkey的集合
smembers(key) ：返回名称为key的set的所有元素
srandmember(key) ：随机返回名称为key的set的一个元素
******************************************************************/
func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connect failed", err.Error())
	}
	fmt.Println(reflect.TypeOf(c))
	defer c.Close()

	/////////////////////////////////////////////////////////////////
	_, err = c.Do("sadd", "myset", "mobike", "foo", "ofo", "bluegogo")
	if err != nil {
		fmt.Println("set add failed", err.Error())
	}
	value, err := redis.Values(c.Do("smembers", "myset"))
	if err != nil {
		fmt.Println("set get members failed", err.Error())
	} else {
		fmt.Printf("myset members :")
		for _, v := range value {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Printf("\n")
	}

	////////////////////////////////////////////////////////////////////////////
	ret, err := c.Do("srandmember", "myset")
	if err != nil {
		fmt.Println("srandmember get failed")
	} else {
		fmt.Printf("srandmember get value is:%s\n", ret)
	}
	////////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("srem", "myset", "bluegogo")
	if err != nil {
		fmt.Println("srem myset error", err.Error())
	} else {
		value, err = redis.Values(c.Do("smembers", "myset"))
		if err != nil {
			fmt.Println("set get members failed", err.Error())
		} else {
			fmt.Printf("myset members :")
			for _, v := range value {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Printf("\n")
		}
	}

	/////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("smove", "myset", "dbset", "mobike")
	if err != nil {
		fmt.Println("smove failed", err.Error())
	} else {
		value, err = redis.Values(c.Do("smembers", "dbset"))
		if err != nil {
			fmt.Println("get new dbset members failed", err.Error())
		} else {
			fmt.Printf("get new dbset members :")
			for _, v := range value {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Printf("\n")
		}
	}
	////////////////////////////////////////////////////////////////////////
	num, err := c.Do("scard", "myset")
	if err != nil {
		fmt.Println("scard error", err.Error())
	} else {
		fmt.Println("scard get num :", num)
	}
	/////////////////////////////////////////////////////////////////////////
	isMember, err := c.Do("sismember", "myset", "foo")
	if err != nil {
		fmt.Println("sismember get failed", err.Error())
	} else {
		fmt.Println("foo is or not myset's member", isMember)
	}
	/////////////////////////////////////////////////////////////////////////
	_, err = c.Do("sadd", "dbset", "foo", "ofo", "xiaolan")
	if err != nil {
		fmt.Println("set add failed", err.Error())
	}
	inner, err := redis.Values(c.Do("sinter", "myset", "dbset"))
	if err != nil {
		fmt.Println("sinter error", err.Error())
	} else {
		fmt.Printf("two set inter is:")
		for _, v := range inner {
			fmt.Printf(" %s ", v.([]byte))
		}
		fmt.Printf("\n")
	}
	//////////////////////////////////////////////////////////////////////////
	_, err = c.Do("sinterstore", "newset", "dbset", "myset")
	if err != nil {
		fmt.Println("sinterstore between myset and dbset error", err.Error())
	} else {
		value, err := redis.Values(c.Do("smembers", "newset"))
		if err != nil {
			fmt.Println("set get members failed", err.Error())
		} else {
			fmt.Printf("newset members :")
			for _, v := range value {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Printf("\n")
		}
	}
	////////////////////////////////////////////////////////////////////////////////
	unino, err := redis.Values(c.Do("sunion", "myset", "dbset"))
	if err != nil {
		fmt.Println("sunion err", err.Error())
	} else {
		fmt.Printf("two set union is:")
		for _, v := range unino {
			fmt.Printf(" %s ", v.([]byte))
		}
		fmt.Printf("\n")
	}
	////////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("sunionstore", "unewset", "myset", "dbset")
	if err != nil {
		fmt.Println("sunionstore failed", err.Error())
	} else {
		value, err := redis.Values(c.Do("smembers", "unewset"))
		if err != nil {
			fmt.Println("set get members failed", err.Error())
		} else {
			fmt.Printf("unewset members :")
			for _, v := range value {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Printf("\n")
		}
	}
	////////////////////////////////////////////////////////////////////////////////
	diff, err := redis.Values(c.Do("sdiff", "dbset", "myset"))
	if err != nil {
		fmt.Println("sdiff err", err.Error())
	} else {
		fmt.Printf("two set diff is:")
		for _, v := range diff {
			fmt.Printf(" %s ", v.([]byte))
		}
		fmt.Printf("\n")
	}
	////////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("sdiffstore", "dnewset", "dbset", "myset")
	if err != nil {
		fmt.Println("sdiffstore failed", err.Error())
	} else {
		value, err := redis.Values(c.Do("smembers", "dnewset"))
		if err != nil {
			fmt.Println("set get members failed", err.Error())
		} else {
			fmt.Printf("dnewset members :")
			for _, v := range value {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Printf("\n")
		}
	}
	/////////////////////////////////////////////////////////////////////////////
	res, err := c.Do("spop", "myset")
	if err != nil {
		fmt.Println("spop failed", err.Error())
	} else {
		fmt.Printf("spop element is:%s\n", res)
	}
	value, err = redis.Values(c.Do("smembers", "myset"))
	if err != nil {
		fmt.Println("after spop  get members failed", err.Error())
	} else {
		fmt.Printf("after spop myset members :")
		for _, v := range value {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Printf("\n")
	}
}

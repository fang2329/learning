package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

/**********************************************************************
zadd(key, score1,member1,...scoreN,memberN) 向有序结合添加（更新）一个或多个成员
zcard(key)：获取有序集合的成员
zcount(key,start,end):计算指定区间的成员数
zincrby(key,increment,member):成员member增加increment
zinterstore(dst,numkey,src1,src2..srcN)：求交集，并将结果存储新的结合
zlexcount(key,start,end):计算字典区间成员数(分数都相同，按照字典排序)
zrange(key,start,end):获取索引区间的成员
zrangebylex (key,start,end):通过字典区间返回区间内有序集合成员
zrangebyscore(key,start,end):通过分数返回区间内的有序集合
zrank (key,member):返回有序结合的索引
zrem(key,members1...membersN):删除一个或多个成员
zremrangebylex(key,start,end):移除集合中给定字典区间的成员
zremrangebyrank(key,start,end):移除有序集合中给定的排名区间的所有成员
zremrangebyscore(key,start,end)：移除给定分数区间的所有元素
zrevange(key,start,end):通过索引，分数由高到低，返回指定区域的元素
zrevrangebyscore(key,member)：分数由高向低返回指定区间的成员数
zrevrank(key,member):分数从小到大，返回指定成员的排名
zscore(key,member):返回有序集中，成员的分数值
zunionstore(dst,numkeys,key1...keyN):返回给定的一个或多个集合的并集，并存储在新的集合中
zscan(key,cursor):迭代有序结合中的元素（包括元素成员和元素分值）
***********************************************************************/
func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connect failed", err.Error())
	}
	fmt.Println(reflect.TypeOf(c))
	defer c.Close()
	//////////////////////////////////////////////////////////////
	_, err1 := c.Do("zadd", "curbike", 1, "mobike", 2, "xiaolan", 3, "ofo", 4, "xiaoming")
	_, err2 := c.Do("zadd", "tmpdata", 0, "mobike", 0, "xiaolan", 0, "mysql", 0, "redis", 0, "mongo", 0, "xiaoming")
	if err1 != nil || err2 != nil {
		fmt.Println("zadd failed", err.Error())
	}

	///////////////////////////所有元素个数////////////////////////////////////
	num, err := c.Do("zcard", "curbike")
	if err != nil {
		fmt.Println("zcard failed", err.Error())
	} else {
		fmt.Println("curbike's size is :", num)
	}
	//////////////////////////////区间个数///////////////////////////////////
	num, err = c.Do("zcount", "curbike", 1, 3)
	if err != nil {
		fmt.Println("zcount failed ", err.Error())
	} else {
		fmt.Println("zcount num is :", num)
	}
	////////////////////////////////////////////////////////////////
	num, err = c.Do("zincrby", "curbike", 3, "xiaolan")
	fmt.Println(reflect.TypeOf(num))
	if err != nil {
		fmt.Println("zincrby failed", err.Error())
	} else {
		fmt.Printf("after zincrby the :%s\n", num)
	}
	////////////////////////////////////////////////////////////////////
	_, err = c.Do("zinterstore", "internewset", 2, "curbike", "tmpdata")
	if err != nil {
		fmt.Println("zinterstore failed", err.Error())
	} else {
		result, err := redis.Values(c.Do("zrange", "internewset", 0, 10))
		if err != nil {
			fmt.Println("interstore failed", err.Error())
		} else {
			fmt.Printf("interstore newset elsements are:")
			for _, v := range result {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Println()
		}
	}
	//////////////////////////////////////////////////////////////////////
	num, err = c.Do("zlexcount", "tmpdata", "[mongo", "[xiaoming")
	if err != nil {
		fmt.Println("zlexcount failed", err.Error())
	} else {
		fmt.Println("zlexcount in tmpdata is :", num)
	}
	////////////////////////////////////////////////////////////////////////
	res, err := redis.Values(c.Do("zrange", "curbike", 0, -1, "withscores"))
	if err != nil {
		fmt.Println("zrange in curbike failed", err.Error())
	} else {
		fmt.Printf("curbike's element are follow:")
		for _, v := range res {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
	////////////////////////////////////////////////////////////////////////////
	res, err = redis.Values(c.Do("zrangebylex", "tmpdata", "[mobike", "[redis"))
	if err != nil {
		fmt.Println("zrangebylex failed", err.Error())
	} else {
		fmt.Printf("zrangebylex in tmpdata:")
		for _, v := range res {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
	////////////////////////////////////////////////////////////////////////////////
	res, err = redis.Values(c.Do("zrangebyscore", "curbike", "(1", "(5"))
	if err != nil {
		fmt.Println("zrangebyscore failed", err.Error())
	} else {
		fmt.Printf("zrangebyscore's element:")
		for _, v := range res {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
	////////////////////////////////////////////////////////////////////////////////
	num, err = c.Do("zrank", "internewset", "xiaoming")
	if err != nil {
		fmt.Println("zrank failed ", err.Error())
	} else {
		fmt.Println("xiaoming's score is ", num)
	}
	////////////////////////////////////////////////////////////////////////////////
	_, err = c.Do("zunionstore", "unewzset", 2, "curbike", "tmpdata")
	if err != nil {
		fmt.Println("zunionstore failed", err.Error())
	} else {
		res, err = redis.Values(c.Do("zrange", "unewzset", 0, 10))
		if err != nil {
			fmt.Println("zunionstore failed", err.Error())
		} else {
			fmt.Printf("union set are:")
			for _, v := range res {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Println()
		}
	}
	////////////////////////////////////////////////////////////////////////////////
	ret, err := c.Do("zscore", "internewset", "xiaolan")
	if err != nil {
		fmt.Println("zscore failed", err.Error())
	} else {
		fmt.Printf("curbike 's xiaolan score is:%s\n", ret)
	}
	////////////////////////////////////////////////////////////////////////////////
	num, err = c.Do("zrevrank", "curbike", "ofo")
	if err != nil {
		fmt.Println("zrevrank failed", err.Error())
	} else {
		fmt.Println("ofo's zrevrank is :", num)
	}
	///////////////////////////////////////////////////////////////////////////////
	res, err = redis.Values(c.Do("zrevrangebyscore", "unewzset", 10, 2))
	if err != nil {
		fmt.Println("zrevrangebyscore failed", err.Error())
	} else {
		fmt.Printf("zrevrangebyscore are:")
		for _, v := range res {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
	////////////////////////////////////////////////////////////////////////////////
	res, err = redis.Values(c.Do("zrevrange", "unewzset", 0, 10))
	if err != nil {
		fmt.Println("zrevrange failed：", err.Error())
	} else {
		fmt.Printf("zrevrange element:")
		for _, v := range res {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
	////////////////////////////////////////////////////////////////////////////////
	num, err = c.Do("zrem", "unewzset", "mysql")
	if err != nil {
		fmt.Println("zrem failed", err.Error())
	} else {
		fmt.Println("zrem result is:", num)
	}
	///////////////////////////////////////////////////////////////////////////////
	num, err = c.Do("zremrangebyrank", "unewzset", 1, 4)
	if err != nil {
		fmt.Println("zremrangebyrank failed", err.Error())
	} else {
		fmt.Println("zremrangebyrank result:", num)
	}
	////////////////////////////////////////////////////////////////////////////////
	num, err = c.Do("zremrangebyscore", "curbike", 2, 5)
	if err != nil {
		fmt.Println("zremrangebyscore failed", err.Error())
	} else {
		fmt.Println("zremrangebyscore result:", num)
	}
}

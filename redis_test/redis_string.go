package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)
failed redigo: unexpected type for String, got type int64
/******************************************************************************
set(key, value)：给数据库中名称为key的string赋予值value
get(key)：返回数据库中名称为key的string的value
getset(key, value)：给名称为key的string赋予上一次的value
mget(key1, key2,…, key N)：返回库中多个string的value
setnx(key, value)：添加string，名称为key，值为value
setex(key, time, value)：向库中添加string，设定过期时间time
mset(key N, value N)：批量设置多个string的值
msetnx(key N, value N)：如果所有名称为key i的string都不存在
incr(key)：名称为key的string增1操作
incrby(key, integer)：名称为key的string增加integer
decr(key)：名称为key的string减1操作
decrby(key, integer)：名称为key的string减少integer
append(key, value)：名称为key的string的值附加value
substr(key, start, end)：返回名称为key的string的value的子串
*******************************************************************************/
func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	defer c.Close()
////////////////////////////////////////////////////////////////////////
	_,err = c.Do("set","myname","Alice")
	if err != nil{
		fmt.Println("redis set failed",err.Error())
	}

	myname,err := redis.String(c.Do("GET","myname"))
	if err!= nil{
		fmt.Println("redis get failed",err.Error())
	}else{
		fmt.Println("GET myname",myname)
	}
//////////////////////////////////////////////////////////////////////////
	_,err = c.Do("set","mydb","redis")
        if nil != err{
                fmt.Println("set failed",err.Error())
        }
	mydb,err := redis.String(c.Do("GET","mydb"))
	if err == nil{
		fmt.Println("Get mydb",mydb)
	}
        var tmpname string
        var tmpdb string
        array,err :=redis.Values( c.Do("MGET","myname","mydb"))
        if err != nil{
                fmt.Println("mget failed",err.Error())
        }else{
                if _,err = redis.Scan(array,&tmpname,&tmpdb);err == nil{
                        fmt.Println("mget :",tmpname,tmpdb)
                }
        }
//////////////////////////////////////////////////////////////////////////

	_,err = c.Do("SET","myage",10,"EX","10")
	if nil != err{
		fmt.Println("redis set failed",err.Error())
	}

	isExist,err := redis.Bool(c.Do("EXISTS","myage"))
	if err != nil{
		fmt.Println("error",err.Error())
	}else{
		fmt.Println("exists or not ",isExist)
	}

	time.Sleep(11*time.Second)
	myage,err := redis.String(c.Do("GET","myage"))
	if err != nil{
		fmt.Println("redis get failed",err.Error())
	}else{
		fmt.Println("get myage",myage)
	}


        isExist,err = redis.Bool(c.Do("EXISTS","myage"))
        if err != nil{
                fmt.Println("error",err.Error())
        }else{
		fmt.Println("10s after exists or not ",isExist)
	}
//////////////////////////////////////////////////////////////////////////
        _,err = c.Do("DEL","myname")
	if err != nil{
		fmt.Println("redis delete failed",err.Error())
	}

	myname,err = redis.String(c.Do("GET","myname"))
	if err != nil{
		fmt.Printf("redis get failed %s  delete success\n",err.Error())
	}else{
		fmt.Println("delete failed")
	}
/////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("SET","myage",10)
	if err != nil{
		fmt.Println("set myage failed",err.Error())
	}

	myage,err = redis.String(c.Do("GET","myage"))
        if err == nil{
                fmt.Println("before myage",myage)
        }
	_,err=c.Do("INCR","myage")
	if err != nil{
		fmt.Println("incr error",err.Error())
	}
	myage,err = redis.String(c.Do("GET","myage"))
	if err == nil{
		fmt.Println("after myage incr 1",myage)
	}
	_,err=c.Do("INCRBY","myage",10)
	if err != nil{
		fmt.Println("incrby error",err.Error())
	}
	myage,err = redis.String(c.Do("GET","myage"))
        if err == nil{
                fmt.Println("after myage incr 10",myage)
        }
///////////////////////////////////////////////////////////////////////////////

        myage,err = redis.String(c.Do("GET","myage"))
        if err == nil{
                fmt.Println("before myage",myage)
        }
        _,err=c.Do("DECR","myage")
        if err != nil{
                fmt.Println("incr error",err.Error())
        }
        myage,err = redis.String(c.Do("GET","myage"))
        if err == nil{
                fmt.Println("after myage DECR 1",myage)
        }
        _,err=c.Do("DECRBY","myage",10)
        if err != nil{
                fmt.Println("incrby error",err.Error())
        }
        myage,err = redis.String(c.Do("GET","myage"))
        if err == nil{
                fmt.Println("after myage DECR 10",myage)
        }
///////////////////////////////////////////////////////////////////////////////
	mydb,err = redis.String(c.Do("GETSET","mydb","mysql"))
	if err != nil{
		fmt.Println("getset failed",err.Error())
	}
	fmt.Println("get before",mydb)
	mydb,err = redis.String(c.Do("get","mydb"))
	if err == nil{
		fmt.Println("mydb:",mydb)
	}
///////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("setnx","foo","bluegogo")
	if err != nil{
		fmt.Println("setnx failed",err.Error())
	}
	foo,err := redis.String(c.Do("get","foo"))
	if err == nil{
		fmt.Println("setnx value",foo)
	}
//////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("setex","ofo",10,"xiaoming")
	if err != nil{
		fmt.Println("setex failed ",err.Error())
	}
	ofo,err := redis.String(c.Do("get","ofo"))
	if err == nil{
		fmt.Println("setex ofo",ofo)
	}
	time.Sleep(11 *time.Second)
	ofo,err = redis.String(c.Do("get","ofo"))
        if err == nil{
                fmt.Println("setex ofo",ofo)
        }else{
		fmt.Println("time out get nil")
	}
/////////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("msetnx","bike1","ofo","bike2","bluegogo","bike3","foo")
	if err != nil{
		fmt.Println("setnx failed")
	}
	var tmpbike1 string
        var tmpbike2 string
        var tmpbike3 string
        array,err =redis.Values( c.Do("MGET","bike1","bike2","bike3"))
        if err != nil{
                fmt.Println("mget failed",err.Error())
        }else{
                if _,err = redis.Scan(array,&tmpbike1,&tmpbike2,&tmpbike3);err == nil{
                        fmt.Println("mget :",tmpbike1,tmpbike2,tmpbike3)
                }
        }
/////////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("mset","bike1","mongo","bike2","redis","bike3","memcached")
	if err != nil{
		fmt.Println("mset error",err.Error())
	}
        array,err =redis.Values( c.Do("MGET","bike1","bike2","bike3"))
        if err != nil{
                fmt.Println("mget failed",err.Error())
        }else{
                if _,err = redis.Scan(array,&tmpbike1,&tmpbike2,&tmpbike3);err == nil{
                        fmt.Println("mget :",tmpbike1,tmpbike2,tmpbike3)
                }
        }
//////////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("APPEND","bike1","xiaomingbike")
	if err != nil{
		fmt.Println("append error",err.Error())
	}
	bike1,err := redis.String(c.Do("GET","bike1"))
	if err == nil{
		fmt.Println("get bike is",bike1)
	}
//////////////////////////////////////////////////////////////////////////////////////
	str,err := redis.String(c.Do("substr","bike1",1,4))
	if err != nil{
		fmt.Println("substr get err",err.Error())
	}
	fmt.Println("substr is",str)
}

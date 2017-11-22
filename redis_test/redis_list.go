package main

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*************************************************************************************
rpush(key, value)：在名称为key的list尾添加一个值为value的元素
lpush(key, value)：在名称为key的list头添加一个值为value的 元素
llen(key)：返回名称为key的list的长度
lrange(key, start, end)：返回名称为key的list中start至end之间的元素
ltrim(key, start, end)：截取名称为key的list
lindex(key, index)：返回名称为key的list中index位置的元素
lset(key, index, value)：给名称为key的list中index位置的元素赋值
lrem(key, count, value)：删除count个key的list中值为value的元素
lpop(key)：返回并删除名称为key的list中的首元素
rpop(key)：返回并删除名称为key的list中的尾元素
blpop(key1, key2,… key N, timeout)：lpop命令的block版本。
brpop(key1, key2,… key N, timeout)：rpop的block版本。
rpoplpush(srckey, dstkey)：返回并删除名称为srckey的list的尾元素，并将该元素添加到名称为dstkey的list的头部
**************************************************************************************/
func main(){
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("conn to redis error",err.Error())
		return
	}	
	
	defer c.Close()
/////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("lpush","mylist","ofo","mobike","foo")
	if err != nil	{
		fmt.Println("redis lpush failed",err.Error())
	}
////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("rpush","mylist","bluegogo","xiaolan","xiaoming")
	if err != nil{
		fmt.Println("redis rpush failed",err.Error())
	}
///////////////////////////////////////////////////////////////////////////////
	num,err := c.Do("llen","mylist")
	if err != nil{
		fmt.Println("mylist get len err",err.Error())
	}else{
		fmt.Println("mylist's len is ",num)
	}
//////////////////////////////////////////////////////////////////////////////
	values,err := redis.Values(c.Do("lrange","mylist",0,10))
	if err != nil{
		fmt.Println("lrange err",err.Error())
	}
	fmt.Printf("mylist is:")
	for _,v := range values{
		fmt.Printf(" %s ",v.([]byte))
	}
	fmt.Printf("\n")
/////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("ltrim","mylist",0,4)
	if err != nil{
		fmt.Println("ltrim error",err.Error())
	}else{
		values,err = redis.Values(c.Do("lrange","mylist",0,4))
		if err != nil{
			fmt.Println("ltrim failed:",err.Error())
		}
		fmt.Printf("ltrim mylist is:")
		for _,v := range values{
			fmt.Printf("%s ",v.([]byte))
		}
		fmt.Printf("\n")
	}

//////////////////////////////////////////////////////////////////////////////
	val,err := c.Do("lindex","mylist",2)
	if err != nil{
		fmt.Println("lindex error:", err.Error())
	}else{
		fmt.Printf("lindex get result :%s\n",val)
	}
////////////////////////////////////////////////////////////////////////////////
	 _,err = c.Do("rpoplpush","mylist","mybike")
        if err != nil{
                fmt.Println("rpoplpush failed:",err.Error())
        }else{
                 values,err = redis.Values(c.Do("lrange","mylist",0,10))
                if err != nil{
                        fmt.Println("lrange failed:",err.Error())
                }
                for _,v := range values{
                        fmt.Printf("rpoplpush %s\n",v.([]byte))
                }


                values,err = redis.Values(c.Do("lrange","mybike",0,10))
                if err != nil{
                        fmt.Println("lrange failed:",err.Error())
                }
                for _,v := range values{
                        //fmt.Println(string(v.([]byte)))
                        fmt.Printf("rpoplpush %s\n",v.([]byte))
                }
        }

////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("lset","mylist",2,"mysql")
	if err != nil{
		fmt.Println("lset error:",err.Error())
	}
	val,err = c.Do("lindex","mylist",2)
        if err != nil{
                fmt.Println("lset error:", err.Error())
        }else{
                fmt.Printf("lset get result:%s\n",val)
        }
//////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("lrem","mylist",1,"mysql")
	if err != nil{
		fmt.Println("lrem error",err.Error())
	}else{
		values,err = redis.Values(c.Do("lrange","mylist",0,10))
                if err != nil{
                        fmt.Println("ltrim failed:",err.Error())
                }
                for _,v := range values{
                        fmt.Printf("lrem mylist: %s",v.([]byte))
                }
		fmt.Printf("\n")
	}
//////////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("lpop","mylist")
	if err != nil{
		fmt.Println("lpop failed:",err.Error())
	}else{
		values,err = redis.Values(c.Do("lrange","mylist",0,10))
                if err != nil{
                        fmt.Println("lpop failed:",err.Error())
                }
		fmt.Printf("lpop mylist :")
                for _,v := range values{
                        fmt.Printf("lpop mylist %s" ,v.([]byte))
                }
		fmt.Printf("\n")
	}
///////////////////////////////////////////////////////////////////////////////////////
	_,err = c.Do("rpop","mylist")
	if err != nil{
		fmt.Println("rpop failed",err.Error())
	}else{
		values,err = redis.Values(c.Do("lrange","mylist",0,10))
                if err != nil{
                        fmt.Println("rpop failed:",err.Error())
                }
		fmt.Printf("rpop mylist :")
                for _,v := range values{
                        //fmt.Println(string(v.([]byte)))
			 fmt.Printf("lpop mylist %s" ,v.([]byte))
                }
		fmt.Printf("\n")
	}
/////////////////////////////////////////////////////////////////////////////////////////
	res,err := c.Do("blpop","mylist",10)
	if err != nil{
		fmt.Println("blpop error")
	}else{
		fmt.Printf("blpop from mylist get:%s\n",res)
	}

	res,err = c.Do("blpop","tmpbike",10)
	if err != nil{
		fmt.Println("blpop time out")
	}else{
		fmt.Println("blpop from tmpbike get:",res)
	}
//////////////////////////////////////////////////////////////////////////////////////////
	res,err = c.Do("brpop","tmpbike",10)
        if err != nil{
                fmt.Println("brpop error")
        }else{
                fmt.Printf("brpop from tmpbike get :%s\n",res)
        }

        res,err = c.Do("brpop","mybike",10)
        if err != nil{
                fmt.Println("brpop time out")
        }else{
                fmt.Printf("brpop from mybike get:%s ",res)

        }
//////////////////////////////////////////////////////////////////////////////////////////
}

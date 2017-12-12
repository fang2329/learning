package main

import (
	"github.com/aarzilli/golua/lua"
	"fmt"
)


func main() {
	L := lua.NewState()
	L.OpenLibs()
	defer L.Close()

	err := L.DoFile("./test.lua")
	if err != nil{
		fmt.Println("err",err.Error())
	}

	//empty params empty results
	L.GetField(lua.LUA_GLOBALSINDEX,"GetStr")
	//L.PushInteger(1)
	L.Call(0,0)


	//two params empty results
	L.GetField(lua.LUA_GLOBALSINDEX,"GetBigger")
	L.PushInteger(6)
	L.PushInteger(3)
	L.Call(2,0)

	//empty params one result
	L.GetField(lua.LUA_GLOBALSINDEX,"GetResult")
	L.Call(0,1)
	ret1 := L.ToString(1)
	fmt.Println(ret1)
	L.Pop(1)	

	//two params one result
	L.GetField(lua.LUA_GLOBALSINDEX,"Compare")
	L.PushInteger(7)
	L.PushInteger(9)
	L.Call(2,1)
	ret2 := L.ToInteger(1)
	fmt.Println(ret2)
	L.Pop(1)

	L.GetField(lua.LUA_GLOBALSINDEX,"MoreReturn")	
	L.PushInteger(10)
	L.Call(1,3)
	ret3 := L.ToString(1)
	ret4 := L.ToString(2)
	ret5 := L.ToString(3)
	fmt.Println(ret3,ret4,ret5)

	
}

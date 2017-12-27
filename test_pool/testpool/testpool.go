package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"test1226/test_pool/pool"
	"time"
)

const(
	maxGroutine = 5
	poolRes = 2
)

func main()  {
	var wg sync.WaitGroup
	wg.Add(maxGroutine)
	
	p,err := pool.NewPool(createConnection,poolRes)
	if err != nil{
		log.Println(err)
		return
	}
	
	for query:=0;query < maxGroutine;query++{
		go func(q int) {
			dbQuery(q,p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("begin close")
	p.Close()
}

func dbQuery(query int,pooll *pool.Pool)  {
	conn,err := pooll.Acquire()
	if err != nil{
		log.Println(err)
		return
	}
	defer pooll.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
	log.Println("id[%d],con[%d]",query,conn.(*dbConnection).ID)
}

type dbConnection struct{
	ID int32
}

func (db *dbConnection)Close()error  {
	log.Println("close connection",db.ID)
	return nil
}

var idCounter int32

func createConnection()(io.Closer,error)  {
	id := atomic.AddInt32(&idCounter,1)
	return &dbConnection{id},nil
}
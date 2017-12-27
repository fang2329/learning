package main

import (
	"log"
	"os"
	"test1226/test_runner/runner"
	"time"
)

func main()  {
	log.Println("begin")
	timeout := 4*time.Second
	r := runner.NewRunner(timeout)
	r.AddTask(creatTask(),creatTask(),creatTask(),creatTask())
	if err := r.Start();err != nil{
		switch err{
		case runner.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		case runner.ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		}
	}
	log.Println("over")
}

func creatTask()func(int)  {
	return func(id int) {
		log.Println("doing %d",id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}

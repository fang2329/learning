package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrInterrupt = errors.New("interrupt")
	ErrTimeOut = errors.New("time out")
)

type Runner struct {
	tasks []func(int)
	complete chan error
	timeout <-chan time.Time
	interrupt chan os.Signal
}

func NewRunner(tm time.Duration) *Runner  {
	return &Runner{
		complete:make(chan error),
		timeout:time.After(tm),
		interrupt:make(chan os.Signal,1),
	}
}

func (this *Runner)AddTask(task ...func(int))  {
	this.tasks = append(this.tasks,task...)
}

func (this *Runner)run()error  {
	for id,task := range this.tasks{
		if this.IsInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (this *Runner)IsInterrupt()bool {
	select {
	case <-this.interrupt:
		signal.Stop(this.interrupt)
		return true
	default:
		return false
	}
}

func (this *Runner)Start()error  {
	signal.Notify(this.interrupt,os.Interrupt)
	go func() {
		this.complete<- this.run()
	}()
	
	select{
	case err := <- this.complete:
		return err
	case <-this.timeout:
		return ErrTimeOut
	}
}



package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct{
	m sync.Mutex
	res chan io.Closer
	factory func()(io.Closer,error)
	closed bool
}

var(
	ErrorPoolClosed = errors.New("pool is closed")
	ErrorPollSize  = errors.New("size is too small")
)

func NewPool(fn func()(io.Closer,error),size int)(*Pool,error)  {
	if size <= 0{
		return nil,ErrorPollSize
	}
	return &Pool{
		factory:fn,
		res : make(chan io.Closer,size),
	},nil
}

func (p *Pool)Acquire()(io.Closer,error)  {
	select{
	case r,ok:=<-p.res:
		log.Println("shared resouce")
		if !ok{
			return nil,ErrorPoolClosed
		}
		return r,nil
	default:
		log.Println("new resource")
		return p.factory()
	}
}

func (p *Pool)Close()  {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed{
		return
	}
	p.closed = true
	close(p.res)
	for r:=range p.res{
		r.Close()
	}
}

func (p *Pool)Release(r io.Closer)  {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed{
		r.Close()
		return
	}
	select {
	case p.res <-r:
		log.Println("released")
	default:
		log.Println("pool is too full")
		r.Close()
	}
}
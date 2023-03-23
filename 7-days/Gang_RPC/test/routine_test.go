package test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func dosome() int64 {
	select {}
	var result int64
	return result
}

func Test_r1(t *testing.T) {
	ch := make(chan int64)
	done := make(chan struct{})
	defer close(done)
	go func() {
		res := dosome()
		select {
		case <-done:
			log.Printf("defer done")
			close(ch)
			return
		case ch <- res:
		default:
			log.Printf("default")
		}
	}()
	select {
	case <-time.After(time.Second * 3):
		log.Printf("timeout")
		return
	case res := <-ch:
		fmt.Println(res)
	}
}

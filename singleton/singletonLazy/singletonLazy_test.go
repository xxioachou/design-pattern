package singletonlazy

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	t1 := GetInstance()
	t2 := GetInstance()
	if t1 != t2 {
		t.Fatal("instance not equal")
	}
}

func TestSingletonParallelly(t *testing.T) {
	const parCount = 100
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*singletonLazy {}
	for i := 0; i < parCount; i ++ {
		go func(index int) {
			// 协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = GetInstance()
		} (i)
		wg.Done()
	}
	
	// 关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i ++ {
		if instances[i] != instances[0] {
			msg := fmt.Sprintf("instance not equal, id is %v", i)
			t.Fatal(msg)
		}
	}
}
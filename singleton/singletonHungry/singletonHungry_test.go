package singletonhungry

import (
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
	instances := [parCount]*singletonhungry {}
	for i := 0; i < parCount; i ++ {
		go func(index int) {
			// 阻塞
			<-start
			instances[index] = GetInstance()
		} (i)
		wg.Done()
	}
	
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i ++ {
		if instances[i] != instances[0] {
			t.Fatal("instance not equal.")
		}
	}
}
package singletonlazy

import (
	"fmt"
	"sync"
)

// 通过接口限制访问，隐藏实现细节：

type ISingleton interface {
	Hello()
}

// 结构体首字母小写，是包私有的
type singletonLazy struct {
	name string
}

func (obj *singletonLazy) Hello() {
	msg := fmt.Sprintf("singletonLazy: hello, %v!", obj.name)
	fmt.Println(msg)
}

var (
	once sync.Once
	instance *singletonLazy
)

// 只能通过这个方法获取结构体的实例
func GetInstance() *singletonLazy {
	once.Do(func ()  {
		instance = &singletonLazy{"lazy man"}
	})
	return instance
}
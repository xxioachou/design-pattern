package singletonhungry

import (
	"fmt"
)

// 通过接口限制访问，隐藏实现细节：

type ISingleton interface {
	Hello()
}

// 结构体首字母小写，是包私有的
type singletonhungry struct {
	name string
}

func (obj *singletonhungry) Hello() {
	msg := fmt.Sprintf("singletonhungry: hello, %v!", obj.name)
	fmt.Println(msg)
}

var instance = &singletonhungry{"hungry man"}

// 只能通过这个方法获取结构体的实例
func GetInstance() *singletonhungry {
	return instance
}
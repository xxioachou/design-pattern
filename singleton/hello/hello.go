package main

import (
	"singleton/singletonLazy"
	"singleton/singletonHungry"
)

func main() {
	singletonlazy.GetInstance().Hello()
	singletonhungry.GetInstance().Hello()
}
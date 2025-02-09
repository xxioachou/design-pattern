module singleton/hello

go 1.23.3

replace singleton/singletonLazy => ../singletonLazy

require (
	singleton/singletonHungry v0.0.0-00010101000000-000000000000
	singleton/singletonLazy v0.0.0-00010101000000-000000000000
)

replace singleton/singletonHungry => ../singletonHungry

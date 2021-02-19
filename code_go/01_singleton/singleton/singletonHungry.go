package singleton

var singletonHungry *singleton

func init() {
	singletonHungry = &singleton{}
}

// GetInstance 获取实例
func GetInstanceHungry() *singleton {
	return singletonHungry
}
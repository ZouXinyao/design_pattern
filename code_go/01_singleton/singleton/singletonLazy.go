package singleton

import "sync"

var instanceLazy *singleton

// GetInstanceLazy public获取实例
func GetInstanceLazy() *singleton {
	if instanceLazy == nil {
		instanceLazy = &singleton{}
	}

	return instanceLazy
}

var instanceLock *singleton
var mu0 sync.Mutex

func GetInstanceLock() *singleton {
	mu0.Lock()
	defer mu0.Unlock()

	if instanceLock == nil {
		instanceLock = &singleton{}
	}

	return instanceLock
}

var instanceTwoCheck *singleton
var mu1 sync.Mutex

func GetInstanceTwoCheck() *singleton {
	if instanceTwoCheck == nil {
		mu1.Lock()
		defer mu1.Unlock()

		if instanceTwoCheck == nil {
			instanceTwoCheck = &singleton{}
		}
	}

	return instanceTwoCheck
}

var instanceOnce *singleton
var once sync.Once

func GetInstanceOnce() *singleton {
	once.Do(func() {
		instanceOnce = &singleton{}
	})
	return instanceOnce
}
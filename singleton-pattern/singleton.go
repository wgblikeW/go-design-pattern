package singleton

import "sync"

// Singleton Pattern

type singleton struct{}

// var ins *singleton = &singleton{}
var ins *singleton
var mu sync.Mutex
var once sync.Once

// GetIns return the instance of singleton
func GetIns() *singleton {
	if ins == nil {
		mu.Lock()
		if ins == nil {
			ins = &singleton{}
		}
		mu.Unlock()
	}
	return ins
}

func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

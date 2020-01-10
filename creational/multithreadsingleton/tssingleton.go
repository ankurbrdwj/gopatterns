package multithreadsingleton

import (
	"fmt"
	"sync"
)

type tssingleton struct {
	data int
	desc string
}

var initialized uint32
var instance *tssingleton
var mu sync.Mutex
var once sync.Once

func (s *tssingleton) GetData() int {
	return s.data
}

func (s *tssingleton) GetDesc() string {
	return s.desc
}

func (s *tssingleton) ToString() {
	fmt.Println(s.desc, s.desc)
}

//GetInstance factory method
func GetInstance() *tssingleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			return new(tssingleton)
		}
	}
	return instance
}

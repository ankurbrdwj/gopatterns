package dooncesingleton

import (
	"fmt"
	"sync"
)

type dosingleton struct {
	data int
	desc string
}

var instance *dosingleton
var once sync.Once

func (s *dosingleton) GetData() int {
	return s.data
}

func (s *dosingleton) GetDesc() string {
	return s.desc
}

func (s *dosingleton) ToString() {
	fmt.Println(s.desc, s.desc)
}

//GetInstance factory method
func GetInstance() *dosingleton {
	once.Do(func() {
		instance = new(dosingleton)
	})
	return instance
}

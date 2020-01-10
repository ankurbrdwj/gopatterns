package singleton

import (
	"fmt"
)

type singleton struct {
	data int
	desc string
}

var instance *singleton

func (s *singleton) GetData() int {
	return s.data
}

func (s *singleton) GetDesc() string {
	return s.desc
}

func (s *singleton) ToString() {
	fmt.Println(s.desc, s.desc)
}

//GetInstance factory method
func GetInstance() *singleton {
	if instance == nil {
		return new(singleton)
	}
	return instance
}

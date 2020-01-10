package singleton_test

import (
	"github.com/ankurbrdwj/gopatterns/creational/dooncesingleton"
	"github.com/ankurbrdwj/gopatterns/creational/multithreadsingleton"
	"github.com/ankurbrdwj/gopatterns/creational/singleton"
	"testing"
)

//simple singleton
func TestGetInstance1(t *testing.T) {
	counter1 := singleton.GetInstance()
	if counter1 == nil {
		//Test 1 failed
		t.Error("expected pointer to singleton struct , found  nil")
	}
	expectedCounter := counter1
	counter2 := singleton.GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Singleton instances must be different")
	}

}

//multi thread singleton
func TestGetInstance2(t *testing.T) {
	counter1 := multithreadsingleton.GetInstance()
	if counter1 == nil {
		//Test 1 failed
		t.Error("expected pointer to singleton struct , found  nil")
	}
	expectedCounter := counter1
	counter2 := multithreadsingleton.GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Singleton instances must be different")
	}

}

// do once singleton
func TestGetInstance3(t *testing.T) {
	counter1 := dooncesingleton.GetInstance()
	if counter1 == nil {
		//Test 1 failed
		t.Error("expected pointer to singleton struct , found  nil")
	}
	expectedCounter := counter1
	counter2 := dooncesingleton.GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Singleton instances must be different")
	}

}

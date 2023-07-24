package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func GetInstance() *single {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

var once sync.Once

type single2 struct {
}

var singleInstance2 *single2

func GetInstance2() *single2 {
	if singleInstance2 == nil {
		once.Do(
			func() {
				fmt.Println("Creating single2 instance now.")
				singleInstance2 = &single2{}
			})
	} else {
		fmt.Println("Single2 instance already created.")
	}

	return singleInstance2
}

func Clear() {
	singleInstance = nil
}

func Clear2() {
	singleInstance2 = nil
	once = sync.Once{}
}

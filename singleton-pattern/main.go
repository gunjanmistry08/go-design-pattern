package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance1() *single {
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

func getInstance2() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single Instance already exist")
	}
	return singleInstance
}

func main() {

	singleInstance = nil
	for i := 0; i < 30; i++ {
		go getInstance2()
	}
	// because no wait group so this i guess
	fmt.Scanln()
}

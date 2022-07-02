package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// using wait group variable. wg is a package scope
var wg sync.WaitGroup

// keep in mind we see one gorutine running at all times. this is the main function which is running in the first gorutine always
func main(){
	// accessingBaiscSysInformation()
	// runningFunctionIntoGorutine()
	//raceConditionExample()
	fixRaceConditionWithMutex()
}

//  runtime package allows us to check couple of things regarding our os that will be usefull for concurrency
func accessingBaiscSysInformation(){

	fmt.Printf("OS: %v, ARCH: %v, CPUs: %v, GORUTINES: %v \t", runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.NumGoroutine())
}

func foo() string{
	for i := 0; i <10; i++ {
		fmt.Printf("#Foo run interation %v \n", i)
	}

	wg.Done() // stating that wait group is done

	return "this will be discarded" // any value returned by function will be DISCARDED. It is good idea to wrap a function in a anonymus function, and us channel to get the value from it
}

func bar(){
	for i := 0; i <10; i++ {
		fmt.Printf("#bar run interation %v \n", i)
	}
}

func runningFunctionIntoGorutine(){
	// to run function in concurrency mode we just add go in front of it. 
	// Important thing to notice is that the code is not going to run in paraller if we have only one CPU!
	go foo() // this will go into itrs own gorutine. We wont see anything printed as after spinning this goruttine rest of the main() will execute and shut down the program.

	 bar()

	// In order to wait for the code we need build in primitives used for this i.e mutex (mutual exclusion lock) or wait group



	wg.Add(1) //wait for one thing
	go foo()

	wg.Wait() // waits untill all things added are done
}

// method set of other type T consists of all methods declared with reciver type T

// in general concurrency has issue due to multiple threads (gorutines?) sharing same variables and causing race conditions. Golang pushes channels as the way to share variables that are used by diffrent threads. Only one gorutine has a access to a value at any given time

var wg1 sync.WaitGroup

func raceConditionExample(){
	fmt.Printf("wg is: %v \n", wg1)
	counter := 0

	const goRutine = 100;
	wg1.Add((goRutine)) // wait for 100 gorutines we will start to finish

	for i := 0; i < goRutine; i++ { // start 100 go rutines
		go func(){
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg1.Done()
		}()

		fmt.Printf("counter %v, go ruttines: %v \n", counter, runtime.NumGoroutine())
	}
	wg1.Wait()
	fmt.Printf("counter %v \n", counter)
}


func fixRaceConditionWithMutex(){

	counter := 0

	const goRutine = 100;
	wg1.Add((goRutine))

	// we lock variable by mutex, so only one gorutine can use it at a time
	var mutex sync.Mutex

	for i := 0; i < goRutine; i++ {
		go func(){
			mutex.Lock() // lock mutating variables

			v := counter
			runtime.Gosched()
			v++
			counter = v

			mutex.Unlock() // unlock mutating variables
			wg1.Done()
		}()

	}
	wg1.Wait()
	fmt.Printf("counter %v \n", counter)
}

func fixingRaceConditionWithAtomic(){
	var counter int64

	const goRutine = 100;
	wg1.Add((goRutine))

	for i := 0; i < goRutine; i++ { 
		go func(){
			atomic.AddInt64(&counter, 1) // another way to mutate values and avoid race conidition is using atomic package
			runtime.Gosched()
			wg1.Done()
		}()

	}
	wg1.Wait()
	fmt.Printf("counter %v \n", counter)
}
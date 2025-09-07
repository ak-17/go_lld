package using_mutexes

import (
	"fmt"
	"sync"
)

func printNumbers(start int, end int, increment int, name string, thisMutex *sync.Mutex, otherMutex *sync.Mutex, group *sync.WaitGroup) {
	defer group.Done()
	for i := start; i <= end; i += increment {
		thisMutex.Lock()
		fmt.Printf("%s , number: %d\n", name, i)
		otherMutex.Unlock()
	}
}

func PrintEvenOdd(end int) {
	var oddMutex sync.Mutex
	var evenMutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)
	// Locking Odd Mutex to force Event thread to start first
	oddMutex.Lock()
	go printNumbers(0, end, 2, "even", &evenMutex, &oddMutex, &wg)
	go printNumbers(1, end, 2, "odd", &oddMutex, &evenMutex, &wg)
	wg.Wait()
	// Unlocking even Mutex at last, the last printed number mutex to be unlocked
	// if end is odd, then odd mutex will be locked last
	if end%2 == 0 {
		evenMutex.Unlock()
	} else {
		oddMutex.Unlock()
	}
}

// Usage
//
//func main() {
//	odd_even.PrintEvenOdd(100)
//}

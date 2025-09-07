package using_channels

import (
	"fmt"
	"sync"
)

func printNumbers(start, end int, increment int, wg *sync.WaitGroup, thisCh chan bool, otherCh chan bool, name string) {
	defer wg.Done()
	for i := start; i <= end; i += increment {
		<-thisCh
		fmt.Printf("%s , number: %d\n", name, i)
		if i+1 <= end {
			otherCh <- true
		}
	}
}

func PrintEvenOdd(end int) {
	evenCh := make(chan bool, 1)
	oddCh := make(chan bool, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	evenCh <- true

	go printNumbers(0, end, 2, &wg, evenCh, oddCh, "even")
	go printNumbers(1, end, 2, &wg, oddCh, evenCh, "odd")

	wg.Wait()

	close(evenCh)
	close(oddCh)

}

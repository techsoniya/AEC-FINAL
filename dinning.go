package main

import (
	"fmt"
	"sync"
)

type Chopstick struct{
	sync.Mutex 
}

func main() {
	var num int
	fmt.Print("enter number:")
	fmt.Scan(& num)

	var wg sync.WaitGroup
	chopsticks := make([]Chopstick, num)

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			left := id
			right := (id + 1) % num

			chopsticks[left].Lock()
			chopsticks[right].Lock()

			fmt.Printf("Philosopher %d is eating\n", id)

			chopsticks[right].Unlock()
			chopsticks[left].Unlock()

			fmt.Printf("Philosopher %d is thinking\n", id)
		}(i)
	}
	wg.Wait()
}
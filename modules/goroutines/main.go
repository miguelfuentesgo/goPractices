package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func RunGoroutines() {

	var items = []Item{
		{Id: "1", Name: "Cook meal"},
		{Id: "2", Name: "Do my tasks"},
		{Id: "3", Name: "Find a new job"},
	}

	var WaitGroup sync.WaitGroup

	for _, item := range items {
		fmt.Println("For started...")
		WaitGroup.Add(1)
		go process(item, &WaitGroup)

	}

	WaitGroup.Wait()
	fmt.Println("Program finished")
}

type Item struct {
	Id   string
	Name string
}

func process(item Item, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Processing item: ", item.Id)
	time.Sleep(1 * time.Second)
	fmt.Println("Procces finished: ", item.Id)

}

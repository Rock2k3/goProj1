package TestApp2

import (
	"app2/repository"
	"fmt"
	"sync"
)

func TestApp2() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		repository.DbMigration()
		wg.Done()
	}()

	fmt.Println("Test App2")

	wg.Wait()
}

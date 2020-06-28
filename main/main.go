package main

import (
	"fmt"
	"sync"

	client_one "gitlab.com/Patrones_de_diseno/client_one"
	client_two "gitlab.com/Patrones_de_diseno/client_two"

	singleton "gitlab.com/Patrones_de_diseno/singleton"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()

		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}
	wg.Wait()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}

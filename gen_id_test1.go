package main

import (
	"fmt"
	"sync"
)

func GenTest1() {
	node := Init()

	fmt.Println("Start")
	fmt.Println(node.Generate())
	fmt.Println(node.Generate())
	fmt.Println(node.Generate())
}

func GenTest2() {
	node := Init()

	fmt.Println("Start")
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println(node.Generate())
	}()
	go func() {
		defer wg.Done()
		fmt.Println(node.Generate())
	}()
	go func() {
		defer wg.Done()
		fmt.Println(node.Generate())
	}()

	wg.Wait()
}

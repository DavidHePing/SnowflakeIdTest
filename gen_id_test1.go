package main

import (
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
)

func GenTest1(node *snowflake.Node) {

	fmt.Println("Start")
	fmt.Println(node.Generate())
	fmt.Println(node.Generate())
	fmt.Println(node.Generate())
}

func GenTest2(node *snowflake.Node) {

	fmt.Println("Start")
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		// fmt.Println(node.Generate())
		fmt.Println(node.Generate().Base2())
	}()
	go func() {
		defer wg.Done()
		// fmt.Println(node.Generate())
		fmt.Println(node.Generate().Base2())
	}()
	go func() {
		defer wg.Done()
		// fmt.Println(node.Generate())
		fmt.Println(node.Generate().Base2())
	}()

	wg.Wait()
}

package main

import "time"

func main() {
	// Define a specific date
	specificDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	node1 := Init2(time.Now().UnixMilli())
	node2 := Init2(specificDate.UnixMilli())

	// GenTest1(node)
	GenTest2(node1)
	GenTest2(node2)

	time.Sleep(3000)

	GenTest2(node1)
	GenTest2(node2)
}

package main

import (
	"container/list"
	"fmt"
)

func main() {
	values := list.New()
	values.PushFront("Prateek")
	values.PushBack("1")
	values.PushFront("Raj")
	values.PushFront("Rajasthan")
	values.PushFront("Raju")

	for i := values.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value, i.Next())
	}
}

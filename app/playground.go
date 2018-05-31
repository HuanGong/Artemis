package main

import (
	"time"
	"math/rand"
	"fmt"
)

type (
	TData struct {
		abc int64
	}
)
var (
	Data map[int64]*TData
)

func GenNewMapData() {
	tricks := time.NewTicker(time.Second * 1)
	for _ = range tricks.C {
		newData := make(map[int64]*TData)
		i := 10000000
		for i > 0 {
			i--
			k := rand.Int63n(0xAFFFFFFFFFFFFFF)
			newData[k] = &TData{
				abc: k,
			}
		}
		Data = newData
	}
}

func ReaderMapData() {
	var i int64 = 0
	for {
		for k, v := range Data {
			if k != v.abc {
				fmt.Println("seems some data was wrong")
			}
			i++
			if i % 10000 == 0 {
				fmt.Println("access 1w item")
			}
		}
		time.Sleep(time.Millisecond * 3)
	}
}

func main() {
	go GenNewMapData()

	go ReaderMapData()
	go ReaderMapData()
	ReaderMapData()
}

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	numberToGenerate = 1000000
)

func main() {
	start := time.Now()
	fmt.Println("Start")

	list := lottoNumbers(numberToGenerate)

	d := time.Now().Sub(start)
	fmt.Println("End ", len(list), " ", d.Seconds())
}

func lottoNumbers(n int) [][]int {
	list := make([][]int, n)
	seed := time.Now().UnixNano()
	var wg sync.WaitGroup
	workers := runtime.GOMAXPROCS(-1) // one for each proc

	for i := 0; i < workers; i++ {
		work := n / workers
		begin := work * i
		end := begin + work

		if i == workers-1 {
			end += n % workers
		}

		r := rand.New(rand.NewSource(seed + int64(i)))

		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := begin; i < end; i++ {
				list[i] = []int{
					r.Intn(49),
					r.Intn(49),
					r.Intn(49),
					r.Intn(49),
					r.Intn(49),
					r.Intn(49),
					r.Intn(49),
				}
			}
		}()
	}

	wg.Wait()

	return list
}

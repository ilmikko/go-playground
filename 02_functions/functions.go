package main

import (
	"fmt"
	"math/rand"
	"time"
)

func average(values ...float32) float32 {
	if len(values) > 0 {
		var avg float32
		for _, val := range values { // discard the index key with _
			avg += val
		}
		return avg / float32(len(values))
	}
	return 0.0
}

func getSlice(N int, max float32) []float32 {
	// generate float32 random values
	var slice []float32
	seed := rand.NewSource(time.Now().UnixNano()) // seed from current nanosecs
	random := rand.New(seed)                      // random object from this seed
	for i := 0; i < N; i++ {
		a := random.Float32()*max + 1                // get range [1-max] instead of [0.0-1.0]
		slice = append(slice, a-(a-float32(int(a)))) // discard decimal place
	}
	return slice
}

// takes in a float32 slice, returns the maximum value (= divide & conq)
// fewer comparaisons compared to generic linear

func main() {
	fmt.Println(average(34.9, 9, 67.0, 4, 23.1), average())
	data := []float32{1, 2, 3}
	fmt.Println(average(data...)) // unpacking equivlent to *args in python

	//pre declare a function variable beforehand if recursion will be used
	var max func(slice []float32, low int, max int) float32
	
	max = func(slice []float32, low int, high int) float32 {
		if low == high || low+1 == high {
			if slice[low] >= slice[high] {
				return slice[low]
			}
			return slice[high]
		} else {
			lowMax := max(slice, low, (low+high)>>1)
			highMax := max(slice, (low+high)>>1, high)

			if lowMax > highMax {
				return lowMax
			}
			return highMax
		}
	}
	slice := getSlice(8,7)
	fmt.Println(slice, ", average =", average(slice...), ", maximum=", max(slice, 0, len(slice)-1))
}

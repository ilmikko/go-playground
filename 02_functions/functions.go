package main

import (
	"fmt"
	"math/rand"
	"time"
)

func average(values ... float32) float32 {
	if len(values) > 0 {
		var avg float32
		for _, val := range values { // discard the index key with _
			avg += val
		}
		return avg/float32(len(values))
	}
	return 0.0
}

func main(){
	fmt.Println(average(34.9,9,67.0,4,23.1), average())
	data := []float32{1,2,3}
	fmt.Println(average(data...)) // unpacking equivlent to *args in python
	
	// generate float32 random values
	var slice []float32
	seed := rand.NewSource(time.Now().UnixNano()) // seed from current nanosecs
	random := rand.New(seed)// random object from this seed
	for i := 0 ; i < 7 ; i++ {
		a := random.Float32() * 7.0 + 1// get range [1-7] instead of [0.0-1.0]
		slice = append(slice,a - (a - float32(int(a)))) // discard decimal place
	}
	fmt.Println(slice,"'s average =" ,average(slice...))
}

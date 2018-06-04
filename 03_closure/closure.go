package main

import (
	"bytes"
	"fmt"
	"time"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"github.com/tati-z/go-playground/00_formatting/utility"
)

type number struct {
	value    int
	isPrime  bool
	isPower2 bool
}

// receivers in both toString and toBytes are passed by value
func (num number) toString() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("%d is ", num.value))
	if num.isPrime || num.isPower2 {
		buffer.WriteString("prime or a power of two")
	} else {
		buffer.WriteString("neither prime nor a power of two")
	}
	return buffer.String()
}

func (num number) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d,%t,%t\n", num.value, num.isPrime, num.isPower2))
	return buffer.Bytes()
}

func (num *number) newNumber(theNumber int){
	// pointer notation *num is not used because THANKFULLY golang
	// automatically deferences pointers.
	num.value = theNumber
	val := math.Log2(float64(num.value)) 
	// can be computer archicture dependant
	num.isPower2 =  val - (float64(int(val))) == 0
	num.isPrime = utility.IsPrime(num.value)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//wrapper receives one file and return a function for writting to this file
func csvGen(filename string)  (func()*number,func(value number)) {
	// create a file of the given filename
	dir, err := os.Getwd()
	checkError(err)
	file, err := ioutil.TempFile(dir, filename)
	defer file.Close()
	
	checkError(err)
	// return the inner funtion with
	writer := func(value number){
		os.OpenFile(file.Name(), os.O_CREATE|os.O_APPEND, 0600)
		_, err := file.Write(value.toBytes())
		checkError(err)
	}
	reader := func() *number{
		//read the number object from file return a number object 
	}
	return reader, writer
}
func main(){
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var num number
	num.newNumber(random.Intn(23))
	fmt.Println(num.toString())

	read, write := csvGen("test1")
	fmt.Printf("%T\n", write)
	write(num)
}


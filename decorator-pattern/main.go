package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"
)

type piFunc func(int) float64

// calculate PI using Leibiniz formula
func calculatePI(n int) float64 {
	cn := make(chan float64)

	for i := 0; i < n; i++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(cn, float64(i))
	}
	result := 0.0

	for i := 0; i < n; i++ {
		result += <-cn
	}
	return result
}

//A cache using sync Map
func syncCache(fun piFunc, c *sync.Map) piFunc {
	return func(n int) float64 {
		fn := func(n int) float64 {
			key := fmt.Sprintf("n=%d", n)
			val, ok := c.Load(key)
			if ok {
				return val.(float64)
			}
			result := fun(n)
			c.Store(key, result)
			return result
		}
		return fn(n)
	}
}

//Logger function to log time and calculate how long it took a function to run
func timeLogger(fun piFunc, l *log.Logger) piFunc {
	return func(n int) float64 {
		fn := func(n int) (result float64) {
			defer func(t time.Time) {
				l.Printf("Time took=%v, iteration=%v, Value if PI=%v", time.Since(t), n, result)
			}(time.Now())

			return fun(n)
		}
		return fn(n)
	}
}

func main() {
	// limit 1000000
	// fmt.Println(calculatePI(100))
	s := sync.Map{}
	g := syncCache(calculatePI, &s)
	f := timeLogger(g, log.New(os.Stdout, "[DEBUG]", 1))
	f(100000)
	f(10000)
	f(1000)
	f(100)
	fmt.Println(s)

}

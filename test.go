package main

import "fmt"

func f() {
	const c = 1.0
	var a float32 = 3
	var X = 6.14*2 + c + a
	var Q, W = 1, 2
	q, w, e := 3, 2, 1
	d := q + w + e + Q + W

	fmt.Printf("%v %v\n", X, d)
}

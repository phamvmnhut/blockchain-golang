package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	x := [10]byte{10, 11, 12}
	fmt.Println(x[9:10])
}
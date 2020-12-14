package main

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Seq-%d: %s\n", i, RandomString())
	}
}

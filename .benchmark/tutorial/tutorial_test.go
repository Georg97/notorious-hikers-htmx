package main

import (
	"testing"
)

// Questions
// what does b.N do?

func BenchmarkSlice(b *testing.B) {
    n := 100

    b.Run("slice non alloc", func(b *testing.B) {
        // outside of bench
        for i := 0; i < b.N; i++ {
            // bench code
            // fmt.Printf("Hello, World")
            ints := []int{}
            for i := 0; i < n; i++ {
                ints = append(ints, i)
            }
        }
    })

    b.Run("slice alloc", func(b *testing.B) {
        // outside of bench
        for i := 0; i < b.N; i++ {
            // bench code
            // fmt.Printf("Hello, World")
            ints := make([]int, n)
            for i := 0; i < n; i++ {
                ints[i] = i
            }
        }
    })
}


package main

import (
	"bytes"
	"testing"
)

// Questions
// what does b.N do? -> It is created by the go benchmark utility and runs the code until go thinks its ok now (until differences are not so big anymore)

func writeBuffer(msg []byte) {
    buf := new(bytes.Buffer)
    buf.Write(msg)
}

func BenchmarkBuffer(b *testing.B) {
    n := 100

    b.Run("rawdog bytes-array", func(b *testing.B) {
        msg := []byte("Hello, World")
        for i := 0; i < b.N; i++ {
            for i := 0; i < n; i++ {
                writeBuffer(msg)
            }
        }
    })
}


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


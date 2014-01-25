// run with
//     $ go run hamming.go

package main

/*
#include "hamming.c"
*/
import "C"

import "fmt"
import "math/rand"
import "time"

func popcount(input uint64) uint8 {

	var mask1, mask2, mask3, mask4, mask5, mask6 uint64
	mask1 = 6148914691236517205 // 01010101...
	mask2 = 3689348814741910323 // 00110011...
	mask3 = 1085102592571150095 // 00001111...
	mask4 = 71777214294589695   // 8 zeroes, 8 ones, etc...
	mask5 = 70367670468607      // 16 zeroes, 16 ones, etc...
	mask6 = 4294967295          // 32 zeroes, 32 ones

	input = (input & mask1) + ((input >> 1) & mask1)
	input = (input & mask2) + ((input >> 2) & mask2)
	input = (input & mask3) + ((input >> 4) & mask3)
	input = (input & mask4) + ((input >> 8) & mask4)
	input = (input & mask5) + ((input >> 16) & mask5)
	input = (input & mask6) + ((input >> 32) & mask6)

	return uint8(input)
}

func main() {

	var sampleSize int = 2000000
	randoms := make([]uint64, sampleSize)
	for i := 0; i < sampleSize; i++ {
		randoms[i] = uint64(rand.Int63())
	}

	// using native code:
	var weight uint64
	start := time.Now()
	for _, v := range randoms {
		weight += uint64(popcount(v))
	}
	elapsed := time.Since(start)
	fmt.Printf("Native Go took %s, popcount: %v\n", elapsed, weight)

	// calling raw C:
	weight = 0
	start = time.Now()
	for _, v := range randoms {
		weight += uint64(C.weight(C.ulonglong(v)))
	}
	elapsed = time.Since(start)
	fmt.Printf("Calling C function from Go (casting every time) took %s, popcount: %v\n", elapsed, weight)
}

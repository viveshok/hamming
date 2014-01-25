// run with
//     $ go run hamming.go

package main

import "fmt"
import "math/rand"
import "time"

func popcount(input int64) uint8 {

	var mask1, mask2, mask3, mask4, mask5, mask6 int64
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
	randoms := make([]int64, sampleSize)
	for i := 0; i < sampleSize; i++ {
		randoms[i] = rand.Int63()
	}

	var weight uint64
	start := time.Now()
	for i := 0; i < sampleSize; i++ {
		weight += uint64(popcount(randoms[i]))
	}
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}

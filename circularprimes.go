package main

import (
	"math"
	"os"
	"strconv"
)

func rotate(i float64) float64 {
	return math.Mod(i, 10)*(math.Pow(10, math.Floor(math.Log10(i)))) + math.Floor(i/10)
}

func sieve(num int64, nums []bool) {
	for i := num * 2; i < int64(len(nums)); i += num {
		nums[i] = false
	}
}

func main() {
	num, _ := strconv.ParseFloat(os.Args[1], 10)
	nums := make([]bool, int64(num))
	for i := range nums {
		nums[i] = true
	}
	for i := 2; i <= int(math.Sqrt(num)); i++ {
		sieve(int64(i), nums)
	}
	circularprimes := 0
	for i := range nums {
		if i > 2 && nums[i] {
			candidate := float64(i)
			ok := true
			desc := strconv.FormatInt(int64(i), 10)
			for j := 0; j < int(math.Floor(math.Log10(float64(i)))); j++ {
				candidate = rotate(candidate)
				if (candidate < num) && !nums[int(candidate)] {
					// this rotation is not prime
					ok = false
					break
				} else {
					// this rotation is prime, so mark the prime as visited
					nums[int(candidate)] = false
					desc += ", "
					desc += strconv.FormatInt(int64(candidate), 10)
				}
			}

			if ok {
				println(desc)
				circularprimes++
			}
		}
	}
	println("found", circularprimes, "primes")
}

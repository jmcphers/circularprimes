package main

import (
	"math"
	"os"
	"strconv"
)

// given a number, "rotates" the number by one by moving the last digit into
// the position of the first digit: e.g. 337 -> 733
func rotate(i float64) float64 {
	return math.Mod(i, 10)*(math.Pow(10, math.Floor(math.Log10(i)))) + math.Floor(i/10)
}

// performs the "sieve of Erostothanes" for a single number
func sieve(num int64, primes []bool) {
	for i := num * 2; i < int64(len(primes)); i += num {
		primes[i] = false
	}
}

func main() {
	// validate arguments
	if len(os.Args) < 2 {
		println("usage: circularprimes.go NUM")
		return
	}

	// make an array and mark all numbers "prime"
	num, _ := strconv.ParseFloat(os.Args[1], 10)
	primes := make([]bool, int64(num))
	for i := range primes {
		primes[i] = true
	}

	// mark all non-prime numbers
	for i := 2; i <= int(math.Sqrt(num)); i++ {
		sieve(int64(i), primes)
	}

	// find primes and check to see if all their rotations are also prime
	circularprimes := 0
	for i := range primes {
		if i > 1 && primes[i] {
			// this prime might be circular
			candidate := float64(i)
			ok := true
			desc := strconv.FormatInt(int64(i), 10)

			// check all its rotations (1 rotation for 2-digit numbers,
			// 2 rotations for 3-digit numbers, etc.)
			var rotations = int(math.Floor(math.Log10(float64(i))))
			var primerotations = 0
			for j := 0; j < rotations; j++ {
				candidate = rotate(candidate)
				if (candidate < num) && !primes[int(candidate)] {
					// this rotation is not prime
					ok = false
					break
				} else if candidate < num && candidate != float64(i) {
					// this rotation is prime, so mark the prime as visited
					primes[int(candidate)] = false
					desc += ", " + strconv.FormatInt(int64(candidate), 10)
					primerotations++
				}
			}

			// if all rotation were prime, output the number
			if ok {
				println(desc)

				// count the number and all its rotations as circular primes
				circularprimes += (1 + primerotations)
			}
		}
	}
	println("found", circularprimes, "primes")
}

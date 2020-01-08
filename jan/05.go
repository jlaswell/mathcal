package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/randall77/factorlib"
	"github.com/randall77/factorlib/big"
)

func main() {
	// Get the prime factorization of all 2 digit numbers.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := log.New(os.Stdout, "", 0)
	allFactors := generatePrimeFactors(r, l)

	// @link https://www.math.upenn.edu/~deturck/m170/wk2/numdivisors.html
	// In general, if you have the prime factorization of the number n,
	// then to calculate how many divisors it has, you take all the exponents
	// in the factorization, add 1 to each, and then multiply these
	// "exponents + 1"s together.

	divisors := map[int]int{}
	for i, factors := range allFactors {
		divisors[i+10] = calculateDivisors(factors)
	}

	for k, v := range divisors {
		if v == 12 {
			log.Printf("%d has 12 divisors", k)
		}
	}
}

func generatePrimeFactors(r *rand.Rand, l *log.Logger) [90][]big.Int {
	allFactors := [90][]big.Int{}
	for i := 10; i < 100; i++ {
		factors, err := factorlib.Factor(big.Int64(int64(i)), "trial", r, l)
		if err != nil {
			log.Println(err)
		}
		allFactors[i-10] = factors
	}

	return allFactors
}

func calculateDivisors(factors []big.Int) int {
	divisorCount := map[int64]int{}
	for _, f := range factors {
		divisorCount[f.Int64()]++
	}
	// Start with 1 to ease multiplication of divisors.
	divisors := 1
	for _, d := range divisorCount {
		divisors *= (d + 1)
	}

	return divisors
}

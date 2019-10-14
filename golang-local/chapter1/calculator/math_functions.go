package main
import (
	"fmt"
	"math"
	)

func main(){
	//x :=[]float64{ 98, 33, 106, 84, 7	}
	val := fact(5)
	fmt.Println(val)

	x := math.Mod(6,3)
	fmt.Println(x)

	fmt.Printd(isPrime(3))
}

func isPrime(x int) bool {
	bool_prime := true

	cap := math.Pow(float64(x),0.5)
	for i:=2.0; i<= cap; i++{
		if math.Mod(float64(x),i) == 0 {
			bool_prime = false
			break
		}
	}
	return bool_prime
}


func countUpTo(x int) int {
	var yield int
	if x == 0{
		yield = 0
	} else {
		yield = x + countUpTo(x-1)
	}
	return yield
}

func fact(x int) int {
	var yield int
	if x <= 0 {
		yield = 1
	} else {
		yield = x * countUpTo(x-1)
	}
	return yield
}

func sum(xs [] float64) float64 {
	var total float64 = 0
	for _, value := range xs{
		total+= value
	}
	return total
}

func prod(xs []float64) float64{
	var total float64 = 1
	for _, value := range xs{
		total *= value
	}
	return total
}

func artih_mean(xs []float64) float64{
	return sum(xs)/float64(len(xs))
}

func harmonic_mean(xs []float64) float64{
	return 	math.Pow( prod(xs), 1/float64(len(xs)))
}
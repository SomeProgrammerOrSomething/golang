package main
import (
	"fmt"
	"math"
)

func main(){
	x :=[]float64{ 98, 33, 106, 84, 7	}


	val := sum(x)
	fmt.Println(val)
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
package main
import "fmt"

func isTriviallySorted( xs []float64) bool {
	possibly_sorted := true
	if len(xs) != 0 && len(xs) != 1 {
		possibly_sorted = false
		//fmt.Println("DEBUG - isTriviallySorted triggered False.\nThe array has length greater than 1")
	}
	return possibly_sorted
}

func isDescending (xs []float64) bool {
	possibly_sorted := true

	for i:=1; i< len(xs) ; i++ {
		if xs[i-1] >= xs[i]{
			//fmt.Printf("DEBUG - isDescending triggered True for check %d\n",i)
		} else {
			//fmt.Printf("DEBUG - isDescending triggered FALSE for check %d\n",i)
			//fmt.Println("The code hit the else statement.")
			possibly_sorted = false
			break
		}
	}
	return possibly_sorted
}
func isAscending(xs []float64) bool {
	possibly_sorted := true

	for i:=1; i< len(xs) ; i++ {
		if xs[i-1] <= xs[i]{
			//fmt.Printf("DEBUG - isASCENDING triggered True for check %d\n",i)
		} else {
			//fmt.Printf("DEBUG - isASCENDING triggered FALSE for check %d\n",i)
			//fmt.Println("The code hit the else statement.")
			possibly_sorted = false
			break
		}
	}
	return possibly_sorted
}

func isSorted (xs []float64) bool {
	return ( isTriviallySorted(xs) || isDescending(xs) || isAscending (xs) )
}

func main(){
	xs := []float64{1,2,3,4,8983,98399}
	fmt.Printf("The statement 'The array %v is sorted' is a %t statement\n",xs,isSorted(xs))
}
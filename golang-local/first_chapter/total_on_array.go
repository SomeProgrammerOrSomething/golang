package main
import "fmt"

func main(){
  var x [5]float64

  x[0] = 324
  x[1] = 3.43
  x[2] = 999
  x[3] = 100*34
  x[4] = 86.58

  var total float64 = 0
  for i := 0; i < 5; i++ {
    total += x[i]
  }
  fmt.Println("The total sum of all elements in the array is:",total)
}
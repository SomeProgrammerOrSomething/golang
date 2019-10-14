package main
import "fmt"

func main(){
  var x [5]float64

  x[0] = 324
  x[1] = 3.43
  x[2] = 999
  x[3] = 100*34
  x[4] = 86.58

  var total,average float64 = 0,0
  for i := 0; i < 5; i++ {
    total += x[i]
  }

  average = total/float64(len(x))
  fmt.Println("The average value of the array is:",average)
}
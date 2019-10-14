package main
import "fmt"

func main(){
  var x [5] int

  for i := 0; i< len(x);i++ {
    fmt.Println("Please enter an integer: ")
    fmt.Scanf("%d",&x[i])
  }

  var min, max int = 0,0
  for i := 0; i<len(x);i++{
    if x[i] < min {
      min = x[i]
    }

    if x[i] > max {
      max = x[i]
    }
  }

  fmt.Printf("The max value is: %d\nThe min value is: %d\n",max,min)
}
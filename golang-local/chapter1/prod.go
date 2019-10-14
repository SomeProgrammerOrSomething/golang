package main
import "fmt"

func main(){
  var x [5] int

  for i := 0; i< len(x);i++ {
    fmt.Println("Please enter an integer: ")
    fmt.Scanf("%d",&x[i])
  }

  var prod int = 1
  for i := 0; i<len(x);i++{
  prod *= x[i]
  }

  fmt.Println("The product of all elements multiplied together is:",prod)
}
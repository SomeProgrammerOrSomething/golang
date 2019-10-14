package main
import "fmt"

func main() {
  fmt.Print("Enter a number: ")
  var count int
  fmt.Scanf("%d%", &count)

  i := 1
  for i <= count {
    fmt.Println(i)
    i += 1
  }

  fmt.Println("Lift Off!")
}
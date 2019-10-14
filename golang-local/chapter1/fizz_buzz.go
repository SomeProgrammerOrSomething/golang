package main
import "fmt"

func main(){
	var count int

	fmt.Println("Hello! Let's Play Fizz Buzz!")
	fmt.Println("Pick a positive number: ")

	fmt.Scanf("%d", &count)

	fmt.Println("\n")
	i := 1
	for i <= count {
		if i%15 == 0{
			fmt.Println("Fizz Buzz!")
		} else if i%3 == 0{
			fmt.Println("Fizz!")
		} else if i%5 == 0{
			fmt.Println("Buzz!")
		} else {
			fmt.Println(i)
		}
		i+=1
	}
	fmt.Println("Thanks for playing!")
}

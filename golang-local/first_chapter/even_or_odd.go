package main
import "fmt"

func main(){
	var num int
	fmt.Println("Pick an integer: ")
	fmt.Scanf("%d",num)

	if num%2 == 0{
		fmt.Println("Your number is even!")
	} else {
		fmt.Println("Your number is odd!")
	}
}
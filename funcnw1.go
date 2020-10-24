package main 

import "fmt"

func main() {
    superman()
    result := multiplyme(3,6)
	fmt.Println(result)
	myesult, mylength, myname := adder(3, 6, 4, 2, 4, 12, 23)
	fmt.Println(myresult, mylenth, myname)

}

func superman() {
    fmt.Println("I am a superman")
}

func multiplyme(v1, v2 int) int {
    return v1 * v2
    
}

func adder)values ...int) (int, int, string) {
	sum := 0
	for i:= range values {
		sum = sum + values[i]
	}
	length := len(values)
	name := "just for fun"

	return sum, length, name
}


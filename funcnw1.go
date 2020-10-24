package main 

import "fmt"

func main() {
    superman()
    result := multiplyme(3,6)
    fmt.Println(result)

}

func superman() {
    fmt.Println("I am a superman")
}

func multiplyme(v1, v2 int) int {
    return v1 * v2
    
}
 package main 
  
import "fmt"
  
 
func main() { 
      
   
    mmap := map[int]string{ 
        24:"Kloudone", 
        34:"KloudoneOfficial", 
        55:"KludoneLearning", 
    } 
    for key, value:= range mmap { 
       fmt.Println(key, value)  
    } 
    
}
package main
 
import (
  "os"
  "fmt" 
  "strings" 
)
 
func main() {   
    // Set custom env variable
    os.Setenv("CUSTOM", "500")
     
    // fetch all env variables
    for _, element := range os.Environ() {
        variable := strings.Split(element, "=")
        fmt.Println(variable[0],"=>",variable[1])        
    }
}
package main  
import "fmt"  
func main(){  
   fmt.Print("Enter text: "); 
   var input int  
   fmt.Scanln(&input); 
   if (input < 0 || input > 100) {  
      fmt.Printf("Please enter valid no");  
   } else if (input >= 0 && input < 50  ) {  
      fmt.Printf(" Fail");  
   } else if (input >= 50 && input < 60) {  
      fmt.Printf(" D Grade");  
   } else if (input >= 60 && input < 70  ) {  
      fmt.Printf(" C Grade");  
   } else if (input >= 70 && input < 80) {  
      fmt.Printf(" B Grade");  
   } else if (input >= 80 && input < 90  ) {  
      fmt.Printf(" A Grade");  
   } else if (input >= 90 && input <= 100) {  
      fmt.Printf(" A+ Grade");  
   }  
}  

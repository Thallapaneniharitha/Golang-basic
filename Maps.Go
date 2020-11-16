1.program:
package main
import "fmt"
func main() {
   var countryCapitalMap map[string]string
  
   countryCapitalMap = make(map[string]string)
   countryCapitalMap["rajampet"] = "Paris"
   countryCapitalMap["kadapa"] = "Rome"
   countryCapitalMap["Tirupathi"] = "Tokyo"
   countryCapitalMap["Rayachoti"] = "New Delhi"
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   capital, ok := countryCapitalMap["United States"]
   
   if(ok){
      fmt.Println("Capital of United States is", capital)  
   } else {
      fmt.Println("Capital of United States is not present") 
   }
}
output:
Capital of Rayachoti is New Delhi
Capital of rajampet is Paris
Capital of kadapa is Rome
Capital of Tirupathi is Tokyo
Capital of United States is not present

2.program:
package main
import "fmt"
func main() {   
  
   countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
   
   fmt.Println("Original map")   
   
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   } 
   
   delete(countryCapitalMap,"France");
   fmt.Println("Entry for France is deleted")   
   fmt.Println("Updated map")   
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
}
output:
Original map
Capital of Japan is Tokyo
Capital of India is New Delhi
Capital of France is Paris
Capital of Italy is Rome
Entry for France is deleted
Updated map
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of India is New Delhi

3.program:
package main  
import "fmt"  
func main() {   
    m_a_p := map[int]string{ 
  
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "parrort", 
    } 
  
    for id, pet := range m_a_p { 
  
        fmt.Println(id, pet) 
    } 
} 
output:
94 parrort
90 Dog
91 Cat
92 Cow
93 Bird

4.program:
package main  
import "fmt" 
func main() { 
  
    m_a_p := map[int]string{ 
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "parrort", 
    } 
  
    fmt.Println("Original map: ", m_a_p) 
  
    m_a_p[95] = "Parrot"
    m_a_p[96] = "Crow"
    fmt.Println("Map after adding new key-value pair:\n", m_a_p) 
  
    
    m_a_p[91] = "PIG"
    m_a_p[93] = "DONKEY"
    fmt.Println("\nMap after updating values of the map:\n", m_a_p) 
} 
output:
Original map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:parrort]
Map after adding new key-value pair:
map[90:Dog 91:Cat 92:Cow 93:Bird 94:parrort 95:Parrot 96:Crow]

Map after updating values of the map:
map[90:Dog 91:PIG 92:Cow 93:DONKEY 94:parrort 95:Parrot 96:Crow]

5.program:
package main  
import "fmt" 
func main() {    
    m_a_p := map[int]string{ 
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
    fmt.Println("Original map: ", m_a_p) 
    value_1 := m_a_p[90] 
    value_2 := m_a_p[93] 
    fmt.Println("Value of key[90]: ", value_1) 
    fmt.Println("Value of key[93]: ", value_2) 
} 
output:
Original map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]
Value of key[90]:  Dog
Value of key[93]:  Bird

6.program:
package main  
import "fmt"
func main() {    
    m_a_p := map[int]string{ 
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
  
    fmt.Println("Original map: ", m_a_p) 
    pet_name, ok := m_a_p[90] 
    fmt.Println("\nKey present or not:", ok) 
    fmt.Println("Value:", pet_name)   
    _, ok1 := m_a_p[92] 
    fmt.Println("\nKey present or not:", ok1) 
} 
output:
Original map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]
Key present or not: true
Value: Dog
Key present or not: true

7.program:
package main  
import "fmt"
func main() {     
    m_a_p := map[int]string{ 
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
  
    fmt.Println("Original map: ", m_a_p) 
    delete(m_a_p, 90) 
    delete(m_a_p, 93) 
    fmt.Println("Map after deletion: ", m_a_p) 
} 
output:
Original map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]
Map after deletion:  map[91:Cat 92:Cow 94:Rabbit]

8.program:
package main   
import "fmt" 
// Main function 
func main() { 
  
    // Creating and initializing a map 
    m_a_p := map[int]string{ 
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
    fmt.Println("Original map: ", m_a_p) 
  
    // Assigned the map into a new variable 
    new_map := m_a_p 
  
    // Perform modification in new_map 
    new_map[96] = "Parrot"
    new_map[98] = "Pig"
  
    // Display after modification 
    fmt.Println("New map: ", new_map) 
    fmt.Println("\nModification done in old map:\n", m_a_p) 
} 
output:
Original map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]
New map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit 96:Parrot 98:Pig]

Modification done in old map:
map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit 96:Parrot 98:Pig]

9.program:
package main 
import "fmt" 
func main() { 
  
    // Creating and initializing empty map 
    // Using var keyword 
    var map_1 map[int]int
  
    // Checking if the map is nil or not 
    if map_1 == nil { 
      
        fmt.Println("True") 
    } else { 
      
        fmt.Println("False") 
    } 
  
    // Creating and initializing a map 
    // Using shorthand declaration and 
    // using map literals 
    map_2 := map[int]string{ 
      
            90: "Dog", 
            91: "Cat", 
            92: "Cow", 
            93: "Bird", 
            94: "Rabbit", 
    } 
      
    fmt.Println("Map-2: ", map_2) 
} 
output:
True
Map-2:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]

10.program:
package main  
import "fmt" 
func main() { 
  
    // Creating a map 
    // Using make() function 
    var My_map = make(map[float64]string) 
    fmt.Println(My_map) 
  
    // As we already know that make() function 
    // always returns a map which is initialized 
    // So, we can add values in it 
    My_map[1.3] = "Rohit"
    My_map[1.5] = "Sumit"
    fmt.Println(My_map) 
} 
output:
map[]
map[1.3:Rohit 1.5:Sumit]

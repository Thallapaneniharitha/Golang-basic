1.program:
package main
import "fmt"
func main() {
   var n [10]int 
   var i,j int
      
   for i = 0; i < 9; i++ {
      n[i] = i + 90 
   }
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}
output:
Element[0] = 90
Element[1] = 91
Element[2] = 92
Element[3] = 93
Element[4] = 94
Element[5] = 95
Element[6] = 96
Element[7] = 97
Element[8] = 98
Element[9] = 0

2.program:
package main  
import "fmt" 
func main() { 
var myarr[3]string
myarr[0] = "HIK"
myarr[1] = "Haritha"
myarr[2] = "SARI"
fmt.Println("Elements of Array:") 
fmt.Println("Element 1: ", myarr[0]) 
fmt.Println("Element 2: ", myarr[1]) 
fmt.Println("Element 3: ", myarr[2]) 
} 
output:
Elements of Array:
Element 1:  HIK
Element 2:  Haritha
Element 3:  SARI

3.program:
package main  
import "fmt" 
func main() { 
arr:= [4]string{"hari", "hik", "sarad123", "prasanthi"} 
fmt.Println("Elements of the array:")  
for i:= 0; i < 3; i++{ 
fmt.Println(arr[i]) 
}   
}
output:
Elements of the array:
hari
hik
sarad123

4.program:
package main  
import "fmt" 
func main() { 
arr:= [3][3]string{{"C#", "C", "Python"},  
                   {"Java", "Css", "AWS"}, 
                    {"C++", "Go", "HTML"},} 
fmt.Println("Elements of Array 1") 
for x:= 0; x < 3; x++{ 
for y:= 0; y < 3; y++{ 
fmt.Println(arr[x][y]) 
} 
} 
var arr1 [2][2] int
arr1[0][0] = 100 
arr1[0][1] = 200 
arr1[1][0] = 300 
arr1[1][1] = 400 
fmt.Println("Elements of array 2") 
for p:= 0; p<2; p++{ 
for q:= 0; q<2; q++{ 
fmt.Println(arr1[p][q])  
} 
} 
} 
output:
Elements of Array 1
C#
C
Python
Java
Css
AWS
C++
Go
HTML
Elements of array 2
100
200
300
400

5.program:
package main   
import "fmt" 
func main() 
{ 
var myarr[5]int 
fmt.Println("Elements of the Array :", myarr) 
  
} 
output:
Elements of the Array : [0 0 0 0 0]

6.program:
package main  
import "fmt"
func main()
 {      
arr1:= [3]int{9,7,6} 
arr2:= [...]int{9,7,6,4,5,3,2,4} 
arr3:= [3]int{9,3,5}  
fmt.Println("Length of the array 1 is:", len(arr1)) 
fmt.Println("Length of the array 2 is:", len(arr2)) 
fmt.Println("Length of the array 3 is:", len(arr3)) 
} 
output:
Length of the array 1 is: 3
Length of the array 2 is: 8
Length of the array 3 is: 3

7.program:
package main  
import "fmt"  
func main() 
{     
myarray:= [...]string{"HIKR", "GFG", "saritha", 
                    "haritha", "HIKR"}  
fmt.Println("Elements of the array: ", myarray) 
fmt.Println("Length of the array is:", len(myarray)) 
} 
output:
Elements of the array:  [HIKR GFG saritha haritha HIKR]
Length of the array is: 5

8.program:
package main 
  
import "fmt"
 
func main() {     
myarray:= [...]int{29, 45, 49, 39, 
                   10, 49, 48, 60}  
for x:=0; x < len(myarray); x++{ 
fmt.Printf("%d\n", myarray[x]) 
} 
} 
output:
29
45
49
39
10
49
48
60

9.program:
package main 
  
import "fmt"
  
func main() { 
      

my_array:= [...]int{100, 200, 300, 400, 500} 
fmt.Println("Original array(Before):", my_array) 
  

new_array := my_array 
  
fmt.Println("New array(before):", new_array) 
  

new_array[0] = 500 
  
fmt.Println("New array(After):", new_array) 
  
fmt.Println("Original array(After):", my_array) 
} 
output:
Original array(Before): [100 200 300 400 500]
New array(before): [100 200 300 400 500]
New array(After): [500 200 300 400 500]
Original array(After): [100 200 300 400 500]

10.program:
package main 
  
import "fmt"
  
func main()
 {   
arr1:= [3]int{9,7,6} 
arr2:= [...]int{9,7,6} 
arr3:= [3]int{9,5,3}  
fmt.Println(arr1==arr2) 
fmt.Println(arr2==arr3) 
fmt.Println(arr1==arr3)  
arr4:= [4]int{9,7,6} 
fmt.Println(arr1==arr4) 

} 
output:
true
false
false

11.program:
package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}   
	var y [5]int = [5]int{10, 20, 30} 

	fmt.Println(x)
	fmt.Println(y)
}
output:
[10 20 30 40 50]
[10 20 30 0 0]

12.program:
package main

import "fmt"

func main() {
	var theArray [3]string
	theArray[0] = "India"  
	theArray[1] = "Rajampet"
	theArray[2] = "Tirupathi"  
	fmt.Println(theArray[0]) 
	fmt.Println(theArray[1]) 
	fmt.Println(theArray[2]) 
}
output:
India
Rajampet
Tirupathi

13.program:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := [...]int{10, 20, 30}

	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(len(x))
}
output:
array
3


 

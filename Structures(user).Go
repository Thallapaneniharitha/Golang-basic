1.program //Accessing Structure Members
package main
import "fmt"
type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books   
   var Book2 Books     
   Book1.title = "java Programming"
   Book1.author = "Haritha"
   Book1.subject = "python Programming Tutorial"
   Book1.book_id = 84942305
   Book2.title = "HTML"
   Book2.author = "Zara Ali"
   Book2.subject = "Css"
   Book2.book_id = 456785
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)  
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}
output:
Book 1 title : java Programming
Book 1 author : Haritha
Book 1 subject : python Programming Tutorial
Book 1 book_id : 84942305
Book 2 title : HTML
Book 2 author : Zara Ali
Book 2 subject : Css
Book 2 book_id : 4567893

2.program:    //Structures as Function Arguments
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    
   var Book2 Books    
 
   
   Book1.title = "java Programming"
   Book1.author = "Haritha"
   Book1.subject = "python Programming Tutorial"
   Book1.book_id = 6495407

  
   Book2.title = "HTML Billing"
   Book2.author = "saritha"
   Book2.subject = "CSS Billing Tutorial"
   Book2.book_id = 6495700
 
   printBook(Book1)
   
   printBook(Book2)
}
func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}
output:
Book title : java Programming
Book author : Haritha
Book subject : python Programming Tutorial
Book book_id : 6495407
Book title : HTML Billing
Book author : saritha
Book subject : CSS Billing Tutorial
Book book_id : 6495700

3.program:
package main
import "fmt"
type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books   
   var Book2 Books   
   Book1.title = "java Programming"
   Book1.author = "Haritha"
   Book1.subject = "python Programming Tutorial"
   Book1.book_id = 6495406   
   Book2.title = "HTML Billing"
   Book2.author = "rani"
   Book2.subject = "rani Billing Tutorial"
   Book2.book_id = 6495700
   printBook(&Book1)
   printBook(&Book2)
}
func printBook( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}
output:
Book title : java Programming
Book author : Haritha
Book subject : python Programming Tutorial
Book book_id : 6495406
Book title : HTML Billing
Book author : rani
Book subject : rani Billing Tutorial
Book book_id : 6495700

4.package main  
import (  
   "fmt"  
)  
type person struct {  
   firstName string  
   lastName  string  
   age       int  
}  
func main() {  
   x := person{age: 25, firstName: "THALLAPANENI", lastName: "HARITHA", }  
   fmt.Println(x)  
   fmt.Println(x.firstName)  
} 
$go run main.go
{THALLAPANENI HARITHA 25}
HARITHA

5.package main 
  
import "fmt"
  
type Car struct { 
    Name, Model, Color string 
    WeightInKg         float64 
} 
  
func main() { 
    c := Car{Name: "Ferrari", Model: "GTC4", 
            Color: "pink", WeightInKg: 1920} 
    fmt.Println("Car Name: ", c.Name) 
    fmt.Println("Car Color: ", c.Color) 
    c.Color = "white"
    fmt.Println("Car: ", c) 
} 
$go run main.go
Car Name:  Ferrari
Car Color:  pink
Car:  {Ferrari GTC4 white 1920}

6.package main 
  
import "fmt"
type Employee struct { 
    firstName, lastName string 
    age, salary int
} 
  
func main() { 
  emp8 := &Employee{"haritha", "Anderson", 25, 4500} 
    fmt.Println("First Name:", (*emp8).firstName) 
    fmt.Println("Age:", (*emp8).age) 
} 
$go run main.go
First Name: haritha
Age: 25

7.package main 
  
import "fmt"
type Employee struct { 
    firstName, lastName string 
    age, salary         int
} 
  
func main() { 
   emp8 := &Employee{"ravi", "Anderson", 50, 4000} 
    fmt.Println("First Name: ", emp8.firstName) 
    fmt.Println("Age: ", emp8.age) 
} 
$go run main.go
First Name:  ravi
Age:  50

8.package main

import "fmt"

type User struct {
    name       string
    occupation string
    age        int
}

func main() {

    u := User{"haritha", "gardener", 25}

    fmt.Printf("%s is %d years old and she is a %s\n", u.name, u.age, u.occupation)
}
$go run main.go
haritha is 25 years old and she is a gardener

9.package main

import "fmt"

type Address struct {
    city    string
    country string
}

type User struct {
    name    string
    age     int
    address Address
}

func main() {

    p := User{
        name: "haritha",
        age:  25,
        address: Address{
            city:    "Delhi",
            country: "USA",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("Country:", p.address.country)
}
$go run main.go
Name: haritha
Age: 25
City:Delhi
Country: USA

10.package main  
import (  
   "fmt"  
)  
type person struct {  
   fname string  
   lname string}  
type employee struct {  
   person  
   empId int  
}  
func (p person) details() {  
   fmt.Println(p, " "+" I am a person")  
}  
func (e employee) details() {  
   fmt.Println(e, " "+"I am a employee")  
}  
func main() {  
   p1 := person{"HARITHA", "CHOWDARY"}  
   p1.details()  
   e1 := employee{person:person{"SARITHA", "THALLAPANENI"}, empId: 11}  
   e1.details()  
}  
$go run main.go
{HARITHA CHOWDARY}   I am a person
{{SARITHA THALLAPANENI} 11}  I am a employee

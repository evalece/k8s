package main
import "fmt"


/*

Golang work dir example 
-go.mod 
-main.go // entry point
-internal //only accessible within this module
		- service
			= service.go
			- interface.go
- pkg //reusabke packages to be imported; source code based
- cmd // public CLi 
- Build // compiled binary structs, functions, interfaces and lib
- models
	-entities.go // basic structs
- tests 
	- service_test.go //test interface 


*/


// variable

var x int =10 

// if else 

if x> 5 {
	fmt.Println("x is greater than 5")
} else if x== 5 {
	fmt.Println ("x is 5")
} else {
	fmt.Println("x is less than 5")
}		

// for loops
for i:=0; i<10; i++ {
	fmt.Println(i)
}

//while loop:
for x < 10 {
	x++
	fmt.Println(x)
}

//switch 
switch x{
case 1:
	fmt.Println("case 1")
case 2:	fmt.Println("case 2")
default: 	fmt.Println("default")

}



//fucntions
func add(a int, b int) int {
	return a+b
}

func division_mod(a, b, int) (int, int) {
	return a/b, a%b
}


//array

arr:= []int {1,2,3
arr= append(arr,4)


// map
hash_map := make(map [string] int)
m["book"]= 3
val : m["book"] // =3 

// OOP
OOP {




	/*
	OOP Basics Ref, note, potential mistake on "
	func (c *CompanyEntity) getType() CompanyEntity {
    return *c // or some specific value
	}" 
	: https://dev.to/parthlaw/object-oriented-go-unraveling-the-power-of-oop-in-golang-49h6

	Class: A class is a blueprint of an object to be created.
	Object: Instance of class containing both data and methods.
	Encapsulation: Limiting access to properties, methods or variables in a class for code external to that class.
	Inheritance: It allows a class (commonly referred to as subclass) to inherit properties and methods from another class (commonly referred to as parent class).
	Polymorphism: Allowing objects of different classes to be treated as objects of a common superclass.
	Abstraction: Abstracting a basic concept of multiple classes in a single abstract class. This simplifies logic and makes code more readable.
		
		*/

	/*
	Struct - contains methods and proterties 
	Recall 

	type CompanyEntity struct {
    Id string
    Name string
    Country string
    manager string
	}

	func (c *CompanyEntity) getType() CompanyEntity {
    return *c // or some specific value
	}


	Go Modules- packages to export and unexport identifiers 
		- i.e, modify struct data
		- set private for module access only
	Inheritance
		- imbed structs in child struct
		- access embedded functions 
		- create new struct functions for embedded struct
		- In the following case where Employee embeds Company, the Employee struct can access (assume not private) Company struct data and method, but the methods are applied to 
			Company data only, not on   Salary int; for example:
			type Employee struct{
				Company;
				Salary int;
				}
		

	Polymorphism (allowing different objects to be treated as the same parent classs)
		- 1. test Interface with struct in test_service.go
			var _ CompanyInterface = &e 
		- 2. calling the function, pkg main: 
			e := Employee{
				CompanyEntity: CompanyEntity{Name: "OpenAI"},
				Salary:        100000,
			}
 		getCompanyNameInLowerCase(&e)

	*/


	// structs and methods
		// structs- classes without methods and holds data only 
	type Person struct {
		Name string
		Age int
		Language string

	} 
		
		// instantiate of struct; no inheritance, use "composition instead", serve more like a protobuff to define data type
			//  composition vs extend ref: https://www.digitalocean.com/community/tutorials/composition-vs-inheritance
			// Inheritance- Maintains heirachy for super(); tight coupling in calling inhereted fucntions 
		p: Person{Name: "Evelece" Age: 18, Language: "Golang"}


		
		//Method- Method Signatures only like member functions-- Person is reciever (i.e., self in py )
		func (p Person) Speak() string{
			return "Person speaks "+ p.Language
		}

		//Interface -- like header files in cpp; method signatures
		type Speaker interface {

			Speak () string
		}
	
	// Override example 
		//- as long as filed in the desired pkg, recommended to bind with structs
	func (e *Employee) getType() string {
    return "Employee"
}


	

}


//Pointers and addresses

var x int =10
var p *int = &x // pointer p holds the address (&x; or we can treate as pointer holding x ) of int x; which has value of 10
*p =20 // dereference + assign; now x= 20
	// Go interface allows type chekcing by passing address/ pointers
	var myinterface Myinterface = &x // assume this works and var p *int= &x, note that Golang's interface does not take the later one


	// We must define our own interface, i.e., in interface.go, under the service name folder and at the same level as the service.go 
	package company

	type CompanyInterface interface {
		saveToDatabase();
		getType() CompanyEntity;
	}


// main function 

func main() {

	fmt.Println("Hello, World!")
}

// Conccurrency and Parallel programming 

/*
Idea from textbook - 
CPU cores - one core one task at a time
Concurrent- concurrent means the core and swtich from doing task x then y then x.
Parallel - Thus, if multi-core, tasks can be assign and performed at the same time; assume no conflicts or race or other concurrency issues.


In the contecxt of Golang; Ref: https://medium.com/@ksandeeptech07/understanding-concurrency-and-parallelism-in-go-golang-f53afaf0387b 

Multi-core+ Parallel in Golang
		- Goruntime with go scheduler to distribute tasks in multi-core processing 
		- 
	Channel - Coordinatitions between tasks 
	(I'm writing from a perspective of C/C++ and Java since this is where I come from, where threads, thread pools, worker nodes are controlled manually). 
	Goroutines 
		- Remove this need for thread pool management mannually
		- concurrent threads, managed and scheduled by Go runtime by M:N scheduling 
		- go doWork(id, task)
	channels - Shared data and inner thread pool comminication
		- 


*/
package main
import "fmt"

type Menber struct{
	name string
	age int
}

func main() {
	var hello string = "hello"
	langage_name := "go"
	var pi float32 = 3.14
	fmt.Println(hello + " " + langage_name)
	fmt.Printf("pi is %.2f\n", pi)
	
	numbers := [5]int8{11,22,33,44,55}
	fmt.Println(numbers[1])

	name := []string{"wrench"}
	name = append(name,"michelangelo")
	fmt.Println(name[1])
	var member1 Menber
	member1.name = "wrench"
	member1.age = 21
	fmt.Println(member1)
	fmt.Println(member1.name)
}

package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")

	lang := "Go"
	name := "Stephen"
	year := 2026

	fmt.Printf("My name is %s and I'm learning %s in %d\n", name, lang, year)

	str := fmt.Sprintf("%s is awesome!", lang)
	fmt.Println(str)

	one, two, three := "one", "two", "three"
	fmt.Println(fmt.Sprintf("%s %s %s", one, two, three))

	fmt.Println(Greet("Stephen"))

	// Create a formatted error
	e := "An expected error occured."
	err := fmt.Errorf("Error: %s", e)
	println(err.Error())

	fmt.Print("Enter a number: ")
	var num string
	_, err = fmt.Scanln(&num)
	if err != nil {
		fmt.Println(fmt.Sprintf("You entered %s", num))
	}
}

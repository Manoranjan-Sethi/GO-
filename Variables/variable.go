package main

import "fmt"

func main (){

	// STRINGS
	//if we declare a variable and don't use it then its an error
	var newName string = "Mano"

	var newName2 = "Ramesh"

	var newName3 string

	fmt.Println(newName, newName2, newName3)

	//updates

	newName2 = "Suresh"

	newName3 = "Ramesh"

	fmt.Println(newName, newName2, newName3)

	//short hand declaration

	newName4 := "GO lets GO"  //automatically infers the variable as string

	fmt.Println(newName4)


	// INTEGERS

	var age int = 22

	var age2 = 32

	age3 := 26

	fmt.Println(age, age2, age3)

}
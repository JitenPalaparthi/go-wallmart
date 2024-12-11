package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "Jitenp@outlook.com", Status: "active", Address: Address{Line1: "Medical College", City: "Trivandrum", Status: "Corporation"}}
	p1.Address.Status = "Capital City"

	p1.PrintPerson()
	p1.PrintAddress()

}

type Person struct {
	Name, Email, Status string
	Address             // promoted field
}

func (p *Person) PrintPerson() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Status:", p.Status)
}

type Address struct {
	Line1, City, Status string
}

func (a *Address) PrintAddress() {
	fmt.Println("Line1:", a.Line1)
	fmt.Println("City:", a.City)
	fmt.Println("Status:", a.Status)
}

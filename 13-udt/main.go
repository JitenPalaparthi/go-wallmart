package main

import "fmt"

func main() {
	cc1 := ColourCode{int: 101, integer: 1, string: "Red"}
	fmt.Println(cc1)

	p1 := Person{Name: "Jiten", Email: "Jitenp@outlook.com", Address: &Address{Line1: "Medical College", City: "Trivandrum"}}
	fmt.Println(p1.Address)
	fmt.Println(p1)

	s1 := Student{Name: "Jiten", Email: "Jitenp@outlook.com", Address: struct {
		Line1 string
		City  string
	}{Line1: "Medical College", City: "Trivandrum"}, SocialMedia: struct {
		Twitter  string
		Linkedin string
	}{Twitter: "jiten_p", Linkedin: "https://linkedin/in/jpalaparthi"}}

	fmt.Println(s1)

	var c1 struct{ Name, Email, WebSite string } = struct {
		Name    string
		Email   string
		WebSite string
	}{Name: "AKA-Labs", Email: "jp@aka-labs.in", WebSite: "https://aka-labs.in"}
	fmt.Println(c1)
	var c2 Company = Company{Name: "AKA-Labs", Email: "jp@aka-labs.in", WebSite: "https://aka-labs.in"}
	fmt.Println(c2)
}

type integer = int // this is not a new type just an alias

type ColourCode struct {
	int
	integer
	string
}

type Person struct {
	Name, Email string
	Address     *Address // field name can be the type name or can give different
}

type Address struct {
	Line1, City string
}

type Student struct {
	Name, Email string
	Address     struct {
		Line1, City string
	}
	SocialMedia struct {
		Twitter, Linkedin string
	}
}

type Company struct{ Name, Email, WebSite string }

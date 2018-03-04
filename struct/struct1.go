package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

func (p *Person) move(newAddr string) string {
	oldAddress := p.Address
	p.Address = newAddr
	return oldAddress
}

func main() {
	p := Person{"Rebert", "Male", 33, "Beijing"}

	oldAddress := p.move("San Francisco")

	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}

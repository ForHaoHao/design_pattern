package main

import "fmt"

type Print interface {
	Print(value string) error
}

type Fax interface {
	Send(value string) error
}

type OldPrinter struct{}

func (o *OldPrinter) Print(value string) error {
	fmt.Printf("Old printer: %s\n", value)
	return nil
}

type NewPrinter struct{}

func (n *NewPrinter) Print(value string) error {
	fmt.Printf("New printer: %s\n", value)
	return nil
}

func (n *NewPrinter) Send(value string) error {
	fmt.Printf("New printer: %s\n", value)
	return nil
}

func main() {
	o := &OldPrinter{}
	n := &NewPrinter{}

	o.Print("舊印表機只會列印")
	n.Print("新印表機能列印還有其他的功能")
	n.Send("新印表機也會傳真")
}

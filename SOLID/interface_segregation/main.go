package main

import "fmt"

/*
介面隔離原則
說明:
	- 客戶端不應被迫依賴它不需要的介面。
	- 與其建立大而全的介面，不如建立多個小而精的介面。
*/

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

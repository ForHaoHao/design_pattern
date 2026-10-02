package main

import "fmt"

/*
開放封閉原則
說明:
	- 對擴展開放，對修改封閉。
	- 當需要新功能時，透過增加程式碼來實現，而不是修改既有程式碼。
*/

// 反面教材，每次增加新的方式都要修改原程式碼
type PaymentProcessor struct{}

func (p *PaymentProcessor) Pay(method string, amount float32) error {
	switch method {
	case "LinePay":
		fmt.Printf("Line pay: %.2f\n", amount)
	case "CreditCard":
		fmt.Printf("Creadit card: %.2f\n", amount)
	default:
		return fmt.Errorf("Payment method not supported.")
	}
	return nil
}

// 正確寫法
// 付款行為介面
type PaymentMethod interface {
	Pay(amount float32) error
}

// 信用卡支付
type CreditCard struct {
	CardNumber string
}

func (c *CreditCard) Pay(amount float32) error {
	fmt.Printf("Creadit card %s: %.2f\n", c.CardNumber, amount)
	return nil
}

// Line 支付
type LinePay struct {
	LineNumber string
}

func (l *LinePay) Pay(amount float32) error {
	fmt.Printf("Line pay %s: %.2f\n", l.LineNumber, amount)
	return nil
}

// 封閉核心邏輯，只依賴介面
type Checkout struct{}

func (c *Checkout) Payment(method PaymentMethod, amount float32) error {
	method.Pay(amount)
	return nil
}

func main() {
	checkout := &Checkout{}

	card := &CreditCard{CardNumber: "HaoCard"}
	line := &LinePay{LineNumber: "HaoLine"}
	_ = checkout.Payment(card, 3.8)
	_ = checkout.Payment(line, 5.2)

}

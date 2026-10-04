package main

import "fmt"

/*
里式替換原則
說明:
	- 子類別必須能替代父類別使用，且不改變程式的正確性
	- 在 Go 語言中，可以理解為只要城市依賴的是某個介面，那麼任何一個符合這個介面的實作，都應該能夠安全替換
*/

type Cat interface {
	Breed() error
}

type persian struct{}

func (p *persian) Breed() error {
	fmt.Printf("波斯貓!\n")
	return nil
}

type britishShorthair struct{}

func (b *britishShorthair) Breed() error {
	fmt.Printf("英國短毛貓!\n")
	return nil
}

func GetCatBreed(cat Cat) error {
	cat.Breed()
	return nil
}

func main() {
	p := &persian{}
	b := &britishShorthair{}

	GetCatBreed(p)
	GetCatBreed(b)
}

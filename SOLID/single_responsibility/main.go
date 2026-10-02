package main

import "fmt"

/*
單一職責原則範例
說明:
	- 每個類別或模組只做一件事，並把它做好。這樣當需求變更時，修改的影響範圍會很小。
*/
// Email 服務
type EmailService struct{}

func (e *EmailService) Send() error {
	fmt.Printf("Send email to user.\n")
	return nil
}

// 使用者資料儲存
type UserRepository struct{}

func (u *UserRepository) Save() error {
	fmt.Printf("Store user profile in the database.\n")
	return nil
}

func main() {
	e := &EmailService{}
	e.Send()

	u := &UserRepository{}
	u.Save()
}

package main

import "fmt"

// ✅ Exercise 1: สร้าง interface ชื่อ Notifier
// 1. สร้าง interface
type Notifier interface {
	Send(to string, message string) error
}

// 2. สร้าง struct ที่ implements interface
type EmailNotifier struct{}

func (e EmailNotifier) Send(to, message string) error {
	fmt.Println("📧 ส่ง Email ถึง", to, ":", message)
	return nil
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(to, message string) error {
	fmt.Println("📱 ส่ง SMS ถึง", to, ":", message)
	return nil
}

// --------------------------

// ✅ Exercise 2: ใช้ DI กับ struct App
type App struct {
	Notifier Notifier
}

func (a App) Run() {
	a.Notifier.Send("admin@system.com", "ระบบกำลังทำงาน")
}

// --------------------------

func main() {
	// ✅ Exercise 1:
	var n Notifier

	n = EmailNotifier{}
	n.Send("user@example.com", "สวัสดีจากระบบ Email")

	n = SMSNotifier{}
	n.Send("0812345678", "รหัส OTP ของคุณคือ 123456")

	// ✅ Exercise 2:
	app := App{Notifier: EmailNotifier{}}
	app.Run()

}

package main

import (
	"fmt"
)

// 🟢 STEP 1: รู้จัก Interface แบบง่ายที่สุด
type Notifier interface {
	Send(to string, message string) error
}

// 🟢 STEP 2: สร้างโครงสร้างที่ “implements” interface
type EmailNotifier struct{}

func (e EmailNotifier) Send(to, message string) error {
	fmt.Println("📧 ส่งอีเมลถึง:", to, "ข้อความ:", message)
	return nil
}

// 🟢 STEP 3: ฟังก์ชันที่รับ Interface
func notifyUser(n Notifier) {
	n.Send("user@example.com", "Hello world!")
}

func main() {
	email := EmailNotifier{}
	notifyUser(email)
	// ✅ ทำงานได้ เพราะ EmailNotifier “มี method” ที่ interface ต้องการ

	app := App{Notifier: EmailNotifier{}}
	app.Run()
}

// 🟡 Concept สำคัญ: Dependency Injection (DI)
// Dependency Injection คือ การส่ง “สิ่งที่ต้องใช้” เข้ามาแทนที่จะสร้างเอง
type App struct {
	Notifier Notifier
}

func (a App) Run() {
	a.Notifier.Send("admin@system.com", "ระบบเริ่มแล้ว")
}

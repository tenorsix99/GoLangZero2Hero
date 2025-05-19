package main

import (
	"context"
	"fmt"
	"time"
)

// 1. สร้าง Notifier interface
type Notifier interface {
	Send(to, message string) error
}

// 2. ตัวอย่าง EmailNotifier
type EmailNotifier struct{}

func (e EmailNotifier) Send(to, msg string) error {
	time.Sleep(1 * time.Second) // จำลอง delay
	fmt.Println("📧 Email sent to", to, ":", msg)
	return nil
}

// 3. ฟังก์ชัน notifyAll ใช้ context + goroutine + timeout
func notifyAll(ctx context.Context, users []string, notifier Notifier) {
	for _, user := range users {
		go func(u string) {
			childCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
			defer cancel()

			done := make(chan error, 1)

			go func() {
				err := notifier.Send(u, "สวัสดีคุณลูกค้า!")
				done <- err
			}()

			select {
			case <-childCtx.Done():
				fmt.Println("⛔ Timeout ส่งถึง", u)
			case err := <-done:
				if err != nil {
					fmt.Println("❌ Error:", err)
				}
			}
		}(user)
	}
}

func main() {
	users := []string{"user1@example.com", "user2@example.com", "user3@example.com"}
	notifier := EmailNotifier{}
	ctx := context.Background()

	notifyAll(ctx, users, notifier)

	// wait ให้ goroutine ส่งครบ (เพื่อไม่ให้โปรแกรมจบก่อน)
	time.Sleep(3 * time.Second)
}

// 🧠 อธิบายสำคัญ:
// | ส่วน                          | ทำอะไร                              |
// | ----------------------------- | ----------------------------------- |
// | `context.WithTimeout`         | จำกัดเวลาไม่ให้ส่งนานเกิน           |
// | `done := make(chan error, 1)` | รอฟังผลลัพธ์การส่ง                  |
// | `select`                      | ถ้าส่งช้ากว่าที่กำหนด → log timeout |
// | goroutine                     | ส่งแต่ละคนแยก thread (พร้อมกัน)     |

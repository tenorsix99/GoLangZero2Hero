// 🧠 โจทย์: Service A → B → C แบบจำลอง (ในไฟล์เดียว)
// Service A: รับคำสั่งเริ่มจาก main()
// Service B: ประมวลผลข้อมูล
// Service C: ทำงานล่าช้า ถ้า timeout → ถูกยกเลิก
// ใช้ context.WithTimeout() ตั้งไว้ที่ A แล้วส่งต่อไปจนถึง C

// ✅ ตัวอย่างโค้ด (จำลอง context propagation)

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ✅ A: เริ่ม context ที่ timeout 2 วินาที
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("🚀 Service A: เริ่มทำงาน")
	err := serviceB(ctx)
	if err != nil {
		fmt.Println("❌ Service A: ล้มเหลว:", err)
	} else {
		fmt.Println("✅ Service A: สำเร็จ")
	}
}

// ✅ B: ทำงานต่อ และส่ง context ไปยัง C
func serviceB(ctx context.Context) error {
	fmt.Println("➡️ Service B: เรียก Service C")

	result, err := serviceC(ctx)
	if err != nil {
		return fmt.Errorf("Service B: C ล้มเหลว: %w", err)
	}

	fmt.Println("📦 Service B: ได้ผลลัพธ์:", result)
	return nil
}

// ✅ C: ทำงานที่อาจใช้เวลานานเกิน
func serviceC(ctx context.Context) (string, error) {
	delay := time.Duration(rand.Intn(4)) * time.Second
	fmt.Printf("🕐 Service C: จะใช้เวลา %v\n", delay)

	select {
	case <-time.After(delay): // จำลองการทำงาน
		return "📄 ข้อมูลจาก Service C", nil
	case <-ctx.Done():
		return "", fmt.Errorf("⏰ ถูกยกเลิก: %v", ctx.Err())
	}
}

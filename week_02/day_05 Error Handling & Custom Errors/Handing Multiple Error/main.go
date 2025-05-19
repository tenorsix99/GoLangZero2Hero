// 📦 สถานการณ์ทั่วไป:
// Client Request → Go service
//                 ├─ call API A
//                 ├─ call API B
//                 └─ call API C
// ❗ ถ้า API ใด API หนึ่งล้มเหลว → ต้องตัดสินใจว่าจะ “หยุดเลย” หรือ “ข้ามไป” หรือ “retry”

// ✅ เทคนิคหลักในการจัดการ Error จากหลาย API

// ✅ 1. ตรวจสอบทุก error ทีละตัว
// ✅ 2. รวบรวม error หลายตัวเข้าด้วยกัน (Go 1.20+)
// ✅ 3. ทำงานแบบ parallel ด้วย Goroutine + channel + select
// ✅ 4. Retry เมื่อ error เป็นแบบ recoverable (เช่น timeout)

// ✅ โจทย์จำลอง:
// 1. เรียก Primary API (/delay/3) ที่ทำงานช้า
// 2. ถ้าเกิน timeout หรือพัง → Retry 3 ครั้ง
// 3. ถ้ายังพัง → เรียก Fallback API (/uuid) แทน

// 🔧 ตัวอย่างโค้ด Go (พร้อม Retry + Fallback)

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func callAPI(ctx context.Context, url string) (string, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}

func callWithRetry(ctx context.Context, url string, maxRetry int) (string, error) {
	var lastErr error
	for i := 1; i <= maxRetry; i++ {
		fmt.Printf("🔁 Attempt %d: Calling %s\n", i, url)

		// สร้าง context ย่อยแบบ timeout
		subCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		data, err := callAPI(subCtx, url)
		if err == nil {
			return data, nil
		}
		fmt.Println("⚠️  Error:", err)
		lastErr = err

		time.Sleep(500 * time.Millisecond) // delay ก่อน retry
	}
	return "", fmt.Errorf("retry failed: %w", lastErr)
}

func main() {
	ctx := context.Background()

	// 🧪 API หลักที่ delay 3 วิ (จะ timeout)
	primaryURL := "https://httpbin.org/delay/1"
	fallbackURL := "https://httpbin.org/uuid"

	fmt.Println("🚀 เรียก API หลัก พร้อม retry...")
	result, err := callWithRetry(ctx, primaryURL, 3)
	if err != nil {
		fmt.Println("❌ Primary ล้มเหลว → เรียก fallback")

		// 🛟 fallback ถ้า API หลักพัง
		data, err2 := callAPI(ctx, fallbackURL)
		if err2 != nil {
			fmt.Println("❌ Fallback ก็ล้มเหลว:", err2)
			return
		}
		fmt.Println("✅ สำเร็จจาก fallback:\n", data)
		return
	}

	fmt.Println("✅ สำเร็จจาก primary API:\n", result)
}

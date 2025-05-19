package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// 1. Call API ดึงข้อมูลลูกค้า
func fetchCustomerData(ctx context.Context) (string, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/2", nil)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("❌ ดึงข้อมูลลูกค้าไม่สำเร็จ: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}

// 2. Call API ส่งผลลัพธ์ไปยัง service อื่น
func sendNotification(ctx context.Context, message string) error {
	req, _ := http.NewRequestWithContext(ctx, "POST", "https://httpbin.org/post", nil)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("❌ แจ้งเตือนไม่สำเร็จ: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("📢 แจ้งเตือนสำเร็จ:", resp.Status)
	return nil
}

func main() {
	// กำหนด timeout ทั้งกระบวนการ 5 วินาที
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("🔄 เริ่มการประมวลผล...")

	// 1. ดึงข้อมูลลูกค้า
	customerData, err := fetchCustomerData(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("✅ ได้ข้อมูลลูกค้า:", customerData[:30], "...")

	// 2. ประมวลผล (จำลอง delay)
	fmt.Println("⚙️ ประมวลผลข้อมูล...")
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("✅ ประมวลผลสำเร็จ")
	case <-ctx.Done():
		fmt.Println("⛔ ประมวลผลถูกยกเลิก:", ctx.Err())
		return
	}

	// 3. ส่งแจ้งเตือน
	err = sendNotification(ctx, "ลูกค้าอัปเดตแล้ว")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("🎉 งานทั้งหมดเสร็จสิ้น")
}

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func callAPI(ctx context.Context, url string) error {
	// สร้าง HTTP request พร้อม context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("สร้าง request ไม่ได้: %v", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("เรียก API ไม่สำเร็จ: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("📦 ได้ข้อมูล:", string(body))
	return nil
}

func main() {
	// จำกัดเวลาแค่ 5 วินาที
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := "https://httpbin.org/delay/3" // จำลอง API ช้า (delay 3 วินาที)

	err := callAPI(ctx, url)
	if err != nil {
		fmt.Println("❌ ERROR:", err)
	} else {
		fmt.Println("✅ สำเร็จ")
	}
}

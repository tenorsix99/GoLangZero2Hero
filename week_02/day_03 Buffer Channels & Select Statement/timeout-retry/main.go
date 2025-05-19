package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doWork() <-chan string {
	result := make(chan string)
	go func() {
		sec := rand.Intn(4) // สุ่มระยะเวลา delay
		fmt.Println("🕒 งานนี้จะใช้เวลา", sec, "วินาที")
		time.Sleep(time.Duration(sec) * time.Second)
		result <- fmt.Sprintf("✅ งานสำเร็จหลัง %d วินาที", sec)
	}()
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	maxRetries := 3
	timeout := 2 * time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("🔁 พยายามครั้งที่ %d...\n", attempt)

		select {
		case msg := <-doWork():
			fmt.Println(msg)
			return // สำเร็จแล้ว ออกเลย
		case <-time.After(timeout):
			fmt.Println("⏰ Timeout! จะลองใหม่...")
		}
	}

	fmt.Println("❌ ลองครบทุกครั้งแล้ว ยังไม่สำเร็จ 😓")
}

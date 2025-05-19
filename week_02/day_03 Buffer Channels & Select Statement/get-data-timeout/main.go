package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dataChan := make(chan string)

	// ส่งข้อมูลล่าช้า (เกิน timeout)
	go func() {
		var sec int
		sec = rand.Intn(5) // สุ่ม 0–4 วินาที
		fmt.Println("⏳ กำลังรอ", sec, "วินาที")
		time.Sleep(time.Duration(sec) * time.Second)
		dataChan <- "ข้อมูลที่ช้ามาก"
	}()

	select {
	case msg := <-dataChan:
		fmt.Println("✅ ได้ข้อมูล:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("⏰ Timeout: รอข้อมูลไม่ไหวแล้ว")
	}
}

package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from Goroutine")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// ✅ สร้าง Goroutine

	go hello() // ❗ เรียก go แล้ว function นี้จะทำงานเบื้องหลัง
	fmt.Println("Hello from Main")
	time.Sleep(1 * time.Second) // รอให้ Goroutine ทำงานทัน

	// ✅ รันหลาย Goroutines พร้อมกัน
	go say("world") // ทำงานเบื้องหลัง
	say("hello")    // ทำงานปกติ

	// Goroutine ไม่ต้องรอ Main
	// ถ้า Main function จบก่อน Goroutine → Goroutine ก็โดน kill ทันที! ❗
	// เพราะฉะนั้น บางครั้งเราต้อง
	// - ใช้ time.Sleep
	// - หรือใช้ advanced concept เช่น WaitGroup

}

// ✅ อะไรคือ Concurrency?
// Concurrency คือการจัดการหลายงาน (task) พร้อม ๆ กัน
// Go ทำได้ง่ายมากด้วย Goroutine
// Goroutine คือ "function" ที่ทำงานพร้อมกันแบบเบา ๆ (lightweight thread)
// (Go สามารถสร้าง ล้าน goroutines ได้โดยไม่พัง)

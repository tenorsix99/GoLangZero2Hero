package main

import (
	"fmt"
	"math/rand"
	"time"
)

// โครงสร้างข้อมูลผลลัพธ์ stock
type StockResult struct {
	Warehouse string
	Available int
}

// mock function ดึงข้อมูล stock จากคลัง
func checkStock(warehouse string, itemCode string, ch chan<- StockResult) {
	// จำลองความล่าช้า (เช่น query database)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	// mock stock value (random 0-100)
	available := rand.Intn(100)

	// ส่งผลลัพธ์กลับ channel
	ch <- StockResult{
		Warehouse: warehouse,
		Available: available,
	}
}

func main() {
	itemCode := "ABC123"
	warehouses := []string{"Warehouse A", "Warehouse B", "Warehouse C"}

	// สร้าง channel สำหรับรับผลลัพธ์ stock
	ch := make(chan StockResult)

	fmt.Println("🔍 ตรวจสอบสินค้า:", itemCode)

	// ดึง stock แบบขนาน
	for _, wh := range warehouses {
		go checkStock(wh, itemCode, ch)
	}

	// รวมผลลัพธ์ทั้งหมด
	for i := 0; i < len(warehouses); i++ {
		result := <-ch
		fmt.Printf("✅ %s เหลือสินค้าอยู่: %d ชิ้น\n", result.Warehouse, result.Available)
	}

}

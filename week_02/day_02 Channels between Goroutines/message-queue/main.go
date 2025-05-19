package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ข้อมูลคำสั่งซื้อ
type Order struct {
	OrderID  int
	ItemCode string
	Quantity int
}

// ผลลัพธ์การตรวจ stock
type StockCheckResult struct {
	OrderID  int
	InStock  bool
	ItemCode string
}

// Mock ตรวจ stock (สุ่มว่ามีของหรือไม่)
func checkStock(order Order, resultChan chan<- StockCheckResult) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000))) // simulate delay
	inStock := rand.Intn(2) == 1                                  // สุ่มมี/ไม่มี

	resultChan <- StockCheckResult{
		OrderID:  order.OrderID,
		InStock:  inStock,
		ItemCode: order.ItemCode,
	}
}

// ระบบแจ้งเตือนผลลัพธ์ (consumer)
func notificationService(resultChan <-chan StockCheckResult) {
	for result := range resultChan {
		if result.InStock {
			fmt.Printf("✅ Order %d: ItemCode %s สินค้ามีใน stock\n", result.OrderID, result.ItemCode)
		} else {
			fmt.Printf("❌ Order %d: ItemCode %s สินค้าหมด\n", result.OrderID, result.ItemCode)
		}
	}
}

func main() {
	orders := []Order{
		{OrderID: 1, ItemCode: "A123", Quantity: 2},
		{OrderID: 2, ItemCode: "B456", Quantity: 1},
		{OrderID: 3, ItemCode: "C789", Quantity: 5},
		{OrderID: 4, ItemCode: "A123", Quantity: 1},
	}

	resultChan := make(chan StockCheckResult)

	// start consumer (แจ้งเตือน)
	go notificationService(resultChan)

	// สั่งตรวจ stock แบบขนาน
	for _, order := range orders {
		go checkStock(order, resultChan)
	}

	// รอให้ทุกคำสั่งทำงานเสร็จ
	time.Sleep(2 * time.Second)
	close(resultChan) // ปิด channel หลังส่งครบ
}

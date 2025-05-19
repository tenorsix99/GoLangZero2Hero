package main

import (
	"fmt"
	"time"
)

// 🔹 Step 1: Channel พื้นฐาน
func main_1() {
	ch := make(chan string) // สร้าง channel สำหรับ string

	go func() {
		ch <- "Hello from goroutine!" // ส่งข้อความไปที่ channel
	}()

	msg := <-ch // รับข้อความจาก channel (จะรอจนกว่ามีคนส่ง)
	fmt.Println("ได้รับข้อความ:", msg)
	// ❗ ถ้าไม่มี <-ch → โปรแกรมจะค้างอยู่ เพราะไม่มีใครรอรับ
}

// 🔹 Step 2: ส่งหลายค่า → ต้องปิด Channel
func main_2() {
	ch := make(chan int)

	// ส่งข้อมูลจาก Goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // ✅ ต้องปิด channel ถ้าจะใช้ for range
		// 📌 ถ้าไม่ปิด ch → for range ch จะค้างอยู่ตลอดไป
	}()

	// รับค่าทั้งหมด
	for value := range ch {
		fmt.Println("ได้ค่า:", value)
	}
}

// 🔹 Step 3: ฟังก์ชันส่งค่า + channel เป็น parameter
func sayHi(ch chan string) {
	ch <- "สวัสดีจาก sayHi()"
}

func main_3() {
	ch := make(chan string)
	go sayHi(ch)
	fmt.Println(<-ch) // รอรับข้อความ
}

// 🔹 Step 4: Directional Channel (ส่งเท่านั้น / รับเท่านั้น)
func sendData(ch chan<- int) { // ส่งเท่านั้น
	ch <- 42
}

func receiveData(ch <-chan int) { // รับเท่านั้น
	fmt.Println("รับได้:", <-ch)
}

func main_4() {
	ch := make(chan int)
	go sendData(ch)
	go receiveData(ch)
	time.Sleep(1 * time.Second)
}

// | พฤติกรรม              	| อธิบาย                            	  |
// | --------------------- | --------------------------------- 		|
// | `ch <- val`           | ส่งค่าเข้า channel                		   |
// | `<-ch`                | รับค่าจาก channel                 		  |
// | `make(chan T)`        | สร้าง channel ที่ส่งค่าประเภท `T` 			 |
// | `for val := range ch` | รับค่าทั้งหมดจนปิด channel         		 |
// | `close(ch)`           | ปิด channel                        	 |
// | `chan<-`, `<-chan`    | จำกัดทิศทางส่ง / รับ              			|

// struct สำหรับผลลัพธ์ log
type LogResult struct {
	Source string
	Data   string
}

// mock service
func fetchLog(source string, ch chan<- LogResult) {
	// mock delay ต่างกัน
	time.Sleep(time.Duration(len(source)) * 200 * time.Millisecond)
	ch <- LogResult{
		Source: source,
		Data:   fmt.Sprintf("Logs from %s at %s", source, time.Now().Format(time.Stamp)),
	}
}

func main() {
	sources := []string{"web-server", "database", "auth-service"}
	ch := make(chan LogResult)

	// สร้าง goroutine สำหรับดึง log จากทุก source
	for _, src := range sources {
		go fetchLog(src, ch)
	}
	resultCh := <-ch
	fmt.Print("Channel resilt => ", resultCh)

	// รอรับผลลัพธ์ทั้งหมด
	for i := 0; i < len(sources); i++ {
		log := <-ch
		fmt.Printf("✅ [%s] %s\n", log.Source, log.Data)
	}

}

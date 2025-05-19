package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job = สิ่งที่ worker ต้องประมวลผล
type Job struct {
	ID       int
	TaskName string
}

// Result = ข้อมูลที่ใช้แจ้งเตือนหลังจาก job เสร็จ
type Result struct {
	JobID    int
	TaskName string
	Success  bool
}

// worker ประมวลผลงาน แล้วส่งผลลัพธ์ไปแจ้งเตือน
func worker(id int, jobs <-chan Job, notify chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("👷‍♂️ Worker %d: ทำ %s (JobID: %d)\n", id, job.TaskName, job.ID)

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // จำลองการทำงาน
		success := rand.Intn(10) != 0                                // 10% fail

		fmt.Printf("✅ Worker %d: เสร็จ %s (%v)\n", id, job.TaskName, success)

		// ส่งผลลัพธ์ให้ระบบแจ้งเตือน
		notify <- Result{
			JobID:    job.ID,
			TaskName: job.TaskName,
			Success:  success,
		}
	}
}

// ฟังก์ชันแจ้งเตือนผลลัพธ์ (consumer)
func notificationService(notify <-chan Result, done chan<- bool) {
	for result := range notify {
		if result.Success {
			fmt.Printf("📣 แจ้งเตือน: งาน %s เสร็จเรียบร้อย (JobID: %d)\n", result.TaskName, result.JobID)
		} else {
			fmt.Printf("⚠️ แจ้งเตือน: งาน %s ล้มเหลว (JobID: %d)\n", result.TaskName, result.JobID)
		}
	}
	done <- true // บอกว่าแจ้งเตือนจบแล้ว
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	notify := make(chan Result, numJobs)
	done := make(chan bool)

	var wg sync.WaitGroup

	// สร้าง worker
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, notify, &wg)
	}

	// สร้าง Goroutine สำหรับแจ้งเตือน
	go notificationService(notify, done)

	// สร้างงาน
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{
			ID:       j,
			TaskName: fmt.Sprintf("ส่งรายงาน %d", j),
		}
	}
	close(jobs) // ปิดช่องส่งงาน

	wg.Wait()     // รอให้ worker ทำงานเสร็จ
	close(notify) // ปิดแจ้งเตือนเมื่อทำครบ
	<-done        // รอให้ระบบแจ้งเตือนทำงานเสร็จ

	fmt.Println("🎉 ระบบทำงานครบทุกส่วนแล้ว")
}

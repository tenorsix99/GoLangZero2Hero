package main

import (
	"fmt"
	"testing"
)

// 🟠 STEP 4: Mock สำหรับการทดสอบ
type MockNotifier struct {
	Called bool
}

func (m *MockNotifier) Send(to, message string) error {
	m.Called = true
	return nil
}

func TestMain(t *testing.T) {
	mock := &MockNotifier{}
	app := App{Notifier: mock}
	app.Run()

	fmt.Println("ถูกเรียกไหม?", mock.Called)
}

// ✅ สรุป flow
// interface → struct implements interface → inject interface → ใช้งาน

// !! อย่าลืม go mod init แล้ว go get package ก่อน test

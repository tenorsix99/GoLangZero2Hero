package main

import (
	"fmt"
	"testing"
)

type MockNotifier struct {
	Called  bool
	Message string
}

func (m *MockNotifier) Send(to, message string) error {
	m.Called = true
	m.Message = message
	fmt.Println("✅ [MOCK] เรียกใช้ mock send กับข้อความ:", message)
	return nil
}

func TestMain(t *testing.T) {
	mock := &MockNotifier{}
	app := App{Notifier: mock}
	app.Run()

	fmt.Println("ตรวจสอบว่า Mock ถูกเรียก:", mock.Called)
	fmt.Println("ข้อความที่ถูกส่ง:", mock.Message)
}

// !! อย่าลืม go mod init แล้ว go get package ก่อน test

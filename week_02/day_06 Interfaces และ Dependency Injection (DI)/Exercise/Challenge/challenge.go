package main

import (
	"fmt"
)

type OTPService interface {
	SendOTP(to string, code string) error
}

type EmailOTP struct{}

func (e EmailOTP) SendOTP(to, code string) error {
	fmt.Println("📧 ส่ง OTP:", code, "ไปยัง Email:", to)
	return nil
}

type SMSOTP struct{}

func (s SMSOTP) SendOTP(to, code string) error {
	fmt.Println("📱 ส่ง OTP:", code, "ไปยังเบอร์:", to)
	return nil
}

func sendOTPToUser(svc OTPService, to string) {
	code := "789456"
	svc.SendOTP(to, code)
}

func main() {
	sendOTPToUser(EmailOTP{}, "user@example.com")
	sendOTPToUser(SMSOTP{}, "0812345678")
}

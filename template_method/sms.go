package template_method

import "fmt"

type Sms struct {
	Otp Otp
}

func (s Sms) GenerateRandomOTP(i int) string {
	randomOtp := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOtp)
	return randomOtp
}

func (s Sms) SaveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s Sms) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s Sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms : %s\n", message)
	return nil
}

func (s Sms) PublishMetric() {
	fmt.Println("SMS: publishing mertics")
}

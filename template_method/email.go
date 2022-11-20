package template_method

import "fmt"

type Email struct {
	Otp Otp
}

func (e Email) GenerateRandomOTP(i int) string {
	randomOtp := "1234"
	fmt.Printf("Email: generating random otp %s\n", randomOtp)
	return randomOtp
}

func (e Email) SaveOTPCache(otp string) {
	fmt.Printf("Email: saving otp: %s to cache\n", otp)
}

func (e Email) GetMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e Email) SendNotification(message string) error {
	fmt.Printf("Email: sending email : %s\n", message)
	return nil
}

func (e Email) PublishMetric() {
	fmt.Println("Email: publishing mertics")
}

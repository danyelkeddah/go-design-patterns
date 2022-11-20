package template_method

type Otp struct {
	Otp OtpContract
}

func (o *Otp) GenerateAndSendOtp(otpLength int) error {
	otp := o.Otp.GenerateRandomOTP(otpLength)
	o.Otp.SaveOTPCache(otp)
	message := o.Otp.GetMessage(otp)
	err := o.Otp.SendNotification(message)
	if err != nil {
		return err
	}
	o.Otp.PublishMetric()
	return nil
}

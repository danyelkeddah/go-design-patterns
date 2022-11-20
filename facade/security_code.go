package facade

import "fmt"

type SecurityCode struct {
	Code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{Code: code}
}

func (s *SecurityCode) CheckCode(incomingCode int) error {
	if s.Code != incomingCode {
		return fmt.Errorf("securityt code is incorrect")
	}
	fmt.Println("Security code verified")
	return nil
}

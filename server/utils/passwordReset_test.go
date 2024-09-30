package utils_test

import (
	"task-inator3000/utils"
	"testing"
)

func TestGenerateOTP(t *testing.T) {
	result := utils.GenerateOTP()

	if len(result) != utils.OTPLength {
		t.Errorf("Length of OTP was not %v. OTP: %v", utils.OTPLength, result)
	}
}

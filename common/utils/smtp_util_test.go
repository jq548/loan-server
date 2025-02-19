package utils

import "testing"

func TestSendEmail(t *testing.T) {

	email, err := SendEmail(
		false,
		"liamc1878@gmail.com",
		"liamc1878@gmail.com",
		"fhqx ayeq fmii waoh",
		"smtp.gmail.com",
		587,
		"1140830756@qq.com",
		"hello",
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(email)
}

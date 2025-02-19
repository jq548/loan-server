package utils

import "testing"

func TestSendEmail(t *testing.T) {

	email, err := SendEmail(
		false,
		"",
		"",
		"",
		"smtp.gmail.com",
		587,
		"",
		"",
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(email)
}

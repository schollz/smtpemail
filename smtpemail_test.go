package smtpemail

import "testing"

func TestSend(t *testing.T) {
	err := Send("someone@something.com", "someoneelse@somethingelse.com", "Hi there", "*This* is an email", "smtpemail.go")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

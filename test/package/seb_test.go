package seb

import "testing"
func TestSeb(t *testing.T) {
	want := "Helo World"
	if  got := Seb(); got != want {
		t.Errorf("Seb() est bug")
	}
}

package main

import (
	"text/template"
	"testing"
)

func TestHTML(t *testing.T) {
	// FIXME: htmlが妥当でなくてもOKになる
	_, err := template.ParseFiles("hogehoge")

	if err != nil {
		t.Error("html is invalid")
	}
}

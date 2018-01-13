package main

import "testing"

func TestParse(t *testing.T) {
	contents := `<html>
	<a href="https://google.com" />
	<a href="https://yahoo.co.jp" />
	</html>`

	Parse(contents)
}

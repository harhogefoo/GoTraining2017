// Echo1は、そのコマンドライン引数を表示します。
package ex01

import (
	"fmt"
	"testing"
)

type Tests struct {
	num float64
	want string
}

func getCelsius() []Tests {
	return []Tests {
		{32, "32 C"},
		{100.5, "100.5 C"},
	}
}

func getFahrenheit() []Tests {
	return []Tests {
		{32, "32 F"},
		{100.5, "100.5 F"},
	}
}

func getKelvin() []Tests {
	return []Tests {
		{32, "32 K"},
		{100.5, "100.5 K"},
	}
}

func TestCelsius_String(t *testing.T) {
	for _, test := range getCelsius() {
		description := fmt.Sprintf("tempconv(%q", test.num)
		c := Celsius(test.num)
		if c.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, c, test.want)
		}
	}
}

func TestFahrenheit_String(t *testing.T) {
	for _, test := range getFahrenheit() {
		description := fmt.Sprintf("tempconv(%q", test.num)
		f := Fahrenheit(test.num)
		if f.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, f, test.want)
		}
	}
}

func TestKelvin_String(t *testing.T) {
	for _, test := range getKelvin() {
		description := fmt.Sprintf("tempconv(%q", test.num)
		k := Kelvin(test.num)
		if k.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, k, test.want)
		}
	}
}

func TestCToF(t *testing.T) {
	var tests = []Tests {
		{32, "89.6 F"},
		{10.5, "50.9 F"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := CToF(Celsius(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []Tests {
		{32, "-241.14999999999998 K"},
		{10.5, "-262.65 K"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := CToK(Celsius(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []Tests {
		{32, "0 C"},
		{10.5, "-11.944444444444445 C"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := FToC(Fahrenheit(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []Tests {
		{32, "-241.14999999999998 C"},
		{10.5, "-262.65 C"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := KToC(Kelvin(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}







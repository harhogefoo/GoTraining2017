// Echo1は、そのコマンドライン引数を表示します。
package ex01

import (
	"fmt"
	"testing"
	"github.com/harhogefoo/go_training2017/ch02/ex01/tempconv"
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
		c := tempconv.Celsius(test.num)
		if c.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, c, test.want)
		}
	}
}

func TestFahrenheit_String(t *testing.T) {
	for _, test := range getFahrenheit() {
		description := fmt.Sprintf("tempconv(%q", test.num)
		f := tempconv.Fahrenheit(test.num)
		if f.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, f, test.want)
		}
	}
}

func TestKelvin_String(t *testing.T) {
	for _, test := range getKelvin() {
		description := fmt.Sprintf("tempconv(%q", test.num)
		k := tempconv.Kelvin(test.num)
		if k.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, k, test.want)
		}
	}
}

func TestCToF(t *testing.T) {
	if tempconv.CToF(tempconv.BoilingC) != 212.0 {
		t.Error(fmt.Sprint(tempconv.CToF(tempconv.BoilingC)))
	}

	var tests = []Tests {
		{32, "89.6 F"},
		{10.5, "50.9 F"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := tempconv.CToF(tempconv.Celsius(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	if tempconv.CToK(0) != -tempconv.Kelvin(tempconv.AbsoluteZeroC) {
		t.Error(fmt.Sprint(tempconv.CToK(0)))
	}
	var tests = []Tests {
		{32, "-241.14999999999998 K"},
		{10.5, "-262.65 K"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := tempconv.CToK(tempconv.Celsius(test.num))
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
		got := tempconv.FToC(tempconv.Fahrenheit(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestKToC(t *testing.T) {
	if tempconv.KToC(0) != tempconv.AbsoluteZeroC {
		t.Error(fmt.Sprint(tempconv.KToC(0)))
	}
	var tests = []Tests {
		{32, "-241.14999999999998 C"},
		{10.5, "-262.65 C"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := tempconv.KToC(tempconv.Kelvin(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

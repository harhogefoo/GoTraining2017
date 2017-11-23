// Echo1は、そのコマンドライン引数を表示します。
package ex02

import (
	"fmt"
	"testing"
	"github.com/harhogefoo/go_training2017/ch02/ex02/tempconv"
	"github.com/harhogefoo/go_training2017/ch02/ex02/lengthconv"
	"github.com/harhogefoo/go_training2017/ch02/ex02/weightconv"
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

func TestFeet_String(t *testing.T) {
	var tests = []Tests {
		{32, "32 feet"},
		{10.5, "10.5 feet"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := lengthconv.Feet(test.num)
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestMeter_String(t *testing.T) {
	var tests = []Tests {
		{32, "32 meter"},
		{10.5, "10.5 meter"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := lengthconv.Meter(test.num)
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestFToM(t *testing.T) {
	var tests = []Tests {
		{32, "9.7536 meter"},
		{10.5, "3.2004 meter"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := lengthconv.FToM(lengthconv.Feet(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}

}

func TestMToF(t *testing.T) {
	var tests = []Tests {
		{32, "104.98687664041994 feet"},
		{10.5, "34.44881889763779 feet"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := lengthconv.MToF(lengthconv.Meter(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestKilogram_String(t *testing.T) {
	var tests = []Tests {
		{32, "32 kilogram"},
		{10.5, "10.5 kilogram"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := weightconv.Kilogram(test.num)
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestPound_String(t *testing.T) {
	var tests = []Tests {
		{32, "32 pound(lb)"},
		{10.5, "10.5 pound(lb)"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := weightconv.Pound(test.num)
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}

func TestKToP(t *testing.T) {
	var tests = []Tests {
		{32, "70.54798144588088 pound(lb)"},
		{10.5, "23.148556411929665 pound(lb)"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := weightconv.KToP(weightconv.Kilogram(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}

}

func TestPToK(t *testing.T) {
	var tests = []Tests {
		{32, "14.514944 kilogram"},
		{10.5, "4.762716 kilogram"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("tempconv(%q", test.num)
		got := weightconv.PToK(weightconv.Pound(test.num))
		if got.String() != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}




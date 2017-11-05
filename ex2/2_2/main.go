// cfは、その数値引数を摂氏と華氏へ変換します
package main

import (
	"os"
	"strconv"
	"fmt"
	"github.com/harhogefoo/go_training2017/ch2/tempconv"
	"github.com/harhogefoo/go_training2017/ex2/lengthconv"
	"github.com/harhogefoo/go_training2017/ex2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		fah := tempconv.Fahrenheit(t)
		cel := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			fah, tempconv.FToC(fah), cel, tempconv.CToF(cel))

		feet := lengthconv.Feet(t)
		meter := lengthconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			feet, lengthconv.FToM(feet), meter, lengthconv.MToF(meter))

		kilogram := weightconv.Kilogram(t)
		pound := weightconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n",
			kilogram, weightconv.KToP(kilogram), pound, weightconv.PToK(pound))
	}
}



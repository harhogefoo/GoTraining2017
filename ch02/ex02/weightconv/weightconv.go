package weightconv

import "fmt"

type Kilogram float64
type Pound float64

const (
	OnePoundPerKilogram = 0.453592
)

func (k Kilogram) String() string    { return fmt.Sprintf("%g kilogram", k) }
func (p Pound) String() string { return fmt.Sprintf("%g pound(lb)", p) }

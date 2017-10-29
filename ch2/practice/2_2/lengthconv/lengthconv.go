package lengthconv

import "fmt"

type Feet float64
type Meter float64

const (
	OneFeetPerMeter = 0.3048
)

func (f Feet) String() string    { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string { return fmt.Sprintf("%g meter", m) }
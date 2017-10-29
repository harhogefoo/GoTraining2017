package lengthconv

// FtoM はFeetをMeterへ変換します。
func FToM(f Feet) Meter { return Meter(f * OneFeetPerMeter) }

// MtoF はMeterをFeetへ変換します。
func MToF(m Meter) Feet { return Feet(m / OneFeetPerMeter) }

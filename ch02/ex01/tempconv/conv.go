package tempconv

// CtoF は摂氏の温度を華氏へ変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

// FtoC は華氏の温度を摂氏へ変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CtoK は摂氏の温度を絶対温度へ変換します。
func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

// KtoC は絶対温度を摂氏へ変換します。
func KToC(k Kelvin) Celsius { return Celsius(k + -273.15) }
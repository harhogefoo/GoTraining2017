package weightconv

// kilogramをpoundに変換します。
func KToP(k Kilogram) Pound { return Pound(k / OnePoundPerKilogram) }

// Poundをkilogramに変換します。
func PToK(p Pound) Kilogram { return Kilogram(p * OnePoundPerKilogram) }

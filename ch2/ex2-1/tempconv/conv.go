package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func FToC(f Fahrenheit) Celsius {
	// A conversion is needed because Fahrenheit can't be directly assigned to Celsius
	return Celsius((f - 32) * 5 / 9)
}
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

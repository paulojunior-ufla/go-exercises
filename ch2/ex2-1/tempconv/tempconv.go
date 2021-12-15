// Copyright © 2021 Paulo A. P. Júnior
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Based on the "Fetchall" program by Alan A. A. Donovan & Brian W. Kernighan.
// https://github.com/adonovan/gopl.io/blob/master/ch2/tempconv/tempconv.go

// Ex2-1 performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	// It works because -273.15 may be assigned to Celsius (both are float64)
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0.0
	BoilingC      Celsius = 100.0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

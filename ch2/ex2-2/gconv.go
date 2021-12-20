// Copyright © 2021 Paulo A. P. Júnior
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Based on the "Tempconv" program by Alan A. A. Donovan & Brian W. Kernighan.
// https://github.com/adonovan/gopl.io/tree/master/ch2/tempconv

// Ex2-2 performs Celsius and Fahrenheit conversions.
package main

import "fmt"

// Temperature
type Celsius float64
type Fahrenheit float64

func (t Celsius) String() string    { return fmt.Sprintf("%.2f°C", t) }
func (t Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", t) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// Distance
type Ft float64
type Km float64

func (d Ft) String() string { return fmt.Sprintf("%.2f ft", d) }
func (d Km) String() string { return fmt.Sprintf("%.2f km", d) }

func FtToKm(d Ft) Km { return Km(d * 0.0003048) }
func KmToFt(d Km) Ft { return Ft(d / 0.0003048) }

// Weight
type Kg float64
type Lb float64

func (w Kg) String() string { return fmt.Sprintf("%.2f kg", w) }
func (w Lb) String() string { return fmt.Sprintf("%.2f lb", w) }

func KgToLb(w Kg) Lb { return Lb(w * 2.20462) }
func LbToKg(w Lb) Kg { return Kg(w / 2.20462) }

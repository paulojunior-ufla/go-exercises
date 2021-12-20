// Copyright © 2021 Paulo A. P. Júnior
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Based on the "Popcount" program by Alan A. A. Donovan & Brian W. Kernighan.
// https://github.com/adonovan/gopl.io/tree/master/ch2/popcount

// Ex2-3 returns the population count (number of set bits) of x.
package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

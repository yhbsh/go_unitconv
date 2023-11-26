// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"playground/unitconv"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := unitconv.Fahrenheit(t)
		c := unitconv.Celsius(t)
		m := unitconv.Meter(t)
		ft := unitconv.Feet(t)

		fmt.Printf("%02s = %02s, %02s = %02s, %02s = %02s, %02s = %02s\n",
			f, unitconv.FToC(f).String(),
			c, unitconv.CToF(c).String(),
			m, unitconv.MToF(m).String(),
			ft, unitconv.FToM(ft).String())
	}
}

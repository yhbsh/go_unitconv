// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"playground/tempconv"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		m := tempconv.Meter(t)
		ft := tempconv.Feet(t)

		fmt.Printf("%02s = %02s, %02s = %02s, %02s = %02s, %02s = %02s\n",
			f, tempconv.FToC(f).String(),
			c, tempconv.CToF(c).String(),
			m, tempconv.MToF(m).String(),
			ft, tempconv.FToM(ft).String())
	}
}

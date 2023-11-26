package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Feet float64
type Meter float64

func (c Celsius) String() string {
	return fmt.Sprintf("%02gºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%02gºF", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%02gºK", k)
}

func (f Feet) String() string {
	return fmt.Sprintf("%02gft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%02gm", m)
}

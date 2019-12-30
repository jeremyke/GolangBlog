package tempconv

import "fmt"

type Celsius float32
type Fahrenheit float32

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g c", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g f", f)
}

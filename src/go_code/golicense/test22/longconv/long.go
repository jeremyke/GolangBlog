package longconv

import "fmt"

type Foot float64
type Mile float64

func (f Foot) String() string {
	return fmt.Sprintf("%g英尺", f)
}
func (m Mile) String() string {
	return fmt.Sprintf("%g米", m)
}

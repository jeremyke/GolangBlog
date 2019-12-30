package longconv

func FToM(f Foot) Mile {
	return Mile(f / 39.37)
}
func MToF(m Mile) Foot {
	return Foot(m * 39.37)
}

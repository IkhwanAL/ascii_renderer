package utils

func ClampUint8(value float64, min, max float64) uint8 {
	if value > max {
		return 255
	}

	if value < min {
		return 0
	}
	return uint8(value)
}

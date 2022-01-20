package TypeNumber

func TypeNumber(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a > 0:
		return "positive"
	default:
		return "zero"
	}
}

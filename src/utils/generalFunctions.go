package utils

func ConvBoolToStr(value bool) string {
	if value{
		return "T"
	}
	return "F"
}

func ConvIntToBool(value int) bool {
	if value == 0 {
		return false
	} else {
		return true
	}
}
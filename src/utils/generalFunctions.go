package utils

func ConvBoolToStr(value bool) string {
	if value{
		return "T"
	}
	return "F"
}

func ConvStrToBool(value string) bool {
	if value == "T" {
		return true
	}
	return false
}
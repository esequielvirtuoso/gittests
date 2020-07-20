package main

import u "github.com/esequielvirtuoso/gittests/src/utils"

func main() {
	u.PrintVars("Boolean",u.ConvBoolToStr(true))
	u.PrintVars("Integer",u.ConvIntToBool(1))
}

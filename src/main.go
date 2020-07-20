package main

import "fmt"
import u "github.com/esequielvirtuoso/gittests/src/utils"

func main() {
	fmt.Println(u.ConvBoolToStr(true))
	fmt.Println(u.ConvStrToBool("F"))
}

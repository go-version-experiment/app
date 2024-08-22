package main

import (
	"fmt"

	"github.com/go-version-experiment/baselib"
	"github.com/go-version-experiment/seclib"
)

func main() {
	fmt.Printf("main: baselib: Version=%s\tGetVersion()=%s\n", baselib.Version, baselib.GetVersion())
	fmt.Printf("main: seclib:  Version=%s\tGetVersion()=%s\n", seclib.Version, seclib.GetVersion())
}

package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/sivchari/unusedtypeparam"
)

func main() { unitchecker.Main(unusedtypeparam.Analyzer) }

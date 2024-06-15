package main

import (
	declarationcomment "github.com/cold-bin/declaration-comment/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(declarationcomment.Analyzer)
}

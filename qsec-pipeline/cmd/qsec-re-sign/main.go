package main

import (
	"log"
	"os"

	"github.com/qsec-pipeline/qsec-pipeline/internal/re-sign"
)

func main() {
	if err := resign.Run(); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"os"

	cts "github.com/ryuichi1208/go-check-tcp-status/lib"
)

func main() {
	os.Setenv("LANG", "C")
	os.Setenv("LC_ALL", "C")
	cts.Do()
}

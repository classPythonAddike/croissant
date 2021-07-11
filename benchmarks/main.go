package main

import (
	"flag"
)

func main() {
	framework := flag.String("framework", "croissant", "Framework to use while testing")

	flag.Parse()

	switch *framework {
	case "croissant":
		TestCroissant()
	case "http":
		TestHttp()
	}
}

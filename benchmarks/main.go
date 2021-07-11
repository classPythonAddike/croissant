package main

import (
	"flag"
	_ "net/http/pprof"
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

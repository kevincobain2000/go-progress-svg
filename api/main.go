package main

import (
	"flag"

	"github.com/kevincobain2000/go-progress-svg/api/pkg"
)

type Flags struct {
	host    string
	port    string
	cors    string
	baseUrl string
}

var f Flags

func main() {
	SetupFlags()
	pkg.StartEcho(pkg.NewEcho(f.baseUrl, f.cors), f.host, f.port)
}

func SetupFlags() {
	flag.StringVar(&f.host, "host", "localhost", "host to serve")
	flag.StringVar(&f.port, "port", "3003", "port to serve")
	flag.StringVar(&f.cors, "cors", "", "cors port to allow")
	flag.StringVar(&f.baseUrl, "base-url", "/", "base url with slash")
	flag.Parse()

}

package main

import (
	"flag"
	"strings"

	"github.com/jtrotsky/govend/vend"
)

var (
	token        string
	domainPrefix string
	tz           string
)

func main() {
	vend.NewClient(token, domainPrefix, tz)
}

func init() {

	// Get auth and retailer info from command line flags.
	flag.StringVar(&domainPrefix, "d", "",
		"The Vend store name (prefix of xxxx.vendhq.com)")
	flag.StringVar(&token, "t", "",
		"Personal API Access Token for the store, generated from Setup -> API Access.")
	flag.StringVar(&tz, "z", "Local",
		"Timezone of the store in zoneinfo format. The default is to try and use the computer's local timezone.")
	flag.Parse()

	// To save people who write DomainPrefix.vendhq.com.
	// Split DomainPrefix on the "." period character then grab the first part.
	parts := strings.Split(domainPrefix, ".")
	domainPrefix = parts[0]
}

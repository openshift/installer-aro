package dnsmasq

import _ "embed"

//go:embed scripts/dnsmasq.conf.gotmpl
var configFile string

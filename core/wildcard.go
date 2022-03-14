package core

import (
	"net"
)

type Lookup struct {
	Domain string
	IP     []net.IP
}

func IsWildCard(domain string) (Lookup, bool) {
	var ret Lookup
	for i := 0; i < 2; i++ {
		subdomain := RandomStr(6) + "." + domain
		ips, err := net.LookupIP(subdomain)
		if err != nil {
			continue
		}
		ret.Domain = subdomain
		ret.IP = ips
		return ret, true
	}
	return ret, false
}

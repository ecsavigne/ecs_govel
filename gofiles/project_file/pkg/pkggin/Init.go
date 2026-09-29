package pkggin

import "slices"

func HostAllow() (host []string) {
	host = []string{
		"localhost",
		"docsgateway.savcoe-services.com",
		"gateway.savcoe-services.com",
	}

	slices.Sort(host)
	return
}

func init() {
	prepare_engine()
}

package pkgutil

import "strings"

func ParseHost(host string) string {
	host, _, _ = strings.Cut(strings.NewReplacer("https://", "", "http://", "").Replace(host), ":")

	return host
}

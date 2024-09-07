package learning

import "strings"

func IsEmployee(address string) bool {
	return strings.HasSuffix(address, "@linkedin.com")
}

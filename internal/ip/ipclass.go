package ip

import (
	"strconv"
	"strings"
)

func IpClass(ipaddr string) string {
	splitted := strings.Split(ipaddr, ".")
	converted, _ := strconv.Atoi(splitted[0])

	if converted <= 127 {
		return "A"
	}

	if converted <= 191 {
		return "B"
	}

	if converted <= 223 {
		return "C"
	}

	return "-"
}

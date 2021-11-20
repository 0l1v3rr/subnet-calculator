package ip

import (
	"strconv"
	"strings"
)

func BroadcastAddress(prefix string, ipInBinary string) string {
	res := ""

	converted, err := strconv.Atoi(prefix)
	if err != nil {
		return ""
	}

	ipWithoutPoints := ""
	splitted := strings.Split(ipInBinary, ".")
	for _, s := range splitted {
		ipWithoutPoints += s
	}

	for i := 0; i < 32; i++ {
		if i < converted {
			if i == 7 || i == 15 || i == 23 {
				res += string(ipWithoutPoints[i])
				res += "."
			} else {
				res += string(ipWithoutPoints[i])
			}
		} else {
			if i == 7 || i == 15 || i == 23 {
				res += "1."
			} else {
				res += "1"
			}
		}
	}

	return res
}

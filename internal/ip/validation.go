package ip

import (
	"strconv"
	"strings"
)

func IsValidIp(ip string) bool {
	splitted := strings.Split(ip, ".")
	if len(splitted) != 4 {
		return false
	}

	if strings.HasPrefix(ip, "0.") {
		return false
	}

	for _, s := range splitted {
		converted, err := strconv.Atoi(s)
		if err != nil {
			return false
		}

		if converted < 0 || converted > 255 {
			return false
		}
	}

	return true
}

func IsValidPrefix(mask string) bool {
	converted, err := strconv.Atoi(mask)
	if err != nil {
		return false
	}
	if converted < 1 || converted >= 30 {
		return false
	}
	return true
}

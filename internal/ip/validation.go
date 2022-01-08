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

func IsValidPrefix(p string) bool {
	converted, err := strconv.Atoi(p)
	if err != nil {
		return false
	}
	if converted < 1 || converted > 30 {
		return false
	}
	return true
}

func IsValidMask(mask string) bool {
	binary := ConvertToBinary(mask)
	if binary == "" {
		return false
	}
	binary = strings.ReplaceAll(binary, ".", "")

	if string(binary[len(binary)-1])+string(binary[len(binary)-2]) != "00" {
		return false
	}

	for i := 0; i < len(binary)-2; i++ {
		if string(binary[i]) == "1" && string(binary[i+1]) == "0" {
			for j := i + 1; j < len(binary); j++ {
				if string(binary[j]) == "1" {
					return false
				}
			}
		}
	}

	return true
}

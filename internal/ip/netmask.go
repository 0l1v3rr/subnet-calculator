package ip

import "strconv"

func CalculateMask(prefix string) string {
	res := ""

	converted, err := strconv.Atoi(prefix)
	if err != nil {
		return ""
	}

	for i := 1; i < 33; i++ {
		if i <= converted {
			if i == 8 || i == 16 || i == 24 {
				res += "1."
			} else {
				res += "1"
			}
		} else {
			if i == 8 || i == 16 || i == 24 {
				res += "0."
			} else {
				res += "0"
			}
		}
	}

	return res
}

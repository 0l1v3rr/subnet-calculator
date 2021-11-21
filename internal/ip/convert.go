package ip

import "strconv"

func ConvertMaskToBits(maskInBinary string) string {
	res := 0

	for i := 0; i < len(maskInBinary); i++ {
		if string(maskInBinary[i]) == "1" {
			res++
		}
	}

	return strconv.Itoa(res)
}

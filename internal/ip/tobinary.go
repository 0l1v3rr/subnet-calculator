package ip

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToBinary(ip string) string {
	res := ""
	splitted := strings.Split(ip, ".")

	for _, s := range splitted {
		c, err := strconv.Atoi(s)
		if err != nil {
			return ""
		}
		binary := strconv.FormatInt(int64(c), 2)
		res += fmt.Sprintf("%s.", addZeros(binary))
	}

	res = strings.TrimSuffix(res, ".")
	return res
}

func addZeros(s string) string {
	switch len(s) {
	case 0:
		return "00000000"
	case 1:
		return fmt.Sprintf("0000000%s", s)
	case 2:
		return fmt.Sprintf("000000%s", s)
	case 3:
		return fmt.Sprintf("00000%s", s)
	case 4:
		return fmt.Sprintf("0000%s", s)
	case 5:
		return fmt.Sprintf("000%s", s)
	case 6:
		return fmt.Sprintf("00%s", s)
	case 7:
		return fmt.Sprintf("0%s", s)
	}

	return s
}

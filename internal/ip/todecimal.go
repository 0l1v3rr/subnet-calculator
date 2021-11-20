package ip

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToDecimal(ip string) string {
	res := ""
	splitted := strings.Split(ip, ".")

	for _, s := range splitted {
		dec, _ := strconv.ParseInt(s, 2, 64)
		res += fmt.Sprintf("%d.", dec)
	}

	res = strings.TrimSuffix(res, ".")
	return res
}

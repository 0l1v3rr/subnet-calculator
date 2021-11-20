package ip

import (
	"fmt"
	"strconv"
	"strings"
)

func SubtractIps(ip1 string, ip2 string) string {
	res := ""

	splitted1 := strings.Split(ip1, ".")
	splitted2 := strings.Split(ip2, ".")

	for i := 0; i < 4; i++ {
		converted1, _ := strconv.Atoi(splitted1[i])
		converted2, _ := strconv.Atoi(splitted2[i])
		res += fmt.Sprintf("%d.", converted1-converted2)
	}

	res = strings.TrimSuffix(res, ".")

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ips "github.com/0l1v3rr/subnet-calculator/internal/ip"
)

var (
	ip   string
	mask string
	bits string
)

func main() {
	title()
	reader := bufio.NewReader(os.Stdin)

	for {
		message("$", "Specify the IP address you want to calculate:")

		prompt("ip")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		fmt.Print("\u001b[0m")

		if strings.HasPrefix(input, "ex") {
			return
		}

		if input == "" || input == " " {
			continue
		}

		if len(strings.Split(input, " ")) >= 2 {
			splitted := strings.Split(input, " ")
			if !ips.IsValidIp(splitted[0]) {
				error("Please specify a valid IP address!\n")
				continue
			}

			if !ips.IsValidMask(splitted[1]) {
				error("Please specify a valid network mask!\n")
				continue
			}

			ip = splitted[0]
			mask = splitted[1]
			bits = ips.ConvertMaskToBits(ips.ConvertToBinary(mask))

		} else if len(strings.Split(input, "/")) == 2 {
			splitted := strings.Split(input, "/")
			if !ips.IsValidIp(splitted[0]) {
				error("Please specify a valid IP address!\n")
				continue
			}

			if !ips.IsValidPrefix(splitted[1]) {
				error("Please specify a valid prefix!\n")
				continue
			}

			ip = splitted[0]
			bits = splitted[1]
			mask = ips.ConvertToDecimal(ips.CalculateMask(splitted[1]))

		} else if ips.IsValidIp(input) {
			ip = input
			if ips.IpClass(ip) == "A" {
				bits = "8"
				mask = "255.0.0.0"
			} else if ips.IpClass(ip) == "B" {
				bits = "16"
				mask = "255.255.0.0"
			} else {
				bits = "24"
				mask = "255.255.255.0"
			}
		} else {
			error("Please specify a valid IP address!\n")
			continue
		}

		resTitle()
		printInfo()
	}
}

func printInfo() {
	fmt.Print("\u001b[34;1m[*] IP Address:      ")
	fmt.Printf("\u001b[0m%s\n", ip)

	fmt.Print("\u001b[34;1m[*] IP in Binary:    ")
	fmt.Printf("\u001b[0m%s\n", ips.ConvertToBinary(ip))

	fmt.Print("\u001b[34;1m[*] Network mask:    ")
	fmt.Printf("\u001b[0m%s\n", mask)

	fmt.Print("\u001b[34;1m[*] Mask in Binary:  ")
	fmt.Printf("\u001b[0m%s\n", ips.CalculateMask(bits))

	fmt.Print("\u001b[34;1m[*] Wildcard:        ")
	fmt.Printf("\u001b[0m%s\n", ips.SubtractIps("255.255.255.255", ips.ConvertToDecimal(ips.CalculateMask(bits))))

	fmt.Print("\u001b[34;1m[*] Network address: ")
	fmt.Printf("\u001b[0m%s\n", ips.ConvertToDecimal(ips.NetworkAddress(bits, ips.ConvertToBinary(ip))))

	fmt.Print("\u001b[34;1m[*] Broadcast:       ")
	fmt.Printf("\u001b[0m%s\n", ips.ConvertToDecimal(ips.BroadcastAddress(bits, ips.ConvertToBinary(ip))))

	fmt.Print("\u001b[34;1m[*] Min IP:          ")
	fmt.Printf("\u001b[0m%s\n", ips.AddIps(ips.ConvertToDecimal(ips.NetworkAddress(bits, ips.ConvertToBinary(ip))), "0.0.0.1"))

	fmt.Print("\u001b[34;1m[*] Max IP:          ")
	fmt.Printf("\u001b[0m%s\n", ips.SubtractIps(ips.ConvertToDecimal(ips.BroadcastAddress(bits, ips.ConvertToBinary(ip))), "0.0.0.1"))

	fmt.Println()
}

func message(prefix string, m string) {
	fmt.Printf("\u001b[32;1m[%s] \u001b[0m%s\n", prefix, m)
}

func error(e string) {
	fmt.Printf("\u001b[31;1m[!] \u001b[0m%s\n", e)
}

func prompt(p string) {
	fmt.Printf("\u001b[4m%s\u001b[0m > \u001b[36;1m", p)
}

func title() {
	fmt.Println()
	fmt.Println("\u001b[33;1m --==<[{ \u001b[37;1mIPv4 Subnet Calculator \u001b[33;1m }]>==--")
	fmt.Println()
}

func resTitle() {
	fmt.Println()
	fmt.Printf("\u001b[32;1m --==<[{ \u001b[37;1m%s/%s \u001b[32;1m }]>==--\n", ip, bits)
	fmt.Println()
}

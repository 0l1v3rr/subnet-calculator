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
)

func main() {
	title()
	reader := bufio.NewReader(os.Stdin)

	message("Please, specify the IP address you want to calculate:")
	for {
		prompt("ip")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		fmt.Print("\u001b[0m")

		if input == "" || input == " " {
			continue
		}

		if ips.IsValidIp(input) {
			ip = input
			break
		}
		error("Please specify a valid IP address!")
	}

	message("Please, specify the netmask you want to calculate with:")
	for {
		prompt("mask")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		fmt.Print("\u001b[0m")

		if input == "" || input == " " {
			continue
		}

		if ips.IsValidMask(input) {
			mask = input
			break
		}
		error("Please specify a valid netmask!")
	}

	resTitle()
	printInfo()
}

func printInfo() {
	fmt.Print("\u001b[34;1m$ IP Address:        ")
	fmt.Printf("\u001b[0m%s\n", ip)

	fmt.Print("\u001b[34;1m$ IP in Binary:      ")
	fmt.Printf("\u001b[0m%s\n", ips.ConvertToBinary(ip))

	fmt.Print("\u001b[34;1m$ Network mask:      ")
	fmt.Printf("\u001b[0m%s\n", ips.ConvertToDecimal(ips.CalculateMask(mask)))

	fmt.Print("\u001b[34;1m$ Mask in Binary:    ")
	fmt.Printf("\u001b[0m%s\n", ips.CalculateMask(mask))

	fmt.Print("\u001b[34;1m$ Wildcard:          ")
	fmt.Printf("\u001b[0m%s\n", ips.SubtractIps("255.255.255.255", ips.ConvertToDecimal(ips.CalculateMask(mask))))

	fmt.Print("\u001b[34;1m$ Network address:   ")
	fmt.Printf("\u001b[0m%s\n", ips.ConvertToDecimal(ips.NetworkAddress(mask, ips.ConvertToBinary(ip))))

	fmt.Print("\u001b[34;1m$ Broadcast address: ")
	fmt.Printf("\u001b[0m%s\n", ips.ConvertToDecimal(ips.BroadcastAddress(mask, ips.ConvertToBinary(ip))))

	fmt.Print("\u001b[34;1m$ Min IP:            ")
	fmt.Printf("\u001b[0m%s\n", ips.AddIps(ips.ConvertToDecimal(ips.NetworkAddress(mask, ips.ConvertToBinary(ip))), "0.0.0.1"))

	fmt.Print("\u001b[34;1m$ Max IP:            ")
	fmt.Printf("\u001b[0m%s\n", ips.SubtractIps(ips.ConvertToDecimal(ips.BroadcastAddress(mask, ips.ConvertToBinary(ip))), "0.0.0.1"))
}

func message(m string) {
	fmt.Printf("\u001b[32;1m[$] \u001b[0m%s\n", m)
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
	fmt.Printf("\u001b[32;1m --==<[{ \u001b[37;1m%s/%s \u001b[32;1m }]>==--\n", ip, mask)
	fmt.Println()
}

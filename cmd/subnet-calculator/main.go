package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ips "github.com/0l1v3rr/subnet-calculator/internal/ip"
)

var (
	ip     string
	mask   string
	bits   string
	choice string
)

func main() {
	title()
	reader := bufio.NewReader(os.Stdin)

	for {
		message("$", "Specify the IP address you want to calculate:")
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

		message("$", "Which one do you want to specify?")
		fmt.Println(" \u001b[32;1m~ [1] \u001b[0mNetwork mask (E.g 255.255.255.0)")
		fmt.Println(" \u001b[32;1m~ [2] \u001b[0mNumber of bits identifying the network (E.g 24)")
		for {
			prompt("choose")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSuffix(input, "\n")
			fmt.Print("\u001b[0m")

			if input == "" || input == " " {
				continue
			}

			if input == "1" || input == "2" {
				choice = input
				break
			}
			error("Please specify 1 or 2!")
		}

		if choice == "1" {
			for {
				prompt("mask")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSuffix(input, "\n")
				fmt.Print("\u001b[0m")

				if input == "" || input == " " {
					continue
				}

				if ips.IsValidIp(input) {
					mask = input
					break
				}
				error("Please specify a valid network mask!")
			}
		} else {
			for {
				prompt("bits")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSuffix(input, "\n")
				fmt.Print("\u001b[0m")

				if input == "" || input == " " {
					continue
				}

				if ips.IsValidMask(input) {
					bits = input
					break
				}
				error("Please specify a valid argument!")
			}
		}

		if choice == "1" {
			bits = ips.ConvertMaskToBits(ips.ConvertToBinary(mask))
		} else {
			mask = ips.ConvertToDecimal(ips.CalculateMask(bits))
		}

		resTitle()
		printInfo()

		message("$", "If you want to exit, type 'yes', else type 'no'!")
		prompt("choose")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		fmt.Print("\u001b[0m")

		if strings.HasPrefix(input, "y") {
			break
		}
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

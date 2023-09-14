package main

import (
	"CosmoEth/feeder/evm"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("Feeder Daemon (server)")
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Printf("    feederd [command]\n")
	fmt.Println()

	fmt.Println("Available Commands:")
	fmt.Printf("    submit-state	:	Submit state value with proofs to the chain.\n")
	fmt.Println()

	fmt.Println("Flags:")
	fmt.Printf("    -h, --help	:	help for feederd.\n")
	fmt.Println()

	fmt.Println("Use \"feederd [command] --help\" for more information about a command.")
}

func main() {
	submitCommand := flag.NewFlagSet("submit-state", flag.ExitOnError)

	addressStr := submitCommand.String("address", "", "address on ethereum network")
	slotStr := submitCommand.String("slot", "", "storage slot at address")

	flag.Usage = printUsage
	flag.Parse()

	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "submit-state":
			err := submitCommand.Parse(os.Args[2:])
			if err != nil {
				panic(err)
			}

			if submitCommand.Parsed() {
				if *addressStr == "" {
					fmt.Println("Please provide an address using the -address flag.")
				}

				if *slotStr == "" {
					fmt.Println("Please provide a slot using the -slot flag.")
				}

				slots := strings.Split(*slotStr, ",")
				sproof, err := evm.GetProof(context.Background(), *addressStr, slots)
				if err != nil {
					panic(err)
				}

				bz, _ := json.Marshal(sproof)
				fmt.Println(string(bz))
			}

		default:
			fmt.Println("submit-state command is only available.")
			os.Exit(1)
		}
	}
}

package main

import (
	"CosmoEth/feeder/evm"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
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

// runTxCommandForStateSubmit handles tx broadcast to CosmoEth chain for submitting state information with proofs
func runTxCommandForStateSubmit(from, address, root, proof string, height uint64) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// Prepare the command with arguments
	command := fmt.Sprintf("CosmoEthd tx cosmoeth add-state \"%s\" %d \"%s\" '%s' --from %s -y --gas auto", address, height, root, proof, from)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command and capture its output
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
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

				// fetch storage proof
				slots := strings.Split(*slotStr, ",")
				sproof, err := evm.GetProof(context.Background(), *addressStr, slots)
				if err != nil {
					panic(err)
				}

				proofArg := ""
				for _, proofResult := range sproof.StorageProof {
					if len(proofArg) > 0 {
						proofArg += ";"
					}
					bz, _ := json.Marshal(proofResult)
					proofArg += string(bz)
				}

				stdOut, stdErr, _ := runTxCommandForStateSubmit("alice", sproof.Address.Hex(), sproof.StorageHash.Hex(), proofArg, sproof.Height.Uint64())
				fmt.Println(stdOut)
				fmt.Println(stdErr)
			}

		default:
			fmt.Println("submit-state command is only available.")
			os.Exit(1)
		}
	}
}

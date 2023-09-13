package utils

import (
	"regexp"
)

func IsEthereumAddress(address string) bool {
	// Define a regular expression pattern for a valid Ethereum address
	ethereumAddressPattern := `^0x[0-9a-fA-F]{40}$`

	// Compile the regular expression
	re := regexp.MustCompile(ethereumAddressPattern)

	// Use the regular expression to match the address
	return re.MatchString(address)
}

func IsValidHexString(hexString string) bool {
	// Define a regular expression pattern for a valid hexadecimal string with "0x" prefix
	hexPattern := `^0x[0-9a-fA-F]+$`

	// Compile the regular expression
	re := regexp.MustCompile(hexPattern)

	// Use the regular expression to match the hex string
	return re.MatchString(hexString)
}

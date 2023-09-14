package utils

import (
	"bytes"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/trie"
)

func IsValidHexString(hexString string) bool {
	// Define a regular expression pattern for a valid hexadecimal string with "0x" prefix
	hexPattern := `^0x[0-9a-fA-F]+$`

	// Compile the regular expression
	re := regexp.MustCompile(hexPattern)

	// Use the regular expression to match the hex string
	return re.MatchString(hexString)
}

func VerifyProof(rootHash common.Hash, key string, value []byte, proofs []string) (bool, error) {
	proofDB := NewMemDB()
	for _, proof := range proofs {
		node := common.FromHex(proof)
		key := crypto.Keccak256(node)
		proofDB.Put(key, node)
	}

	path := crypto.Keccak256(common.LeftPadBytes(common.FromHex(key), 32))

	res, err := trie.VerifyProof(rootHash, path, proofDB)
	if err != nil {
		return false, err
	}
	return bytes.Equal(value, res), nil
}

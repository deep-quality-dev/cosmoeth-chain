package types

const (
	// ModuleName defines the module name
	ModuleName = "cosmoeth"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosmoeth"

	// StateKey defines the state store prefix key
	StateKey = "state"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

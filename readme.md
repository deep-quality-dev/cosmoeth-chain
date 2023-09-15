# CosmoEth and Feeder
**CosmoEth** is a blockchain constructed using the Cosmos SDK and Tendermint. Its primary purpose is to securely store verified state values sourced from the Ethereum network.

**Feeder**, on the other hand, is a Go (Golang) application designed to retrieve state values through Ethereum Virtual Machine (EVM) addresses and storage slot indices. Subsequently, it facilitates the submission of transactions to the CosmoEth chain, complete with state values and corresponding proofs, ensuring data integrity and security.

## Get Started
### 1. Setup CosmoEth localnet
#### *Install ignite CLI*
```
curl https://get.ignite.com/cli\! | bash
```
Please ensure that the Ignite CLI version is at or above v0.27.1.

#### *Spin up the localnet*

```
ignite chain serve
```
`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### 2. Setup feeder
```
make build-feeder
```
Please ensure the feeder binary is created in /build folder.

## How to test
### *Submit ethereum state to CosmoEth chain*
```
./build/feeder submit-state -address [evm address] -slot [hex presentation of slot index]
```
example => 
```
./build/feeder submit-state -address 0xdac17f958d2ee523a2206206994597c13d831ec7 -slot 0x1
```

### *Query the storage from CosmoEth chain*
```
CosmoEthd query cosmoeth state-value "0xdac17f958d2ee523a2206206994597c13d831EC7" "0x1"  
```

## Future improvements
- Setup real time threads for submitting latest state value to the chain (currently, state values stored into the chain are updated only by running the command manually by feeder)
- Setup a system to incentivize feeders
- Setup feeder whitelist and manage it through governance
- Enable RPC node configurability in the feeder application, implement RPC node list management, and enable automatic switching to alternative nodes in case of issues.
- Prevent submitting state value to the chain if the latest value fetched from RPC is same as the stored one.
- Optimize state kvstore by decreasing the size of state metadata

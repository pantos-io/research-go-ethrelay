# Go-ETH Relay
This project contains a Go-library and a command-line interface (CLI) to interact with the [ETH Relay](https://github.com/pantos-io/ethrelay) prototype.
 
ETH Relay enables the cross-blockchain verification of transactions.
This makes it possible that a "verifying" (destination) blockchain can verify if a certain transaction (or receipt, or state) is included 
in a different "target" (source) blockchain without relying or trusting a third-party.

Detailed information about how the prototype works can be found [here](https://dsg.tuwien.ac.at/projects/tast/pub/tast-white-paper-6.pdf).

> _Important: ETH Relay is a research prototype. 
    It represents ongoing work conducted within the [TAST](https://dsg.tuwien.ac.at/projects/tast/) 
    research project. Use with care._

## Prerequisites
You need to have [Golang](https://golang.org/doc/install) and [Ganache](https://www.trufflesuite.com/ganache) (>= 2.1.0) installed. 
## Get Started
_The following setup will take you through the deployment of ETH Relay with a local Ethereum blockchain (Ganache)
as verifying chain and the main Ethereum chain as target chain.
Information on how to connect other blockchains can be found [here](#Configuration)._


1. Install the library and CLI with `$ go get github.com/pantos-io/go-ethrelay`.
Check that the CLI was installed correctly by running `$ go-ethrelay --help`.
If you want to install the library manually, you can simply clone this repository and run any command in the cloned direcoty with `go run main.go [command]`.

2. Run `go-ethrelay init` or `go run main.go init` to initialize the client.
If you encounter any problems calling this command, get sure the rights are properly adjusted so Go can create the testimonium.yml config file in the current folder.
It is also possible to generate the file by hand or change the example config file named testimonium.example.yml contained in this repo. 

3. Start Ganache (should start on the default port 7545, if not, change this in the config file)

4. Deploy the Ethash contract with `go-ethrelay deploy ethash`. 
This deploys the contract responsible for verifying the Proof of Work (PoW) of a block.

5. Submit the correct epoch data to the Ethash contract with `go-ethrelay submit epoch <EPOCH_NO>`.
Depending on which block will be submitted as genesis block to the ETH Relay contract, 
the correct epoch data can be calculated as `EPOCH_NO = BLOCK_NO / 30000` floored. This may take a while. 
    > e.g., for genesis block 8084509, the correct epoch data is 269

6. Deploy the ETH Relay contract with `go-ethrelay deploy ethrelay --genesis <BLOCK_NO>`.
This deploys the contract responsible for the verification of transactions (or receipts, or state). 
The `genesis` parameter specifies the first block of the target chain which will be submitted to 
the ETH Relay contract. Verifications will be possible for all subsequent blocks.

###
ETH Relay is now setup. In order to submit blocks, a stake has to be submitted first. After you deposited some stake, you can submit block data from the target chain to the verifying chain, 
and request verifications of transactions, and dispute illegal blocks. 

## Usage
The CLI can be started with `go-ethrelay [command]` where `[command]` is one of the commands below.

Use `go-ethrelay [command] --help` for more information about a command.

---

`init`: Initializes the client by creating a ethrelay.yml file in the current directory that acts as config file for all command calls.

`account`: Prints the address of the current account

`balance`: Prints the balance of the current account

`deploy ethash`: Deploys the Ethash smart contract on the verifying chain

`deploy ethrelay`: Deploys the ETH Relay contract on the verifying chain

`dispute [blockHash]`: Disputes the submitted block header with the specified hash

`get block [blockHash]`: Retrieves the block with the specified hash

`get transaction [txHash]`: Retrieves the transaction with the specified hash

`get longestchainendpoint`: Retrieves the most recent block hash of the longest chain in the eth relay contract on the verifying chain

`stake`: Retrieves the amount of stake deposited in the relay-contract on the verifying chain

`stake deposit [amountInWei]`: Deposits amountInWei stake of the account balance in the contract

> e.g. `stake deposit 25000000000000000000` deposits 25 ETH

`stake withdraw [amountInWei]`: Withdraws the submitted stake back to the account balance. Remember that stake can be locked in the contract when a block was submitted and you have to wait until it is unlocked again.

`submit block [blocknumber]`: Submits the specified block header from the target chain to the verifying chain

`submit epoch [epoch]`: Sets the epoch data for the specified epoch on the verifying chain

`verify block [blockHash]`: Verifies a block from the target chain on the verifying chain

`verify transaction [txHash]`: Verifies a transaction from the target chain on the verifying chain

`verify receipt [txHash]`: Verifies a receipt from the target chain on the verifying chain

## Quick Setup

There is also a shell script in this repository named `setup-relay.sh`. This script helps researchers and developers to quickly setup
a working version of the relay with default-values and assuming a local Ganache instance. The call to this script is:

```
sh setup-realy.sh [ganacheAccountPrivateKey] [genesisBlock] [stakeInETH]

// example call
sh setup-realy.sh 0x45b5ffd7266ec7131f31f94da843b99fd270b46d94bf01368ceeb936649dfc3b 11367417 25
```

## Configuration
The relay client uses a configuration file called `ethrelay.yml` file.

The default file looks like this:

    privatekey: <YOUR PRIVATE KEY>
    chains:
        0:
            url: mainnet.infura.io
        1:
            type: http
            url: localhost
            port: 8545

Chain ID 0 contains connection configuration for the main Ethereum chain (via Infura).
Chain ID 1 contains connection configuration for a local chain (e.g., run via Ganache).

You can configure the relay client for other Ethereum blockchains (there is no upper limit).
Just manually add or edit a chain entry under the `chains` key.
Key `type` refers to the connections type (e.g., http, https, ws, wss), 
`url` refers to the URL, 
and `port` refers to the port number under which the specific chain is reachable. If no type is specified, https is used.
If no port is defined, it is determined by the default port of the type.


If you have already deployed the Ethash and ETH Relay contracts, you might find further entries
`ethashAddress` and `ethrelayAddress` under a specific chain config:

    ...
    chains:
        ...
        1:
            type: http
            url: localhost
            port: 8545
            ethrelayAddress: 0xabc123...
            ethashAddress: 0x123abc...

These are the addresses that the client uses to interact with the ETH Relay smart contracts.
If you deployed the contracts manually, just add the entries.

## Troubleshooting
#### Dispute causes error: "VM Exception while processing transaction: revert"
If disputing a certain block causes a generic revert exception, make sure you are running Ganache version >= 2.1.0.

#### Client won't start because of Go problems
Get sure your Go-path variables like GOHOME, GOPATH und GOBIN are set properly.

#### Client won't work because issues to permissions or projectId (e.g. forbidden: project ID is required)
If you are using Infura, check that your url and protocol is correct and the permissions are properly set in the Infura admin panel.

## How to Contribute
ETH Relay is a research prototype. We welcome anyone to contribute.
File a bug report or submit feature requests through the [issue tracker](https://github.com/pantos-io/go-ethrelay/issues). 
If you want to contribute feel free to submit a pull request.

## Acknowledgement
The development of this prototype was funded by [Pantos](https://pantos.io/) within the [TAST](https://dsg.tuwien.ac.at/projects/tast/) research project.

## Licence
This project is licensed under the [MIT License](LICENSE).
